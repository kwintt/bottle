package bot

import (
	"fmt"
	"github.com/gempir/go-twitch-irc"
	"os"
)

func Start() {
	client := twitch.NewClient(
		os.Getenv("TWITCH_NAME"),
		os.Getenv("TWITCH_OAUTH"),
	)
	client.OnNewMessage(func(_ string, user twitch.User, message twitch.Message) {
		// https://en.wikipedia.org/wiki/ANSI_escape_code#8-bit
		r, g, b := Color(user.Color)
		fmt.Printf("\u001b[48;2;%d;%d;%dm%s\u001b[0m: %s\n", r, g, b, user.Username, message.Text)
	})
	client.Join("kwintt")
	if err := client.Connect(); err != nil {
		panic(err)
	}
}
