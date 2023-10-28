NAME := quail-gui
run:
	wails dev
.PHONY: build
build:
	wails build

build-darwin:
	GOOS=darwin wails build
build-windows:
	GOOS=windows wails build -skipbindings
build-linux:
	GOOS=linux wails build -skipbindings
build-run:
	wails build
	cd build/bin && start ${NAME}.app