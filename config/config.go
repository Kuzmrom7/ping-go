package config

type Configurations struct {
	Bot     *Bot
	Url     string
}

type Bot struct {
	Token string
	Chat  int64
}
