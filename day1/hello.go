package main

func sayHi(name string) string {
	return "Hello " + name
}

func main() {
	message := sayHi("somkiat")
	println(message)
}
