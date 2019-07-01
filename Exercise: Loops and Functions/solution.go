package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)     // the inital number to try
	var closetZ float64 // the answer to return
	maxTry := 100       // try only 100 times
	acceptedChange := 0.0000000001
	for i := 1; i <= maxTry; i++ {
		var delta = (z*z - x) / (2 * z)
		comparson := delta
		//convert minus number to postiive number
		if comparson < 0 {
			comparson = -comparson
		}

		//condition of stopping to find, meaning we got answer in last z
		if comparson < acceptedChange {
			fmt.Printf("Run %d times to find answer %f. Not bad.\n", i, z)
			return closetZ
		}
		z -= delta
		fmt.Printf("delta %d is delta %e\n", i, delta)
		closetZ = z

		if i == maxTry {
			fmt.Printf("\nRun %d times. No good!\n", maxTry)
			fmt.Println("Accepted change should be larger than %e", acceptedChange)
			fmt.Println("z is ", z)
		}
	}
	return closetZ
}

func main() {
	ans := Sqrt(2000000000)
	fmt.Printf("ans^2 is %f", ans*ans)
}
