package notifiers

import (
	"fmt"

	"github.com/missionMeteora/fileInformant/internal/config"
	"github.com/missionMeteora/mandrill"
)

const (
	subject      = "File informant notification"
	mandrillTmpl = `<h2>Oh no!</h2>
<p>It appears that %s on %s has not been updated within the requested interval.</p>
<p>Please look into this to ensure nothing is amiss!</p>`
)

var (
	tags = []string{"notifications"}
)

func newMandrill(c *config.Mandrill) Mandrill {
	return Mandrill{
		clnt: mandrill.New(c.Key, c.SubAccount, c.FromEmail, c.FromName),
	}
}

type Mandrill struct {
	clnt *mandrill.Client
}

func (m Mandrill) Send(subs []config.Subscriber, msg string) {
	for _, s := range subs {
		if len(s.Email) > 0 {
			m.clnt.SendMessage(msg, subject, s.Email, s.Name, tags)
		}
	}
}

func (m Mandrill) GetMessage(a, b string) string {
	return fmt.Sprintf(mandrillTmpl, a, b)
}
