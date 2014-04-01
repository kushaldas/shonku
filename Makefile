all:
	go build -o shonku.bin scrdkd.go

clean:
	rm .scrdkd.db
	rm -rf output/*.html