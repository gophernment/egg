package egg

type Worker func() interface{}

type Responder func() <-chan interface{}

func New(w Worker) Responder {
	resp := make(chan interface{})
	go func(ch chan<- interface{}) {
		defer func() {
			if r := recover(); r != nil {
				resp <- r
			}
		}()

		ch <- w()
	}(resp)
	return func() <-chan interface{} {
		return resp
	}
}
