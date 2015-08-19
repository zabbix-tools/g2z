g2z.so: g2z.go module.h
	go build -x -buildmode=c-shared -o g2z.so

clean:
	rm -f g2z.so g2z.h

docker:
	docker build -t cavaliercoder/g2z .

.PHONY: clean docker