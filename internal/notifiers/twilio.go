package notifiers

import (
	"fmt"

	"github.com/missionMeteora/fileInformant/internal/config"
	"github.com/missionMeteora/twilio"
)

const twilioTmpl = `Oh no! It appears that %s on %s has not been updated within the requested interval. Please look into this to ensure nothing is amiss!`

func newTwilio(c *config.Twilio) Twilio {
	return Twilio{
		clnt: twilio.New(c.Key, c.Token, c.FromPhone),
	}
}

type Twilio struct {
	clnt *twilio.Client
}

func (t Twilio) Send(subs []config.Subscriber, msg string) {
	for _, s := range subs {
		if len(s.Phone) > 0 {
			t.clnt.Send(s.Phone, msg)
		}
	}
}

func (t Twilio) GetMessage(a, b string) string {
	return fmt.Sprintf(twilioTmpl, a, b)
}
