package main

import (
	"fmt"
	"swapper/modules/twitter"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("checking username thisisatestlol_")
	// start := time.Now()
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)

		wg.Add(1)
		go StartChecking(i)
	}

	wg.Wait()
}

func StartChecking(i int) {
	guestToken := twitter.GetGuestToken()

	for {
		// time.Sleep(50 * time.Millisecond)
		avai, ratelimit := twitter.CheckUser("exgoju", guestToken)

		if ratelimit < 10 {
			guestToken = twitter.GetGuestToken()
			fmt.Println("refreshed token")
		}
		if avai {
			fmt.Println("username is available")
			break
		} else {
			fmt.Println("[", i, "]", "username not available, requests before token refresh:", ratelimit)
		}
		// duration := time.Since(start)
		// fmt.Println(duration)
	}
}
