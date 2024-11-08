package main // start from package main

import (
	"fmt"

	"github.com/NhongSun/Learn-Go/nuzzy" // import sunshine package
	"github.com/google/uuid"             // import uuid package
)

// start from function main
func main() {
	fmt.Println("Hi Hi Hi")

	id := uuid.New()
	fmt.Printf("UUID: %s", id)

	nuzzy.SayHiNuzzy()
	nuzzy.SayHiBuzzy()
}
