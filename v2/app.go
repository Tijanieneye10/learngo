// You can write Go code here!
package main

import (
	"fmt"
	"math"
)

func mains() {
	var investmentAmount, years float64 = 1000, 10
	expectedRateOfReturn := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedRateOfReturn/100, years)
	fmt.Println(futureValue)

	fmt.Scan(&investmentAmount)
	fmt.Println(investmentAmount)
}
