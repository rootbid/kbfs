package libkbfs

import (
	"testing"

	"github.com/keybase/go-codec/codec"
)

type dirEntryFuture struct {
	DirEntry
	extra
}

func (cof dirEntryFuture) toCurrent() DirEntry {
	return cof.DirEntry
}

func (cof dirEntryFuture) toCurrentStruct() currentStruct {
	return cof.toCurrent()
}

func makeFakeDirEntryFuture(t *testing.T) dirEntryFuture {
	cof := dirEntryFuture{
		DirEntry{
			makeFakeBlockInfo(t),
			EntryInfo{
				Dir,
				100,
				"fake sym path",
				101,
				102,
			},
			codec.UnknownFieldSetHandler{},
		},
		makeExtraOrBust("dirEntry", t),
	}
	return cof
}

func TestDirEntryUnknownFields(t *testing.T) {
	testStructUnknownFields(t, makeFakeDirEntryFuture(t))
}
