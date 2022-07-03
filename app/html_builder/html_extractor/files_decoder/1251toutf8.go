// Package files_decoder
// Coded by Faralaks https://github.com/Faralaks
// This package provides functionality  which allows u convert files from one encoding to another.
package files_decoder

import (
	"context"
	. "github.com/faralaks/go-vk-expander/app/html_builder/dialog_files"
)

// DecoderWin1251ToUTF8 Implementation of DecoderRunner which decodes from windows-1251 to UTF-8
type DecoderWin1251ToUTF8 struct {
}

//Run Decoder Runner
func (d DecoderWin1251ToUTF8) Run(ctx context.Context, dialogsChan chan *Dialog) {
	for {
		select {
		case _ = <-ctx.Done():
			println("done!")
			return
		case dialog := <-dialogsChan:
			println(dialog.GetFileList()[0])
		}
	}
}

func NewDecoderWin1251ToUTF8() *DecoderWin1251ToUTF8 {
	return &DecoderWin1251ToUTF8{}
}
