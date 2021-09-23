package utils

import (
	"log"
	"os"
	"strings"

	"github.com/blinkbean/dingtalk"
)

func PushMessage(msgs []string) {

	dt := dingtalk.InitDingTalk([]string{os.Getenv("DINGTALK_TOKEN")}, "")

	err := dt.SendTextMessage(strings.Join(msgs, "\n"))
	if err != nil {
		log.Fatalf("Dingtalk send error: %v", err)
	}
}
