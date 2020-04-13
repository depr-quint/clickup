package clickup

import (
	"log"
	"os"
	"testing"
)

var testClient *Client

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	secret, ok := os.LookupEnv("CLICKUP_TOKEN")
	if !ok {
		log.Fatal("could not get ClickUp API token")
	}
	testClient = NewClient(nil, &secret)
}
