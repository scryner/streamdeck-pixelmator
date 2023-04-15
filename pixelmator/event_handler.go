package pixelmator

import (
	"encoding/json"
	"github.com/scryner/streamdeck-go-sdk/sdk"
	"log"
	"strconv"
	"strings"
)

func WillAppear(px *Pixelmator) sdk.HandlerFunc {
	return func(plugin *sdk.Plugin, event interface{}) {
		p := event.(*sdk.WillAppearEvent)

		// parse adjustment
		adj, err := parseAdjustment(p.Action)
		if err != nil {
			log.Printf("failed to parse adjustment from '%s': %v", p.Event, err)
			return
		}

		// register adjustment
		if err := px.registerAdjustment(adj); err != nil {
			log.Printf("failed to register adjustment '%s': %v", adj.getTerm().osascriptTerm, err)
		}
	}
}

func WillDisappear(px *Pixelmator) sdk.HandlerFunc {
	return func(plugin *sdk.Plugin, event interface{}) {
		p := event.(*sdk.WillAppearEvent)

		// parse adjustment
		adj, err := parseAdjustment(p.Action)
		if err != nil {
			log.Printf("failed to parse adjustment from '%s': %v", p.Event, err)
			return
		}

		// register adjustment
		if err := px.deregisterAdjustment(adj); err != nil {
			log.Printf("failed to de-register adjustment '%s': %v", adj.getTerm().osascriptTerm, err)
		}
	}
}

func TouchTap(px *Pixelmator) sdk.HandlerFunc {
	return func(plugin *sdk.Plugin, event interface{}) {

	}
}

func DialUp(px *Pixelmator) sdk.HandlerFunc {
	return func(plugin *sdk.Plugin, event interface{}) {
		p := event.(*sdk.DialEvent)
		b, _ := json.MarshalIndent(p, "", "  ")

		log.Println("------", string(b))
	}
}

func DialRotate(px *Pixelmator) sdk.HandlerFunc {
	return func(plugin *sdk.Plugin, event interface{}) {
		p := event.(*sdk.DialRotateEvent)

		// parse adjustment
		adj, err := parseAdjustment(p.Action)
		if err != nil {
			log.Printf("failed to parse adjustment from '%s': %v", p.Event, err)
			return
		}

		if err := px.adjust(adj, p.Payload.Ticks); err != nil {
			log.Printf("failed to adjust for %s: %v", adj.getTerm().osascriptTerm, err)
		}
	}
}

func DeviceConnect(px *Pixelmator) sdk.HandlerFunc {
	return func(plugin *sdk.Plugin, event interface{}) {

	}
}

func DeviceDisconnect(px *Pixelmator) sdk.HandlerFunc {
	return func(plugin *sdk.Plugin, event interface{}) {

	}
}

func parseAdjustment(s string) (ColorAdjustment, error) {
	lastI := strings.LastIndex(s, ".")
	v := s[lastI+1:]
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return Noop, err
	}

	return ColorAdjustment(int(i)), nil
}
