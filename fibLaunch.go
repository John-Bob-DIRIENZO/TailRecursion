package tailRecursion

import (
	"fmt"
	"sync"
	"time"
)

func FibLaunch(n int) {
	start := time.Now().UnixMicro()

	wg := new(sync.WaitGroup)
	wg.Add(2)

	fmt.Println("========== START ==========")

	go func() {
		fmt.Println("Classic recursion ===== ", FibRec(n))
		fmt.Println("Needed Time : ", time.Now().UnixMicro()-start, " microseconds")
		fmt.Println("===== /Classic =====")
		wg.Done()
	}()

	go func() {
		fmt.Println("Tail recursion ===== ", FibTailRec(n))
		fmt.Println("Needed Time : ", time.Now().UnixMicro()-start, " microseconds")
		fmt.Println("===== /Tail =====")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("========== END ==========")
}

func FibStraight() {
	fib := []int{0, 1}

	//wait := time.Now().Add(100 * time.Millisecond).Unix()

	i := 2
	for i < 50 {
		fib = append(fib, fib[i-2]+fib[i-1])
		i++
	}

	fmt.Println("===== No recursion =====")
	fmt.Println(fib)
}
