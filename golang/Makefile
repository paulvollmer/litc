
build: pre
	@echo "build litc"
	@go build -o litc

pre:
	@echo "pre-processing main.go.md"
	@./litc-bootstrap main.go.md main.go -remove
	go fmt main.go

bootstrap:
	@echo "bootstrap litc"
	@go build -o litc-bootstrap
	@./litc-bootstrap version

test-lit:
	@echo "test lit source"
	@cd lit && go test

test: build test-lit
	@./litc version
	@mkdir -p tmp
	@echo "build test sample"
	@./litc samples/sample.go.md tmp/main.go
	@echo "run test sample\n---"
	@go run tmp/main.go
	@echo "---"
	@rm -rf tmp

clean:
	@rm -f litc
	@rm -f litc-bootstrap
