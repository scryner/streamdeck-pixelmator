package main

import (
	"fmt"
	"github.com/scryner/streamdeck-go-sdk/sdk"
	"github.com/scryner/streamdeck-pixelmator/pixelmator"
	"log"
	"os"
)

func main() {
	// make plugin
	plugin, err := sdk.NewPlugin()
	if err != nil {
		log.Printf("failed to make streamdeck plugin: %v", err)
		os.Exit(1)
	}

	// make pixelmator
	px := pixelmator.NewPixelmator()

	// registry handler to plugin
	plugin.Handle(fmt.Sprintf("com.scryner.pixelmator/%s", sdk.EventWillAppear), pixelmator.WillAppear(px))
	plugin.Handle(fmt.Sprintf("com.scryner.pixelmator/%s", sdk.EventWillAppear), pixelmator.WillDisappear(px))
	plugin.Handle(fmt.Sprintf("com.scryner.pixelmator/%s", sdk.EventTouchTap), pixelmator.TouchTap(px))
	plugin.Handle(fmt.Sprintf("com.scryner.pixelmator/%s", sdk.EventDialUp), pixelmator.DialUp(px))
	plugin.Handle(fmt.Sprintf("com.scryner.pixelmator/%s", sdk.EventDialRotate), pixelmator.DialRotate(px))

	// set logger
	f, err := os.OpenFile("/tmp/streamdeck-pixelmator.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("failed to open log file: %v", err)
		os.Exit(1)
	}

	defer f.Close()
	log.SetOutput(f)

	// run plugin
	if err := plugin.Run(); err != nil {
		log.Printf("plugin error: %v", err)
		os.Exit(1)
	}
}
