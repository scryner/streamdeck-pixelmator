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
	supportedAdjustments := []pixelmator.ColorAdjustment{
		pixelmator.Exposure,
		pixelmator.Tint,
	}

	for _, adj := range supportedAdjustments {
		plugin.Handle(fmt.Sprintf("com.scryner.pixelmator.adjust.%d/%s", adj, sdk.EventTouchTap), pixelmator.TouchTap(px))
		plugin.Handle(fmt.Sprintf("com.scryner.pixelmator.adjust.%d/%s", adj, sdk.EventDialUp), pixelmator.DialUp(px))
		plugin.Handle(fmt.Sprintf("com.scryner.pixelmator.adjust.%d/%s", adj, sdk.EventDialRotate), pixelmator.DialRotate(px))
		plugin.Handle(fmt.Sprintf("com.scryner.pixelmator.adjust.%d/%s", adj, sdk.EventWillAppear), pixelmator.WillAppear(px))
		plugin.Handle(fmt.Sprintf("com.scryner.pixelmator.adjust.%d/%s", adj, sdk.EventWillDisappear), pixelmator.WillDisappear(px))
	}

	plugin.Handle(fmt.Sprintf("com.scryner.pixelmator/%s", sdk.EventDeviceDidConnect), pixelmator.DeviceConnect(px))
	plugin.Handle(fmt.Sprintf("com.scryner.pixelmator/%s", sdk.EventDeviceDidDisconnect), pixelmator.DeviceDisconnect(px))

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
