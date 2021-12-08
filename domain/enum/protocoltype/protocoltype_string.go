// Code generated by "stringer -type ProtocolType"; DO NOT EDIT.

package protocoltype

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Undefined-0]
	_ = x[HTTP-1]
	_ = x[HTTPS-2]
	_ = x[FTP-3]
	_ = x[FTPS-4]
	_ = x[TCP-5]
}

const _ProtocolType_name = "UndefinedHTTPHTTPSFTPFTPSTCP"

var _ProtocolType_index = [...]uint8{0, 9, 13, 18, 21, 25, 28}

func (i ProtocolType) String() string {
	if i < 0 || i >= ProtocolType(len(_ProtocolType_index)-1) {
		return "ProtocolType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ProtocolType_name[_ProtocolType_index[i]:_ProtocolType_index[i+1]]
}
