.PHONY: run build windows linux macos server
FILES=lib/helpers.go lib/copy_dir.go lib/structs.go lib/main.go

run r:
	go run ${FILES}

build b:
	make macos
	make windows
	make linux
	rm -rf bin
	mkdir -p bin
	mv pangolin.darwin bin/
	mv pangolin.exe bin/
	mv pangolin bin/

macos m:
	GOOS=darwin GOARCH=amd64 go build -o pangolin.darwin ${FILES}

windows w:
	GOOS=windows GOARCH=amd64 go build -o pangolin.exe ${FILES}

linux l:
	GOOS=linux GOARCH=amd64 go build -o pangolin ${FILES}

server s:
	cd _build && python3 -m http.server
