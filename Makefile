all: g2z.so

g2z.so: g2z.go module.go agent.go cfuncs.go log.go log.h module.h zbxtypes.h 
	go build -x -buildmode=c-shared -o g2z.so

clean:
	rm -f g2z.so g2z.h

docker-build:
	docker build -t cavaliercoder/g2z .

docker-run:
	docker run --rm -it -v $(PWD):/usr/src/g2z -w /usr/src/g2z cavaliercoder/g2z

.PHONY: all clean docker-build docker-run
