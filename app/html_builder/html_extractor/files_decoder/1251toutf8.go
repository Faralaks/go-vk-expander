// Package files_decoder
// Coded by Faralaks https://github.com/Faralaks
// This package provides functionality  which allows u convert files from one encoding to another.
package files_decoder

type FileListGetter interface {
	Lock()
	Unlock()
	GetFileList() []string
}

type DecoderWin1251ToUTF8 struct {
	Dialogs []FileListGetter
}

func NewDecoderWin1251ToUTF8(dialogs ...FileListGetter) *DecoderWin1251ToUTF8 {
	return &DecoderWin1251ToUTF8{Dialogs: dialogs}
}
