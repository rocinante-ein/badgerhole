APPNAME=badgerhole
APPVERSION=0.1.0

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SRCDIR=cmd
BUILDWORKDIR=work
WEBDIR=web
CONFIGDIR=configs
LOGDIR=logs
RELEASEDIR=release
BINARYDIR=bin


all: test clean deps gbuild

gbuild:
	mkdir -p $(BUILDWORKDIR)/$(APPNAME)/$(BINARYDIR)
	mkdir -p $(BUILDWORKDIR)/$(APPNAME)/$(LOGDIR)
	cp -Rp $(WEBDIR) $(BUILDWORKDIR)/$(APPNAME)/
	cp -Rp $(CONFIGDIR) $(BUILDWORKDIR)/$(APPNAME)/
	$(GOBUILD) -o $(BUILDWORKDIR)/$(APPNAME)/$(BINARYDIR)/$(APPNAME) -v $(SRCDIR)/$(APPNAME).go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILDWORKDIR)/*

run: gbuild
	./$(BUILDWORKDIR)/$(APPNAME)/$(BINARYDIR)/$(APPNAME)

deps:
	$(GOGET) github.com/google/uuid
	$(GOGET) github.com/spf13/pflag
	$(GOGET) github.com/spf13/viper

# make install ... badgerhole admin install server
install: clean deps gbuild
	mkdir -p /opt/$(APPNAME)
	cp -Rp $(BUILDWORKDIR)/$(APPNAME) /opt/$(APPNAME)/

# release source files
release-src: clean deps gbuild
	rm -f $(RELEASEDIR)/$(APPNAME)-$(APPVERSION).tar.gz
	tar -zcvf $(RELEASEDIR)/$(APPNAME)-$(APPVERSION).tar.gz -C $(BUILDWORKDIR)/ .

# cross compile
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARYUNIX) -v