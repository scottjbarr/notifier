package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scottjbarr/notifier"
)

func main() {
	msg := ""

	flag.StringVar(&msg, "message", "", "Message to send")
	flag.Parse()

	if len(msg) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	url := os.Getenv("NOTIFIER_SLACK_URL")

	if len(url) == 0 {
		panic(fmt.Errorf("Env var NOTIFIER_SLACK_URL must be specified"))
	}

	n := notifier.NewSlackNotifier(url)

	b := []byte(msg)

	if err := n.Notify(b); err != nil {
		panic(err)
	}
}
