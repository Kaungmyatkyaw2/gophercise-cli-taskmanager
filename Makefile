.PHONY: cmd/install
cmd/install:
	@echo 'Building the cmd line tool...'
	go build ./
	@echo 'Moving up to binary directory...'
	sudo mv ./task /usr/bin