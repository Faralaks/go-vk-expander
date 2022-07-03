// Package files_decoder
// Coded by Faralaks https://github.com/Faralaks
// This package provides functionality  which allows u convert files from one encoding to another.
package files_decoder

import (
	"context"
	. "github.com/faralaks/go-vk-expander/app/html_builder/dialog_files"
)

type DecoderWin1251ToUTF8 struct {
}

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
