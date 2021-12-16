all:
	go build -buildmode=c-shared -o out_prettyslack.so

clean:
	rm -rf *.so *.h *~