package genl

import (
	"syscall"
)

const (
	ID_CTRL = iota + syscall.NLMSG_MIN_TYPE
	ID_VFS_DQUOT
	ID_PMCRAID
	START_ALLOC
)

const (
	ID_MIN = syscall.NLMSG_MIN_TYPE
	ID_MAX = 1023
)
