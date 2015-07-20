package notifiers

import (
	"github.com/missionMeteora/fileInformant/internal/config"
)

func New(nfrs config.Notifiers) (n Notifiers) {
	if nfrs.Twilio != nil {
		n = append(n, newTwilio(nfrs.Twilio))
	}

	if nfrs.Mandrill != nil {
		n = append(n, newMandrill(nfrs.Mandrill))
	}

	return
}

type Notifiers []Notifier

type Notifier interface {
	Send(subs []config.Subscriber, msg string)
	GetMessage(a, b string) string
}
