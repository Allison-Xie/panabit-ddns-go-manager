.PHONY: all

GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean

DIST_DIR=./build/dist
PACKAGE=./build/panabit-ddns-go-manager.tar.gz


all: clean build package

clean:
	rm -rf $(DIST_DIR)

build: build-ctl build-cgi build-hooks

build-ctl:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(DIST_DIR)/appctl -v ./cmd
build-cgi:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(DIST_DIR)/web/cgi/webmain -v ./cmd/cgi
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(DIST_DIR)/web/cgi/api -v ./cmd/cgi/api
build-hooks:
	GOOS=linux GOARCH=arm64 $(GOBUILD) -o $(DIST_DIR)/afterinstall -v ./cmd/hooks/postinstall

DDNSGO=./static/bin/ddns-go
DDNSGO_PATH=./static/bin
DDNSGO_TARBALL=./static/bin/ddns-go.tar.gz
DDNSGO_URL=https://github.com/jeessy2/ddns-go/releases/download/v5.6.6/ddns-go_5.6.6_linux_arm64.tar.gz

package: $(DDNSGO)
	cp -r ./static/* $(DIST_DIR)
	chmod +x $(DIST_DIR)/appctrl
	tar -czvf $(PACKAGE) -C $(DIST_DIR) --exclude='.gitkeep' .

$(DDNSGO):
	wget -O $(DDNSGO_TARBALL) $(DDNSGO_URL)
	tar -xzvf $(DDNSGO_TARBALL) -C $(DDNSGO_PATH)
	rm $(DDNSGO_TARBALL)
