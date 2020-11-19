
package main
import (
"fmt"
)
func main(){
	fmt.Println(Hello("Abdul"))
}

func Hello(name string)string  {
	return "Hello, " +name
}