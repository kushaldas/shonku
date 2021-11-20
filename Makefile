all:
	GO111MODULE=auto go build -o shonku.bin scrdkd.go posts.go bindata.go

clean:
	rm -f .scrdkd.db
	rm -rf output/*.html
	rm -f shonku.bin
