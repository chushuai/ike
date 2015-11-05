// generated by stringer -type NotificationType -output notification_type_strings.go; DO NOT EDIT

package protocol

import "fmt"

const _NotificationType_name = "UNSUPPORTED_CRITICAL_PAYLOADINVALID_IKE_SPIINVALID_MAJOR_VERSIONINVALID_SYNTAXINVALID_MESSAGE_IDINVALID_SPINO_PROPOSAL_CHOSENINVALID_KE_PAYLOADAUTHENTICATION_FAILEDSINGLE_PAIR_REQUIREDNO_ADDITIONAL_SASINTERNAL_ADDRESS_FAILUREFAILED_CP_REQUIREDTS_UNACCEPTABLEINVALID_SELECTORSTEMPORARY_FAILURECHILD_SA_NOT_FOUNDINITIAL_CONTACTSET_WINDOW_SIZEADDITIONAL_TS_POSSIBLEIPCOMP_SUPPORTEDNAT_DETECTION_SOURCE_IPNAT_DETECTION_DESTINATION_IPCOOKIEUSE_TRANSPORT_MODEHTTP_CERT_LOOKUP_SUPPORTEDREKEY_SAESP_TFC_PADDING_NOT_SUPPORTEDNON_FIRST_FRAGMENTS_ALSOMOBIKE_SUPPORTEDADDITIONAL_IP4_ADDRESSADDITIONAL_IP6_ADDRESSNO_ADDITIONAL_ADDRESSESUPDATE_SA_ADDRESSESCOOKIE2NO_NATS_ALLOWEDAUTH_LIFETIMEMULTIPLE_AUTH_SUPPORTEDANOTHER_AUTH_FOLLOWSREDIRECT_SUPPORTEDREDIRECTREDIRECTED_FROMTICKET_LT_OPAQUETICKET_REQUESTTICKET_ACKTICKET_NACKTICKET_OPAQUELINK_IDUSE_WESP_MODEROHC_SUPPORTEDEAP_ONLY_AUTHENTICATIONCHILDLESS_IKEV2_SUPPORTEDQUICK_CRASH_DETECTIONIKEV2_MESSAGE_ID_SYNC_SUPPORTEDIPSEC_REPLAY_COUNTER_SYNC_SUPPORTEDIKEV2_MESSAGE_ID_SYNCIPSEC_REPLAY_COUNTER_SYNCSECURE_PASSWORD_METHODSPSK_PERSISTPSK_CONFIRMERX_SUPPORTEDIFOM_CAPABILITYSENDER_REQUEST_IDIKEV2_FRAGMENTATION_SUPPORTEDSIGNATURE_HASH_ALGORITHMS"

var _NotificationType_map = map[NotificationType]string{
	1:     _NotificationType_name[0:28],
	4:     _NotificationType_name[28:43],
	5:     _NotificationType_name[43:64],
	7:     _NotificationType_name[64:78],
	9:     _NotificationType_name[78:96],
	11:    _NotificationType_name[96:107],
	14:    _NotificationType_name[107:125],
	17:    _NotificationType_name[125:143],
	24:    _NotificationType_name[143:164],
	34:    _NotificationType_name[164:184],
	35:    _NotificationType_name[184:201],
	36:    _NotificationType_name[201:225],
	37:    _NotificationType_name[225:243],
	38:    _NotificationType_name[243:258],
	39:    _NotificationType_name[258:275],
	43:    _NotificationType_name[275:292],
	44:    _NotificationType_name[292:310],
	16384: _NotificationType_name[310:325],
	16385: _NotificationType_name[325:340],
	16386: _NotificationType_name[340:362],
	16387: _NotificationType_name[362:378],
	16388: _NotificationType_name[378:401],
	16389: _NotificationType_name[401:429],
	16390: _NotificationType_name[429:435],
	16391: _NotificationType_name[435:453],
	16392: _NotificationType_name[453:479],
	16393: _NotificationType_name[479:487],
	16394: _NotificationType_name[487:516],
	16395: _NotificationType_name[516:540],
	16396: _NotificationType_name[540:556],
	16397: _NotificationType_name[556:578],
	16398: _NotificationType_name[578:600],
	16399: _NotificationType_name[600:623],
	16400: _NotificationType_name[623:642],
	16401: _NotificationType_name[642:649],
	16402: _NotificationType_name[649:664],
	16403: _NotificationType_name[664:677],
	16404: _NotificationType_name[677:700],
	16405: _NotificationType_name[700:720],
	16406: _NotificationType_name[720:738],
	16407: _NotificationType_name[738:746],
	16408: _NotificationType_name[746:761],
	16409: _NotificationType_name[761:777],
	16410: _NotificationType_name[777:791],
	16411: _NotificationType_name[791:801],
	16412: _NotificationType_name[801:812],
	16413: _NotificationType_name[812:825],
	16414: _NotificationType_name[825:832],
	16415: _NotificationType_name[832:845],
	16416: _NotificationType_name[845:859],
	16417: _NotificationType_name[859:882],
	16418: _NotificationType_name[882:907],
	16419: _NotificationType_name[907:928],
	16420: _NotificationType_name[928:959],
	16421: _NotificationType_name[959:994],
	16422: _NotificationType_name[994:1015],
	16423: _NotificationType_name[1015:1040],
	16424: _NotificationType_name[1040:1063],
	16425: _NotificationType_name[1063:1074],
	16426: _NotificationType_name[1074:1085],
	16427: _NotificationType_name[1085:1098],
	16428: _NotificationType_name[1098:1113],
	16429: _NotificationType_name[1113:1130],
	16430: _NotificationType_name[1130:1159],
	16431: _NotificationType_name[1159:1184],
}

func (i NotificationType) String() string {
	if str, ok := _NotificationType_map[i]; ok {
		return str
	}
	return fmt.Sprintf("NotificationType(%d)", i)
}
