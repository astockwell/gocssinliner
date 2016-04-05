build:
	# Darwin
	cd cmd/gocssinliner && go build
	mv cmd/gocssinliner/gocssinliner cmd/gocssinliner/gocssinliner_darwin-amd64

	# Linux
	cd cmd/gocssinliner && env GOOS=linux GOARCH=amd64 go build
	mv cmd/gocssinliner/gocssinliner cmd/gocssinliner/gocssinliner_linux-amd64

	# Windows
	docker-machine start default
	docker-machine env default
	eval "$$(docker-machine env default)"
	cd cmd/gocssinliner && xgo --targets=windows/amd64 .
	docker-machine stop default

.PHONY: build
