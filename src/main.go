package main

const memorySize = 2000000

func main() {
	// Initialize a simulated memory
	var m Memory
	m.Init(memorySize)

	// Test load function
	m.Load("../data/array_test1.txt")
}