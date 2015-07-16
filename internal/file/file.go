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

func (f *File) wasModified() bool {
	fi, err := os.Stat(f.loc)
	if err != nil {
		return f.existCheck(err)
	}

	return f.statsCheck(fi)
}

func (f *File) existCheck(err error) bool {
	if err == os.ErrNotExist {
		if f.exists {
			f.mux.Lock()
			f.exists = false
			f.mux.Unlock()
			return true
		}
	}

	return false
}

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
