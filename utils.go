package assistentby

import (
	"log"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func replaceId(url string, id string) string {
	return strings.Replace(url, ":id", id, -1)
}
