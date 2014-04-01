all:
	go build -o shonku.bin scrdkd.go posts.go

clean:
	rm .scrdkd.db
	rm -rf output/*.html