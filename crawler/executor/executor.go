package executor

type LinearExecutor[P any, M any] struct {
	Producer  func() []P
	Processor func(item P) M
	Consumer  func(item M)
}

func (exe *LinearExecutor[P, M]) Perform() {
	for _, produced := range exe.Producer() {
		intermediate := exe.Processor(produced)
		exe.Consumer(intermediate)
	}
}
