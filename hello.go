
package main
import (
"fmt"
)

const PrefixName = "Hello, "
const spanishPrefix = "Halo, "
const hindiPrefix = "Hallo, "
const Hindi = "Hindi"
const Spanish = "Spanish"
func main(){
	fmt.Println(Hello("",""))
}

func Hello(name string,language string)string  {
	if name == ""{
		name = "World"
	}



	return greetingPrefix(language) +name
}
func greetingPrefix(language string)(prefix string)  {
	switch language {
	case Hindi:
		prefix = hindiPrefix
	case Spanish:
		prefix = spanishPrefix
	default:
		prefix = PrefixName
	}
	return
}