package genl

import (
	"errors"
	"syscall"

	"github.com/khirono/go-nl"
)

const (
	CTRL_NAME = "nlctrl"
)

func GetFamily(c *nl.Client, name string) (*Family, error) {
	req := nl.NewRequest(ID_CTRL, 0)
	err := req.Append(Header{
		Cmd:     CTRL_CMD_GETFAMILY,
		Version: CTRL_VERSION,
	})
	if err != nil {
		return nil, err
	}
	err = req.Append(&nl.Attr{
		Type:  CTRL_ATTR_FAMILY_NAME,
		Value: nl.AttrString(name),
	})
	if err != nil {
		return nil, err
	}
	rsps, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if len(rsps) < 1 {
		return nil, errors.New("not found")
	}
	f, err := DecodeFamily(rsps[0].Body[SizeofHeader:])
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GetFamilyAll(c *nl.Client) ([]Family, error) {
	flags := syscall.NLM_F_DUMP
	req := nl.NewRequest(ID_CTRL, flags)
	err := req.Append(Header{
		Cmd:     CTRL_CMD_GETFAMILY,
		Version: CTRL_VERSION,
	})
	if err != nil {
		return nil, err
	}
	rsps, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var fs []Family
	for _, rsp := range rsps {
		f, err := DecodeFamily(rsp.Body[SizeofHeader:])
		if err != nil {
			return nil, err
		}
		fs = append(fs, *f)
	}
	return fs, nil
}
