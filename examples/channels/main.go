package main

import (
	"strings"
	"time"
)

/*

Example showing simple logging.

$ go run .
2023/10/15 13:12:06 INFO +0.0s [0.mushroom risotto] new customer order
2023/10/15 13:12:06 INFO +0.0s [20.mushroom risotto] ready
2023/10/15 13:12:06 INFO +0.0s [0.pineapple pizza] new customer order
2023/10/15 13:12:06 DEBUG +0.0s [20.pineapple pizza] the customer wanted pineapple
2023/10/15 13:12:07 ERROR +1.0s [20.pineapple pizza] the chef rages about pineapple on pizza
2023/10/15 13:12:07 INFO +1.0s [20.pineapple pizza] ready

*/

func main() {

	// universal done channel
	done := make(chan interface{})
	defer close(done)

	// message channels are created bidirectional and prefixed with ch
	chOrders := make(chan *Order)
	defer close(chOrders)

	// Use inch and ouch prefixes to add meaning to channel parameters.
	// I also give channel parameters a typed direction.
	kitchen := func(done chan interface{}, inchOrders <-chan *Order) <-chan *Order {

		chReady := make(chan *Order)

		go func() {

			defer close(chReady) // the creator closes the channel

			// use the for-select loop pattern, with a done channel
			for {
				select {

				case order := <-inchOrders:
					order.LogInfo("new customer order")
					order.ChefId = 20

					if strings.Contains(order.Food, "pizza") {
						order.LogDebug("the customer wanted pineapple")

						if strings.Contains(order.Food, "pineapple") {
							time.Sleep(1 * time.Second)
							order.LogError("the chef rages about pineapple on pizza")
						}
					}
					// work complete, send the order on its way
					chReady <- order

				case <-done:
					return
				}
			}
		}()
		return chReady
	}

	// some test orders, the second enables debug
	orders := []*Order{NewOrder("mushroom risotto"), NewOrder("pineapple pizza")}
	orders[1].Debug = true

	// launch the kitchen go routine, and get back its output channel
	inchReady := kitchen(done, chOrders)

	for _, order := range orders {
		chOrders <- order
		ready := <-inchReady
		ready.LogInfo("ready")
	}
}
