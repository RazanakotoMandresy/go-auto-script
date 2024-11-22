build:
	@go build -C app -o ../go-auto-script
run: build
	./go-auto-script