default: build

build:
	go build github.com/TheCreeper/UselessLocker/cmd/locker
	zip -r assets.zip assets/
	cat assets.zip >> locker
	zip -A locker

build-windows:
	GOOS=windows GOARCH=386 go build github.com/TheCreeper/UselessLocker/cmd/locker
	zip -r assets.zip assets/
	cat assets.zip >> locker.exe
	zip -A locker.exe