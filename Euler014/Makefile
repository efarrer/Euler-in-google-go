results: main
	./main > results

main: main.go
	gofmt -w=true *.go && go build *.go

clean:
	rm main results
