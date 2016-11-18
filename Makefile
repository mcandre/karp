VERSION=0.0.1

.PHONY: build-ports clean clean-ports

tall: test

test:
	karp karp.gif

govet:
	go list ./... | grep -v vendor | xargs go vet -v

gofmt:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec gofmt -s -w {} \;

goimport:
	find . -path '*/vendor/*' -prune -o -name '*.go' -type f -exec goimports -w {} \;

lint: govet gofmt goimport

build-ports:
	sh -c "cd cmd/karp && mkdir -p bin/karp-$(VERSION)/linux/amd64 && env GOOS=linux GOARCH=amd64 go build -o bin/karp-$(VERSION)/linux/amd64/karp"
	sh -c "cd cmd/karp && mkdir -p bin/karp-$(VERSION)/linux/386 && env GOOS=linux GOARCH=386 go build -o bin/karp-$(VERSION)/linux/386/karp"
	sh -c "cd cmd/karp && mkdir -p bin/karp-$(VERSION)/darwin/amd64 && env GOOS=darwin GOARCH=amd64 go build -o bin/karp-$(VERSION)/darwin/amd64/karp"
	sh -c "cd cmd/karp && mkdir -p bin/karp-$(VERSION)/darwin/386 && env GOOS=darwin GOARCH=386 go build -o bin/karp-$(VERSION)/darwin/386/karp"
	sh -c "cd cmd/karp && mkdir -p bin/karp-$(VERSION)/windows/amd64 && env GOOS=windows GOARCH=amd64 go build -o bin/karp-$(VERSION)/windows/amd64/karp.exe"
	sh -c "cd cmd/karp && mkdir -p bin/karp-$(VERSION)/windows/386 && env GOOS=windows GOARCH=386 go build -o bin/karp-$(VERSION)/windows/386/karp.exe"
	sh -c "cd cmd/karp/bin && zip -r karp-$(VERSION).zip karp-$(VERSION)/"

clean: clean-ports

clean-ports:
	rm -rf cmd/karp/bin
