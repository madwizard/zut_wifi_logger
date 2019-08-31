intel:
	go build -o build/scanner config.go scanner.go wifi.go db.go gps.go
	go build -o build/client server.go config.go db.go

arm:
	env GOOS=linux GOARCH=arm GOARM=5 go build -o build/scanner_arm config.go scanner.go wifi.go db.go gps.go
	env GOOS=linux GOARCH=arm GOARM=5 go build -o build/client_arm server.go config.go db.go
