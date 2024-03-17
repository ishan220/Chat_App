package model

type Chat struct {
	ID       string `json:id`
	FromUser string `json:from_user`
	ToUser   string `json:to_user`
	Message  string `json:message`
	Time     int64  `json:timestamp`
}

type ContactList struct {
	Username     string `json:"username"`
	LastActivity int64  `json:"last_activity"`
	MemberName   string `json:"member"`
}
