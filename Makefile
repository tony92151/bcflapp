build:
	go build ./x/bcflapp/
	go build -o bcflappd ./cmd/bcflappd/main.go ./cmd/bcflappd/genaccounts.go
	go build -o bcflappcli ./cmd/bcflappcli/main.go
	@echo "\nBuild success."
	@echo "\n Use ./bcflapp [command]"

install:
	# go install ./cmd/tttd
	# go install ./cmd/tttcli
