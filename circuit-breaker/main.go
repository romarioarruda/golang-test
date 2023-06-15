package main

import(
	"fmt"
	"time"

	"github.com/sony/gobreaker"
	"golang-test/circuit-breaker/server"
	"golang-test/circuit-breaker/client"
)

func main() {
	go server()

	// call with circuit breaker
	cb := gobreaker.NewCircuitBreaker(
		gobreaker.Settings{
			Name:        "my-circuit-breaker",
			MaxRequests: 3,
			Timeout:     3 * time.Second,
			Interval:    1 * time.Second,
			ReadyToTrip: func(counts gobreaker.Counts) bool {
				return counts.ConsecutiveFailures > 3
			},
			OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
				fmt.Printf("CircuitBreaker '%s' changed from '%s' to '%s'\n", name, from, to)
			},
		},
	)

	fmt.Println("Call with circuit breaker")
	for i := 0; i < 100; i++ {
		_, err := cb.Execute(func() (interface{}, error) {
			err := DoReq()
			return nil, err
		})
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)

	}
}