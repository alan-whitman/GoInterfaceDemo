// package main

// import "fmt"

// type MyInterface interface {
// 	DoTheThing()
// }

// type Strategy1 struct {
// 	myname string
// }

// func (s *Strategy1) DoTheThing() {
// 	fmt.Println(s.myname)
// }

// type Strategy2 struct {
// 	myname string
// }

// func (s1 *Strategy2) DoTheThing() {
// 	fmt.Println("Strategy2")
// }

// type Strategy3 struct{}

// func (s1 *Strategy3) DoTheThing() {
// 	fmt.Println("Strategy3")
// }

// func main() {
// 	strats := make(map[string]MyInterface)
// 	strats["One"] = &Strategy1{"Strategy 1, yo"}
// 	strats["Two"] = &Strategy2{}
// 	strats["Three"] = &Strategy3{}

// 	strats["One"].DoTheThing()
// 	strats["Two"].DoTheThing()
// 	strats["Three"].DoTheThing()

// }

package main

import "fmt"

func Strategy1(s string) {
	fmt.Println("hi from strat 1: ", s)
}

func Strategy2(s string) {
	fmt.Println("hi from strat 2: ", s)
}

func main() {
	strats := make(map[string]interface{})
	strats["One"] = Strategy1
	strats["Two"] = Strategy2

	strats["One"].(func(string))("hi")
	strats["Two"].(func(string))("Hello")

	// v.(func(string))("astring")
}
