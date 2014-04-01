all:
	go build -o mark.bin mark.go
	go build -o logit.bin scrdkd.go

clean:
	rm .scrdkd.db
	rm -rf output/*.html