gev:
	templ generate && sudo wails dev

build-windows:
	@sudo rm -f build/bin/metrom.exe && \
	sudo wails build -platform windows/amd64

build-macos:
	@sudo rm -rf build/bin/metrom.app && \
	sudo wails build -platform darwin/arm64 && \
	cd build/bin && \
	sudo zip -r metrom-darwin-m1.zip metrom.app && \
	rm -rf metrom.app && \
	cd ../..

build: build-windows build-macos
