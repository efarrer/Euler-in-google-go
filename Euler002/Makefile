results: main
	./main > results

main: main.6
	6l -o main *.6

main.6: main.go
	gofmt -w=true *.go && 6g *.go

clean:
	rm main *.6 results
