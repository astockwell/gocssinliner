build:
	cd cmd/gocssinliner && go build
	mv cmd/gocssinliner/gocssinliner cmd/gocssinliner/gocssinliner_darwin-amd64

	cd cmd/gocssinliner && env GOOS=linux GOARCH=amd64 go build
	mv cmd/gocssinliner/gocssinliner cmd/gocssinliner/gocssinliner_linux-amd64

.PHONY: deploy
