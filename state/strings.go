// generated by stringer -type IkeSaState,IkeEventId -output strings.go; DO NOT EDIT

package state

import "fmt"

const _IkeSaState_name = "SMI_INITSMI_INIT_WAITSMI_AUTH_WAITSMI_AUTH_PEERSMI_EAPSMI_INSTALLCSA_DLSMI_INSTALLCSASMR_INITSMR_AUTHSMR_AUTH_FINALIZESMR_AUTH_RESPONSE_IDSMR_AUTH_RESPONSESMR_EAP_INITATOR_REQUESTSMR_EAP_AAA_REQUESTSMR_AUTH_DL_PEERSMR_CFG_WAITSM_MATURESM_REKEYSM_CRL_UPDATESM_REAUTHSM_TERMINATESM_DYINGSM_DEAD"

var _IkeSaState_index = [...]uint16{0, 8, 21, 34, 47, 54, 71, 85, 93, 101, 118, 138, 155, 179, 198, 214, 226, 235, 243, 256, 265, 277, 285, 292}

func (i IkeSaState) String() string {
	i -= 1
	if i < 0 || i+1 >= IkeSaState(len(_IkeSaState_index)) {
		return fmt.Sprintf("IkeSaState(%d)", i+1)
	}
	return _IkeSaState_name[_IkeSaState_index[i]:_IkeSaState_index[i+1]]
}

const _IkeEventId_name = "ACQUIRECONNECTREAUTHN_COOKIEN_INVALID_KEN_NO_PROPOSAL_CHOSENIKE_SA_INITIKE_AUTHDELETE_IKE_SACREATE_CHILD_SAIKE_SA_INIT_SUCCESSIKE_AUTH_SUCCESSDELETE_IKE_SA_SUCCESSCREATE_CHILD_SA_SUCCESSINVALID_KECONNECTION_ERRORMSG_IKE_REKEYMSG_IKE_DPDMSG_IKE_CRL_UPDATEMSG_IKE_REAUTHMSG_IKE_TERMINATEMSG_DELETE_IKE_SAIKE_TIMEOUTStateEntryStateExit"

var _IkeEventId_index = [...]uint16{0, 7, 14, 20, 28, 40, 60, 71, 79, 92, 107, 126, 142, 163, 186, 196, 212, 225, 236, 254, 268, 285, 302, 313, 323, 332}

func (i IkeEventId) String() string {
	i -= 1
	if i < 0 || i+1 >= IkeEventId(len(_IkeEventId_index)) {
		return fmt.Sprintf("IkeEventId(%d)", i+1)
	}
	return _IkeEventId_name[_IkeEventId_index[i]:_IkeEventId_index[i+1]]
}
