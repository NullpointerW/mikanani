package test

import (
	"fmt"
	"testing"

	TORR "github.com/NullpointerW/mikanani/download/torrent"
)

func TestDL(t *testing.T) {
	h, err := TORR.Add("magnet:?xt=urn:btih:3522edcc5e979347bf1bc6a99cf12c15b5e66170&tr=http://open.acgtracker.com:1096/announce", "./dl", "mikan@subject-123")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(h)
}

func TestDLcompl(t *testing.T) {
	c, err := TORR.DLcompl("3522edcc5e979347bf1bc6a99cf12c15b5e66170")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Println(c)
}
