package testModule

import "fmt"

func TestFunc(name string) string {
	message := fmt.Sprintf("Hello, %v, and welcome!", name)
	return message
}
