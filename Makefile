PLUGIN_UUID   = com.scryner.pixelmator
PLUGIN_BINARY = pixelmator
PLUGIN_DIR    = $(PLUGIN_UUID).sdPlugin

all: build

.PHONY: build
build:
	mkdir -p $@/$(PLUGIN_DIR)
	GOARCH=amd64 CGO_ENABLED=1 go build -o $@/$(PLUGIN_BINARY)-amd64 main.go
	GOARCH=arm64 CGO_ENABLED=1 go build -o $@/$(PLUGIN_BINARY)-arm64 main.go
	lipo -create -output $@/$(PLUGIN_DIR)/$(PLUGIN_BINARY) $@/$(PLUGIN_BINARY)-amd64 $@/$(PLUGIN_BINARY)-arm64
	cp manifest.json $@/$(PLUGIN_DIR)
	cp -R assets/* $@/$(PLUGIN_DIR)
	cd $@ && zip -r $(PLUGIN_DIR).streamDeckPlugin $(PLUGIN_DIR)

install: build
	unzip -o $</$(PLUGIN_UUID).sdPlugin.streamDeckPlugin -d ~/Library/Application\ Support/com.elgato.StreamDeck/Plugins

uninstall:
	rm -rf ~/Library/Application\ Support/com.elgato.StreamDeck/Plugins/$(PLUGIN_UUID).sdPlugin

clean:
	rm -rf build