package main // start from package main

import (
	"fmt"

	"github.com/NhongSun/Learing-Go/nuzzy" // import sunshine package
	"github.com/google/uuid"               // import uuid package
)

// start from function main
// test
func main() {
	fmt.Println("Hi Hi Hi")

	id := uuid.New()
	fmt.Println(id)

	nuzzy.SayHiNuzzy()
	nuzzy.SayHiBuzzy()
}
