package lazy

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestLazy(t *testing.T) {
	b1 := Egg(func() interface{} {
		fmt.Println("working b1")
		return "b1"
	})
	b2 := Egg(func() interface{} {
		<-time.After(1 * time.Second)
		fmt.Println("working b2")
		return "b2"
	})
	b3 := Egg(func() interface{} {
		fmt.Println("working b3")
		return "b3"
	})

	fmt.Println(<-b1())
	fmt.Println(<-b2())
	fmt.Println(<-b3())
}

func TestLazyPanic(t *testing.T) {
	b1 := Egg(func() interface{} {
		fmt.Println("working b1")
		return "b1"
	})
	b2 := Egg(func() interface{} {
		panic(errors.New("panic!!"))
		fmt.Println("working b2")
		return "b2"
	})
	b3 := Egg(func() interface{} {
		fmt.Println("working b3")
		return "b3"
	})

	fmt.Println(<-b1())
	fmt.Println(<-b2())
	fmt.Println(<-b3())
}
