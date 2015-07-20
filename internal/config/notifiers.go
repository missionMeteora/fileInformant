package config

type Notifiers struct {
	Mandrill *Mandrill `json:"mandrill"`
	Twilio   *Twilio   `json:"twilio"`
}

type Mandrill struct {
	Key        string `json:"key"`
	SubAccount string `json:"subAccount"`
	FromEmail  string `json:"fromEmail"`
	FromName   string `json:"fromName"`
}

type Twilio struct {
	Key       string `json:"key"`
	Token     string `json:"token"`
	FromPhone string `json:"fromPhone"`
}
