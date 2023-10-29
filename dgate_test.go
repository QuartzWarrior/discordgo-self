package dgate

import (
	"fmt"
	"os"
	"testing"

	"github.com/luminaldev/dgate/types"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Test_Main(t *testing.T) {
	token, err := os.ReadFile("test_token")
	handleErr(err)

	c := NewClient(string(token))

	c.AddHandler(types.ReadyEventHandler, func(e *types.ReadyEventData) {
		fmt.Println(e)
	})
	c.AddHandler(types.MessageCreateEventHandler, func(e *types.MessageEventData) {
		fmt.Println(e)
	})
	c.AddHandler(types.MessageUpdateEventHandler, func(e *types.MessageEventData) {
		fmt.Println(e)
	})

	handleErr(c.Connect())
}
