run:
	wails dev
.PHONY: build
build:
	wails build

build-windows:
	GOOS=windows wails build -skipbindings