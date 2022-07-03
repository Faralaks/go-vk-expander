package files_decoder

import (
	"context"
	. "github.com/faralaks/go-vk-expander/app/html_builder/dialog_files"
	"testing"
	"time"
)

func TestDecoderWin1251ToUTF8_Run(t *testing.T) {
	d := NewDecoderWin1251ToUTF8()
	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan *Dialog)
	go d.Run(ctx, c)
	time.Sleep(500 * time.Millisecond)
	cancel()
}
