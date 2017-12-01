all:
	go build -buildmode=c-shared -o out_stdout.so .

fast:
	go build out_stdout.go

clean:
	rm -rf *.so *.h *~

test:
	go test 
