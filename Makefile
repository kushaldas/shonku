all:
	go build -o shonku.bin scrdkd.go posts.go bindata.go

clean:
	rm .scrdkd.db
	rm -rf output/*.html