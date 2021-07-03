package main

import (
	"fmt"
	"math"
)

// type Car struct {
// 	NumberOfTires int
// 	Luxury        bool
// 	BucketSeats   bool
// 	Make          string
// 	Model         string
// 	Year          int
// }

// func main() {
// 	myCar := Car{
// 		NumberOfTires: 4,
// 		Luxury:        false,
// 		BucketSeats:   true,
// 		Make:          "Nissan",
// 		Model:         "Rogue",
// 		Year:          2014,
// 	}
// 	fmt.Println(myCar)
// }

// func main() {
// 	var animals []string
// 	animals = append(animals, "fish")
// 	animals = append(animals, "pony")
// 	animals = append(animals, "hippo")
// 	fmt.Println(animals)
// 	fmt.Println(animals[len(animals)-1])
// }

type Pair struct {
	base  int
	power int
}

func (p *Pair) exponent() int {
	return int(math.Pow(float64(p.base), float64(p.power)))
}

func main() {
	myPair := Pair{
		base:  2,
		power: 3,
	}
	fmt.Println(exponent(2, 3))
	fmt.Println(myPair.exponent())
	fmt.Println((&Pair{base: 3, power: 2}).exponent())
}

func exponent(base, power int) int {
	return int(math.Pow(float64(base), float64(power)))
}
