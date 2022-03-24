package main

import "fmt"

type Human struct {
	Name    string
	Age     int
	Weight  int
	Fingers int
}

func (h Human) fingersNumber() int {
	return 5
}

//наследование от Human
type Action struct {
	Human
}

func (a Action) SayName() string {
	return "Hello, my Name is " + a.Name
}

func main() {
	action := Action{
		Human{
			Name:   "maui",
			Age:    20,
			Weight: 70,
		},
	}

	fmt.Println(action.SayName())
	fmt.Println(action.Human.fingersNumber())
	fmt.Println(action.fingersNumber())
}
