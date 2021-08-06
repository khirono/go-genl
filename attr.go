package genl

import (
	"github.com/khirono/go-nl"
)

const (
	CTRL_ATTR_UNSPEC = iota
	CTRL_ATTR_FAMILY_ID
	CTRL_ATTR_FAMILY_NAME
	CTRL_ATTR_VERSION
	CTRL_ATTR_HDRSIZE
	CTRL_ATTR_MAXATTR
	CTRL_ATTR_OPS
	CTRL_ATTR_MCAST_GROUPS
	CTRL_ATTR_POLICY
	CTRL_ATTR_OP_POLICY
	CTRL_ATTR_OP
)

const (
	CTRL_ATTR_OP_UNSPEC = iota
	CTRL_ATTR_OP_ID
	CTRL_ATTR_OP_FLAGS
)

const (
	CTRL_ATTR_MCAST_GRP_UNSPEC = iota
	CTRL_ATTR_MCAST_GRP_NAME
	CTRL_ATTR_MCAST_GRP_ID
)

const (
	CTRL_ATTR_POLICY_UNSPEC = iota
	CTRL_ATTR_POLICY_DO
	CTRL_ATTR_POLICY_DUMP
)

type Op struct {
	ID    uint32
	Flags uint32
}

func DecodeOp(b []byte) (Op, error) {
	var op Op
	for len(b) > 0 {
		hdr, n, err := nl.DecodeAttrHdr(b)
		if err != nil {
			return op, err
		}
		switch hdr.MaskedType() {
		case CTRL_ATTR_OP_ID:
			op.ID = native.Uint32(b[n:])
		case CTRL_ATTR_OP_FLAGS:
			op.Flags = native.Uint32(b[n:])
		}
		b = b[hdr.Len.Align():]
	}
	return op, nil
}

func DecodeOps(b []byte) ([]Op, error) {
	var ops []Op
	for len(b) > 0 {
		hdr, n, err := nl.DecodeAttrHdr(b)
		if err != nil {
			return ops, err
		}
		op, err := DecodeOp(b[n:])
		if err != nil {
			return ops, err
		}
		ops = append(ops, op)
		b = b[hdr.Len.Align():]
	}
	return ops, nil
}

type MulticastGroup struct {
	ID   uint32
	Name string
}

func DecodeMulticastGroup(b []byte) (MulticastGroup, error) {
	var g MulticastGroup
	for len(b) > 0 {
		hdr, n, err := nl.DecodeAttrHdr(b)
		if err != nil {
			return g, err
		}
		switch hdr.MaskedType() {
		case CTRL_ATTR_MCAST_GRP_NAME:
			g.Name, _, _ = nl.DecodeAttrString(b[n:])
		case CTRL_ATTR_MCAST_GRP_ID:
			g.ID = native.Uint32(b[n:])
		}
		b = b[hdr.Len.Align():]
	}
	return g, nil
}

func DecodeMulticastGroups(b []byte) ([]MulticastGroup, error) {
	var gs []MulticastGroup
	for len(b) > 0 {
		hdr, n, err := nl.DecodeAttrHdr(b)
		if err != nil {
			return gs, err
		}
		g, err := DecodeMulticastGroup(b[n:])
		if err != nil {
			return gs, err
		}
		gs = append(gs, g)
		b = b[hdr.Len.Align():]
	}
	return gs, nil
}

type Family struct {
	ID      uint16
	HdrSize uint32
	Name    string
	Version uint32
	MaxAttr uint32
	Ops     []Op
	Groups  []MulticastGroup
}

func DecodeFamily(b []byte) (*Family, error) {
	f := new(Family)
	for len(b) > 0 {
		hdr, n, err := nl.DecodeAttrHdr(b)
		if err != nil {
			return nil, err
		}
		switch hdr.MaskedType() {
		case CTRL_ATTR_FAMILY_NAME:
			f.Name, _, _ = nl.DecodeAttrString(b[n:])
		case CTRL_ATTR_FAMILY_ID:
			f.ID = native.Uint16(b[n:])
		case CTRL_ATTR_VERSION:
			f.Version = native.Uint32(b[n:])
		case CTRL_ATTR_HDRSIZE:
			f.HdrSize = native.Uint32(b[n:])
		case CTRL_ATTR_MAXATTR:
			f.MaxAttr = native.Uint32(b[n:])
		case CTRL_ATTR_OPS:
			f.Ops, err = DecodeOps(b[n:])
			if err != nil {
				return nil, err
			}
		case CTRL_ATTR_MCAST_GROUPS:
			f.Groups, err = DecodeMulticastGroups(b[n:])
			if err != nil {
				return nil, err
			}
		}
		b = b[hdr.Len.Align():]
	}
	return f, nil
}
