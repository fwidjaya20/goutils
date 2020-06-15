package main

import (
	"fmt"
	"github.com/fwidjaya20/goutils"
)

func main() {
	fmt.Println("Test Implementation")

	foo := goutils.NewList(Person{})
	foo.Push(
		Person{
			name: "Catharine",
			age:  25,
		},
		Person{
			name: "Fredrick",
			age:  22,
		},
		Person{
			name: "Jeffrey",
			age:  19,
		},
		Person{
			name: "John",
			age:  35,
		},
		Person{
			name: "Dwayne",
			age:  13,
		},
		Person{
			name: "Walter",
			age:  25,
		},
	)

	fmt.Println("[TEST] Get : ", foo.Get().([]Person))

	finalFilter := foo.Filter(func(i int, el interface{}) interface{} {
		if el.(Person).age > 20 {
			return el
		}
		return nil
	})
	fmt.Println("[TEST] Filter : ", finalFilter.([]Person))

	finalFind := foo.Find(func(index int, el interface{}) interface{} {
		if el.(Person).age == 35 {
			return el
		}
		return nil
	})
	fmt.Println("[TEST] Find : ", finalFind.(Person))

	finalFindIndex := foo.FindIndex(func(index int, el interface{}) interface{} {
		if el.(Person).age == 35 {
			return el
		}
		return nil
	})
	fmt.Println("[TEST] FindIndex : ", finalFindIndex)

	finalInclude := foo.Some(func(index int, el interface{}) interface{} {
		if el.(Person).age == 35 {
			return el
		}
		return nil
	})
	fmt.Println("[TEST] Some : ", finalInclude)

	finalMap := foo.Map(&[]PersonNameOnly{}, func(index int, el interface{}) interface{} {
		if el.(Person).age > 15 {
			return PersonNameOnly{
				name: el.(Person).name,
			}
		}
		return nil
 	})
	fmt.Println("[TEST] Map : ", finalMap.([]PersonNameOnly))

	finalSlice := foo.Slice(1, foo.Count())
	fmt.Println("[TEST] Slice : ", finalSlice.([]Person))

	foo.Pop()
	fmt.Println("[TEST] Pop : ", foo.Get())

	foo.Shift()
	fmt.Println("[TEST] Shift : ", foo.Get())

	foo.Clear()
	fmt.Println("[TEST] Clear : ", foo.Get().([]Person))
}

type Person struct {
	name string
	age  int
}

type PersonNameOnly struct {
	name string
}