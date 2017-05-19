package ctrie

import "github.com/badgerodon/goreify/generics"

type (
	inode struct {
		cmain *cnode
		smain *snode
	}
	cnode struct {
		bmp   int64
		children [256]
	}
	snode struct {
		key   generics.T1
		value generics.T2
		tomb  bool
	}
	CTrie struct {
		root *inode
	}
)

func insert(k, v []byte) {

}
