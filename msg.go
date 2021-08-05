package genl

const (
	CTRL_CMD_UNSPEC = iota
	CTRL_CMD_NEWFAMILY
	CTRL_CMD_DELFAMILY
	CTRL_CMD_GETFAMILY
	CTRL_CMD_NEWOPS
	CTRL_CMD_DELOPS
	CTRL_CMD_GETOPS
	CTRL_CMD_NEWMCAST_GRP
	CTRL_CMD_DELMCAST_GRP
	CTRL_CMD_GETMCAST_GRP // unused
	CTRL_CMD_GETPOLICY
)

const (
	CTRL_VERSION = 2
)

const (
	SizeofHeader = 4
)

type Header struct {
	Cmd     uint8
	Version uint8
}

func DecodeHeader(b []byte) (Header, error) {
	var h Header
	h.Cmd = b[0]
	h.Version = b[1]
	return h, nil
}

func (h Header) Len() int {
	return SizeofHeader
}

func (h Header) Encode(b []byte) (int, error) {
	b[0] = h.Cmd
	b[1] = h.Version
	return h.Len(), nil
}
