.DEFAULT_GOAL := win

win: wslexecquiet.go
	GOOS=windows GOARCH=amd64 go build -o wslexecquiet.exe wslexecquiet.go

wingui: wslexecquiet.go
	GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o wslexecquiet.exe wslexecquiet.go