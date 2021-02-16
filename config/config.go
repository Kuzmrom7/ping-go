package config

type Configurations struct {
	Bot     *Bot
	Url     string
	Timeout uint64
}

type Bot struct {
	Token string
	Chat  int64
}
