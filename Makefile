build:
	templ generate internal/view
	$(MAKE) build-js
	$(MAKE) build-css
	go build -o ./bin/app cmd/cat-story/main.go
run:
	$(MAKE) build
	./bin/app
build-js:
	npm --prefix ./web run build:js

build-css:
	npm --prefix ./web run build:css
