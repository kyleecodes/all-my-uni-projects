package awesomeProject

import "fmt"

func main() {
	fmt.Println(sayHi("Marco"))
}

func sayHi(person string) string {
	return fmt.Sprintf("Hi %s", person)
}