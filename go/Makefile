
build:
	@./litc-bootstrap main.md main.go
	@go build -o litc

init:
	@go build -o litc-bootstrap

test: build
	@mkdir -p tmp
	./litc samples/sample.md tmp/main.go
	@go run tmp/main.go
	@rm -rf tmp

clean:
	rm -f litc
	rm -f litc-bootstrap
