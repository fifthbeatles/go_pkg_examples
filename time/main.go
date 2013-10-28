package main

import (
	"fmt"
	"time"
)

func testTimeout(s1, s2 time.Duration) {
	c := make(chan int)
	go func() {
		time.Sleep(s1 * time.Second)
		c <- 10
	}()

	select {
	case i := <-c:
		fmt.Println("Get:", i)
	case t := <-time.After(s2 * time.Second):
		fmt.Println("Timeout at", t)
	}
}

func testTimer() {
	timer := time.AfterFunc(3*time.Second, func() {
		fmt.Println("After 3 seconds")
	})
	fmt.Println(timer.Stop()) // true

	fmt.Println(time.Now()) // start
	timer = time.NewTimer(time.Second)
	timer.Reset(3 * time.Second)
	select {
	case t := <-timer.C:
		fmt.Println(t) // start + 3 seconds
	}
	fmt.Println(timer.Stop()) // false
}

func testTicker() {
	c := time.Tick(time.Second)
	n := 4
	for now := range c {
		if n--; n < 0 {
			break
		}
		fmt.Println(now)
	}

	ticker := time.NewTicker(time.Second)
	n = 4
	for now := range ticker.C {
		if n--; n < 0 {
			break
			ticker.Stop()
		}
		fmt.Println(now)
	}
}

func testTime() {
	// Duration
	d1, err := time.ParseDuration("1h25m")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d1, d1.Hours(), d1.Minutes(), d1.Seconds())

	// Month && Weekday
	for i := 1; i <= 13; i++ {
		fmt.Printf("%s ", time.Month(i))
	}
	fmt.Println()
	for i := 0; i <= 7; i++ {
		fmt.Printf("%s ", time.Weekday(i))
	}
	fmt.Println()

	// Time
	t1 := time.Now()
	t2 := time.Date(2013, time.July, 1, 12, 0, 0, 0, time.UTC)
	t3, err := time.Parse("2006-01-02 15:04:05", "2013-08-01 15:00:00")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t1, t2, t3, t3.Local())
}

func main() {
	//testTimeout(1, 2)
	//testTimeout(2, 1)
	//testTimer()
	//testTicker()
	testTime()
}
