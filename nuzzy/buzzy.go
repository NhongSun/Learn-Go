package nuzzy

import (
	"fmt"
)

func sayHiBuzzy() { // -> private function
	fmt.Println("Hi Buzzy")
}

func SayHiBuzzy() {
	sayHiBuzzy()
}
