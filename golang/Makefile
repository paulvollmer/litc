VERSION=0.2.5

build: pre
	@echo "build litc v${VERSION}"
	@go build -ldflags "-X main.version=${VERSION}" -o litc

pre:
	@echo "pre-processing main.go.md"
	@./litc-bootstrap main.go.md main.go

bootstrap:
	@echo "bootstrap litc v${VERSION}..."
	@go build -ldflags "-X main.version=${VERSION}" -o litc-bootstrap
	@./litc-bootstrap version

test: build
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
