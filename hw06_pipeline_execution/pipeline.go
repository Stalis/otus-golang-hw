package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func stageWrapper(done In, in In, stage Stage) Out {
	res := make(Bi)
	go func() {
		defer close(res)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				res <- v
			}
		}
	}()
	return stage(res)
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, st := range stages {
		if st != nil {
			out = stageWrapper(done, out, st)
		}
	}

	return out
}
