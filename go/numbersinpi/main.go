package main

import (
	"fmt"

	"github.com/vagnerclementino/algorithm/go/numbersinpi/pi"
)

func main() {
	const PI = "3141592653589793238462643383279"
	numbers := []string{"314159265358979323846", "26433", "8", "3279", "314159265", "35897932384626433832", "79"}
	fmt.Println(pi.NumbersInPi(PI, numbers))
}
