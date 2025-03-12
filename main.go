package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("Learn GO")
	todos.add("Learn GO more")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v", todos)
}
