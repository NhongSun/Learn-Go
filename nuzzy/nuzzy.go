package nuzzy

import (
	"fmt"

	"github.com/NhongSun/Learing-Go/nuzzy/internal/secret"
)

func SayHiNuzzy() {
	fmt.Println("Hi Nuzzy")
	// sayHiBuzzy() can be called because it's in same package

	secret.SayHiSecret()
}
