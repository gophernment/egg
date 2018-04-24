package lazy

import "fmt"

type Worker func() interface{}

type Responder func() <-chan interface{}

func Egg(w Worker) Responder {
	resp := make(chan interface{})
	go func(ch chan interface{}) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
				resp <- nil
			}
		}()

		ch <- w()
	}(resp)
	return func() <-chan interface{} {
		return resp
	}
}
