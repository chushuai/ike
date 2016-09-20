package main

import (
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/net/ipv4"

	"github.com/msgboxio/context"
	"github.com/msgboxio/ike"
	"github.com/msgboxio/ike/platform"
	"github.com/msgboxio/log"
)

func waitForSignal(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	sig := <-c
	// sig is a ^C, handle it
	cancel(errors.New("received signal: " + sig.String()))
}

func loadConfig() (*ike.Config, string, string, *x509.CertPool, []*x509.Certificate, *rsa.PrivateKey) {
	var localString, remoteString string
	flag.StringVar(&localString, "local", "0.0.0.0:5000", "address to bind to")
	flag.StringVar(&remoteString, "remote", "", "address to connect to")

	var isTunnelMode bool
	flag.BoolVar(&isTunnelMode, "tunnel", false, "use tunnel mode?")

	var caFile, certFile, keyFile string
	flag.StringVar(&caFile, "ca", "", "PEM encoded ca certificate")
	flag.StringVar(&certFile, "cert", "", "PEM encoded peer certificate")
	flag.StringVar(&keyFile, "key", "", "PEM encoded peer key")

	flag.Set("logtostderr", "true")
	flag.Parse()

	config := ike.DefaultConfig()
	if !isTunnelMode {
		config.IsTransportMode = true
	}
	roots, err := ike.LoadRoot(caFile)
	if err != nil {
		log.Fatal(err)
	}
	certs, err := ike.LoadCerts(certFile)
	if err != nil {
		log.Warningf("Cert: %s", err)
	}
	key, err := ike.LoadKey(keyFile)
	if err != nil {
		log.Warningf("Key: %s", err)
	}

	return config, localString, remoteString, roots, certs, key
}

// map of initiator spi -> session
var sessions = make(map[uint64]*ike.Session)

// runs in goroutine.
// TODO - need exclusive access to sessions map
// TODO - merge into session.run goroutine; reduce from 3 -> 2 per conn
func runSession(spi uint64, session *ike.Session, pconn *ipv4.PacketConn, to net.Addr) {
	sessions[spi] = session
	for {
		select {
		case reply, ok := <-session.Replies():
			if !ok {
				break
			}
			if err := ike.WritePacket(pconn, reply, to); err != nil {
				session.Close(err)
				break
			}
		case <-session.Done():
			delete(sessions, spi)
			log.Infof("Finished SA 0x%x", spi)
			return
		}
	}
}

// var localId = &ike.PskIdentities{
// Primary: "ak@msgbox.io",
// Ids:     map[string][]byte{"ak@msgbox.io": []byte("foo")},
// }
var localId = &ike.RsaCertIdentity{}

// var remoteId = &ike.PskIdentities{
// Primary: "bk@msgbox.io",
// Ids:     map[string][]byte{"bk@msgbox.io": []byte("foo")},
// }
var remoteId = &ike.RsaCertIdentity{}

// runs on main thread
// loops until there is a socket error
func processPackets(pconn *ipv4.PacketConn, config *ike.Config) {
	for {
		msg, err := ike.ReadMessage(pconn)
		if err != nil {
			log.Error(err)
			break
		}
		// convert spi to uint64 for map lookup
		spi := ike.SpiToInt(msg.IkeHeader.SpiI)
		// check if a session exists
		session, found := sessions[spi]
		if !found {
			// create and run session
			session, err = ike.NewResponder(context.Background(), localId, remoteId, config, msg)
			if err != nil {
				log.Error(err)
				continue
			}
			go runSession(spi, session, pconn, msg.RemoteAddr)
		}
		session.HandleMessage(msg)
	}
}

func main() {
	config, localString, remoteString, roots, cert, key := loadConfig()
	cxt, cancel := context.WithCancel(context.Background())
	go waitForSignal(cancel)

	if cert != nil {
		localId.Certificate = cert[0]
	}
	localId.PrivateKey = key
	remoteId.Roots = roots

	// this should load the xfrm modules
	// requires root
	if xfrm := platform.ListenForEvents(cxt); xfrm != nil {
		go func() {
			<-xfrm.Done()
			if err := xfrm.Err(); err != context.Canceled {
				log.Error(err)
			}
			xfrm.Close()
		}()
	}

	pconn, err := ike.Listen(localString)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("socket listening: %s", pconn.Conn.LocalAddr())

	// requires root
	if err := platform.SetSocketBypas(pconn.Conn, syscall.AF_INET); err != nil {
		log.Error(err)
	}

	if remoteString != "" {
		remoteAddr, _ := net.ResolveUDPAddr("udp4", remoteString)
		initiator := ike.NewInitiator(context.Background(), localId, remoteId, ike.AddrToIp(remoteAddr).To4(), config)
		go runSession(ike.SpiToInt(initiator.IkeSpiI), initiator, pconn, remoteAddr)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		// wait for app shutdown
		<-cxt.Done()
		// shutdown sessions
		for _, session := range sessions {
			// rely on this to drain replies
			session.Close(cxt.Err())
			// wait until client is done
			<-session.Done()
		}
		pconn.Close()
		wg.Done()
	}()

	// this will return when there is a socket error
	// usually caused by the close call above
	processPackets(pconn, config)
	cancel(context.Canceled)

	wg.Wait()
	fmt.Printf("shutdown: %v\n", cxt.Err())
}
