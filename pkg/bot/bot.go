package bot

import (
	"github.com/nicklaw5/helix/v2"
)

type Bot struct {
	client *helix.Client
}

func (b *Bot) Init() {
	var err error
	b.client, err = helix.NewClient(&helix.Options{
		ClientID: "test",
	})

	if err != nil {

	}
}
