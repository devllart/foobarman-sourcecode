build:	
	GOOS=linux   GOARCH=386   go build -o ./bin/foobarman-386-linux         ./cmd/app
	GOOS=linux   GOARCH=amd64 go build -o ./bin/foobarman-amd64-linux       ./cmd/app
	GOOS=linux   GOARCH=arm   go build -o ./bin/foobarman-arm-linux         ./cmd/app
	GOOS=linux   GOARCH=arm64 go build -o ./bin/foobarman-arm64-linux       ./cmd/app
	
	GOOS=windows GOARCH=386   go build -o ./bin/foobarman-386-windows.exe   ./cmd/app
	GOOS=windows GOARCH=amd64 go build -o ./bin/foobarman-amd64-windows.exe ./cmd/app
	GOOS=windows GOARCH=arm   go build -o ./bin/foobarman-arm-windows.exe   ./cmd/app
	GOOS=windows GOARCH=arm64 go build -o ./bin/foobarman-arm64-windows.exe ./cmd/app

buildonce:
	go build -o ./foobarman ./cmd/app

run:
	ENVIRONMENT=development go run ./cmd/app

readmeupdate:
	rm -rf bin/assets
	cp -r assets bin/
	rm -f bin/Readme.md
	cp Readme.md bin/Readme.md

buildfull:
	make readmeupdate
	make build
