package main

import (
	"fmt"
	"time"
)

func main() {

	stn := Station{
		Coordinates: &Point{
			X:                2000,
			Y:                4000,
			Z:                288,
			CoordinateSystem: "Ghana National Grid",
		},
		Distance: 1000,
		unit:     "m",
		next:     nil,
		prev:     nil,
	}

	stn2 := Station{
		Coordinates: &Point{
			X:                2500,
			Y:                7000,
			Z:                300,
			CoordinateSystem: "Ghana National Grid",
		},
		Distance: 1500,
		unit:     "m",
		next:     nil,
		prev:     nil,
	}

	t := Traverse{}

	done := make(chan interface{}, 100)
	for i := 0; i < 100; i++ {
		go func() {
			t.PushStationBack(stn)
			t.PushStationBack(stn2)
			done <- "Goroutine is done!"

		}()

	}
	message := <-done
	fmt.Println(message)
	start := time.Now()
	t.traverse()
	elapsed := time.Since(start)
	fmt.Printf("%.10f has elapsed\n", elapsed.Seconds())
	fmt.Println(t.size)

}
