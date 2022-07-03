package dialog_files

import "sync"

// MsgFiles presents list of messages filenames from one dialog
type MsgFiles []string

// Dialog presents messages file list from one dialog amd following mutex
// implements file_decoder.FileListGetter
type Dialog struct {
	Decoded bool
	MsgFiles
	sync.Mutex
}

func NewDialog(msgFiles MsgFiles) *Dialog {
	return &Dialog{Decoded: false, MsgFiles: msgFiles}
}

// GetFileList returns list of msg filenames
func (d *Dialog) GetFileList() []string {
	return d.MsgFiles
}
