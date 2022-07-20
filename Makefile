PROJ_NAME=hl-api

all: build

build:
	go build -o $(PROJ_NAME) .

clean:
	rm -rf $(PROJ_NAME)

install:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp -f hl-api $(DESTDIR)$(PREFIX)/bin
	chmod 755 $(DESTDIR)$(PREFIX)/bin/$(PROJ_NAME)

uninstall:
	rm -f $(DESTDIR)$(PREFIX)/bin/$(PROJ_NAME)

.PHONY: all build clean install uninstall
