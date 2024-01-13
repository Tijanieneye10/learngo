// You can write Go code here!
package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	var expectedRateOfReturn = 5.5
	var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1 + expectedRateOfReturn / 100, float64(years))
	fmt.Println(futureValue)
}