// generated by stringer -type IkeExchangeType -output ikeexchange_type_strings.go; DO NOT EDIT

package protocol

import "fmt"

const _IkeExchangeType_name = "IKE_SA_INITIKE_AUTHCREATE_CHILD_SAINFORMATIONALIKE_SESSION_RESUMEGSA_AUTHGSA_REGISTRATIONGSA_REKEY"

var _IkeExchangeType_index = [...]uint8{0, 11, 19, 34, 47, 65, 73, 89, 98}

func (i IkeExchangeType) String() string {
	i -= 34
	if i+1 >= IkeExchangeType(len(_IkeExchangeType_index)) {
		return fmt.Sprintf("IkeExchangeType(%d)", i+34)
	}
	return _IkeExchangeType_name[_IkeExchangeType_index[i]:_IkeExchangeType_index[i+1]]
}
