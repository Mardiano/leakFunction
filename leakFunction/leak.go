package leak

import (
	"fmt"
	"sync"
	"time"
)

func LeakyFunction(wg sync.WaitGroup) {
	fmt.Println("LEAK FUNCTION")
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pandas")
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
