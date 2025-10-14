.PHONY: build
build: clean
	go build -gcflags="all=-N -l" -o exe ./main.go

.PHONY: clean
clean:
	rm -f exe
