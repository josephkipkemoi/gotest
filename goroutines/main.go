package goroutines

import (
	"math/rand"
	"fmt"
	"time"
)

func F(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n,":",i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

 