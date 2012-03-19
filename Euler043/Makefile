results: main
	\time -f "%e" -o time ./main > results 2> stderr

main: main.go
	gofmt -w=true *.go && go build *.go

clean:
	rm main results
