
build: pre
	@go build -o litc

pre:
	@./litc-bootstrap main.go.md main.go

bootstrap:
	@echo "bootstrap"
	@go build -o litc-bootstrap

test: build
	@mkdir -p tmp
	@./litc samples/sample.go.md tmp/main.go
	@go run tmp/main.go
	@rm -rf tmp

clean:
	@rm -f litc
	@rm -f litc-bootstrap
