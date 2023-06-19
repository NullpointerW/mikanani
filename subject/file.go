package subject

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	CFG "github.com/NullpointerW/mikanani/conf"
	"github.com/NullpointerW/mikanani/download/rss"
	TORR "github.com/NullpointerW/mikanani/download/torrent"
	"github.com/NullpointerW/mikanani/errs"
	"github.com/NullpointerW/mikanani/util"

	// "strconv"
	"strings"
)

var HOME string = CFG.Env.SubjPath

func Scan() {
	home := trimPath(HOME)
	if fs, err := os.ReadDir(home); err == nil {
		for _, f := range fs {
			if f.IsDir() {
				log.Println("scan:found folder:" + home + string(os.PathSeparator) + f.Name())
				if jsraw, err := os.ReadFile(home + `/` + f.Name() + `/` + jsonfileName); err == nil {
					var s Subject
					err := json.Unmarshal(jsraw, &s)
					if err != nil {
						log.Println(err)
					}
					s.runtimeInit(true)
				} else {
					log.Println(err)
				}
			}
		}
	} else {
		log.Println(err)
	}
}

// Initialize the content library in OS file for the `subject` in the memory.
// Path can be used to monitor the download status of resources
// apply to RSS and Torrent type.
// If initialization is successful, write the path to Subject.Path.
func initFolder(subject *Subject) (err error) {
	var folderPath string

	folderPath = trimPath(HOME)
	sd, err := util.ParseShortTime(subject.StartTime)
	if err != nil {
		return err
	}
	folderPath += "/" + "[" + sd + "]" + subject.Name

	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return err
	}

	if ap, err := filepath.Abs(folderPath); err == nil {
		subject.Path = ap
	}

	return
}

func rmFolder(s *Subject) error {
	return os.RemoveAll(s.Path)
}

func (s *Subject) writeJson() (err error) {
	b, _ := json.Marshal(*s)
	fldrp := s.Path
	err = os.WriteFile(fldrp+"/"+jsonfileName, b, os.ModePerm)
	return err
}

func (s *Subject) RmRes() error {
	wrap := errs.ErrWrapper{}
	if s.ResourceTyp == RSS {
		wrap.Handle(func() error {
			return rss.RmRss(s.RssPath())
		})
	}
	wrap.Handle(func() error {
		return TORR.DelTorrs(s.Path)
	})
	wrap.Handle(func() error {
		return TORR.DelTag(s.QbtTag())
	})
	return wrap.Error()
}

func trimPath(n string) string {
	return strings.TrimSuffix(strings.TrimSuffix(CFG.Env.SubjPath, "\\"), "/")
}
