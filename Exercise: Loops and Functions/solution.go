package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	var closetZ float64
	acceptedChange := 0.00000000001
	for i := 1; i <= 100 && z < x; i++ {
		var delta = (z*z - x) / (2 * z)

		//condition of stopping to find, meaing get answer in last z
		if math.Abs(delta) < acceptedChange {
			fmt.Printf("Run %d times to find answer %f. Not bad.\n", i, z)
			return closetZ
		}
		z -= delta
		fmt.Printf("delta %d is delta %e\n", i, delta)
		closetZ = z

		if i == 100 {
			fmt.Println("\nRun 100 times. No good!")
			fmt.Println("Accepted change should be larger than %e", acceptedChange)
			fmt.Println("z is ", z)
		}
	}
	return closetZ
}

func main() {
	ans := Sqrt(2000000000000000)
	fmt.Printf("ans^2 is %f", ans*ans)
}
