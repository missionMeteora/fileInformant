package file

import (
	"fmt"
)

const (
	tmpl = `<h2>Oh no!</h2>
<p>It appears the %s has not been updated within the requested interval.</p>
<p>Please look into this to ensure nothing is amiss!</p>`
	emailSubject = "File informant notification"
)

var (
	emailTags = []string{"notifications"}
)

func getMessage(m string) string {
	return fmt.Sprintf(tmpl, m)
}
