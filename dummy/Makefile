all: dummy.so

dummy.so:
	go build -x -buildmode=c-shared -o dummy.so

clean:
	go clean

.PHONY: all clean
