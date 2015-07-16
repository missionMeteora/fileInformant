package file

import (
	"os"
	"sync"
	"time"

	"github.com/missionMeteora/mandrill"
	"github.com/missionMeteora/twilio"

	"github.com/missionMeteora/fileInformant/internal/config"
)

func New(name, loc, interval string, subs []config.Subscriber, ec *mandrill.Client, tc *twilio.Client) (*File, error) {
	d, err := time.ParseDuration(interval)
	if err != nil {
		return nil, err
	}

	f := File{
		ec: ec,
		tc: tc,

		name:     name,
		loc:      loc,
		interval: d,
		subs:     subs,
	}

	// Called to set initial values of exists, size, lastModified for file
	f.wasModified()
	f.setInterval()

	return &f, nil
}

type File struct {
	mux sync.RWMutex

	ec *mandrill.Client
	tc *twilio.Client

	// Name of server
	name     string
	loc      string
	interval time.Duration
	subs     []config.Subscriber

	exists       bool
	size         int64
	lastModified time.Time
}

// Returns boolean to determine if file has been changed since the last time
//	wasModified was called.
//
//	If an error occurs while os.Stat is called on a file,
// 	error is passed to f.existsCheck and value of this fn is returned.
//
//	If file information is available, return value of f.statsCheck fn.
func (f *File) wasModified() bool {
	fi, err := os.Stat(f.loc)
	if err != nil {
		return f.existCheck(err)
	}

	return f.statsCheck(fi)
}

// If provided error == os.ErrNotExist and f.exists is set to true
//		Set f.exists to false. Return true (indicating a change has occurred)
// Else, return false (indicating that f.exists was already set to false)
func (f *File) existCheck(err error) bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	if err == os.ErrNotExist && f.exists {
		f.exists = false
		return true
	}

	return false
}

// If fileInfo.Size does not equal f.size
//		Update f.size to new value, set return boolean to true
// If fileInfo.ModTime does not equal f.lastModified
//		Update f.lastModified to new value, set return boolean to true
func (f *File) statsCheck(fi os.FileInfo) (m bool) {
	f.mux.Lock()

	if s := fi.Size(); s != f.size {
		m = true
		f.size = s
	}

	if l := fi.ModTime(); l != f.lastModified {
		m = true
		f.lastModified = l
	}

	f.mux.Unlock()
	return
}

func (f *File) setInterval() {
	tkr := time.NewTicker(f.interval)

	go func() {
		for _ = range tkr.C {
			if !f.wasModified() {
				f.notify()
			}
		}
	}()
}

func (f *File) notify() {
	for _, s := range f.subs {
		if len(s.Email) > 0 {
			f.ec.SendMessage(getEmailMessage(f.loc, f.name), emailSubject, s.Email, s.Name, emailTags)
		}

		if len(s.Phone) > 0 {
			f.tc.Send(s.Phone, getSmsMessage(f.loc, f.name))
		}
	}
}
