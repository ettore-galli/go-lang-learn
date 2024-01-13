package streams

type StreamSource[S any] struct {
	Source func() chan S
}

type StreamNode[S any] interface {
	Pipe(streamSource StreamSource[S])
}

func produceStrings() *chan string {
	data := make(chan string)
	go func() {
		data <- "aaa"
		data <- "bbb"
		data <- "ccc"
		close(data)
	}()
	return &data
}

type StringSource struct {
	Source func()
}
