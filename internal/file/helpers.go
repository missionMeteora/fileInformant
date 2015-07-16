package file

import (
	"fmt"
)

const (
	emailSubject = "File informant notification"
	emailTmpl    = `<h2>Oh no!</h2>
<p>It appears that %s on %s has not been updated within the requested interval.</p>
<p>Please look into this to ensure nothing is amiss!</p>`
	smsTmpl = `Oh no! It appears that %s on %s has not been updated within the requested interval. Please look into this to ensure nothing is amiss!`
)

var (
	emailTags = []string{"notifications"}
)

func getEmailMessage(m, n string) string {
	return fmt.Sprintf(emailTmpl, m, n)
}

func getSmsMessage(m, n string) string {
	return fmt.Sprintf(smsTmpl, m, n)
}
