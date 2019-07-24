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
BINARYDIR=bin
BINARYNAME=badgerhole
BINARYUNIX=$(BINARYNAME)_unix

all: test clean deps gbuild

gbuild:
	mkdir -p $(BUILDWORKDIR)/$(BINARYDIR)
	mkdir -p $(BUILDWORKDIR)/$(LOGDIR)
	cp -Rp $(WEBDIR) $(BUILDWORKDIR)/
	cp -Rp $(CONFIGDIR) $(BUILDWORKDIR)/
	$(GOBUILD) -o $(BUILDWORKDIR)/$(BINARYDIR)/$(BINARYNAME) -v $(SRCDIR)/$(BINARYNAME).go

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BUILDWORKDIR)/$(BINARYDIR)
	rm -rf $(BUILDWORKDIR)/$(LOGDIR)
	rm -rf $(BUILDWORKDIR)/$(WEBDIR)
	rm -rf $(BUILDWORKDIR)/$(CONFIGDIR)

run: gbuild
	./$(BUILDWORKDIR)/$(BINARYDIR)/$(BINARYNAME)

deps:
	$(GOGET) github.com/google/uuid
	$(GOGET) github.com/spf13/pflag
	$(GOGET) github.com/spf13/viper

# cross compile
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARYUNIX) -v