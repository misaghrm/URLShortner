package protocoltype

type ProtocolType int64

const (
	Undefined ProtocolType = iota
	HTTP
	HTTPS
	FTP
	FTPS
	TCP
)
