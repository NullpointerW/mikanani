package test

import (
	"fmt"
	"strings"
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

func TestGetViaPath(t *testing.T) {
	ts, err := TORR.GetViaPath("C:/Users/W/Downloads")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	fmt.Printf("%#+v", ts[0])
}

func TestStr(t *testing.T) {
	str := "jook     lol"
	fs := strings.Fields(str)
	for _, s := range fs {
		fmt.Println(s)
	}
}
