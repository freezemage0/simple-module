// simple-module project simple-module.go
package simple.module


import "fmt"


func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}