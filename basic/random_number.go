package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println(rand.Intn(101))
	fmt.Println(rand.Intn(101) + 1)

	rand := rand.New(rand.NewSource(time.Now().Unix()))
	fmt.Println(rand.Intn(101))

	fmt.Println(rand.Float64()) // will be always between 0.0 and 1.0

}
