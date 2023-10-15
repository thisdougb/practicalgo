package main

import (
	"fmt"
	"time"
)

// An example of an entity that is sent through channels.
// I embed the ChannelMessage type, to get the logging functions.
type Order struct {
	ChannelMessage
	ChefId  int64  `json:"chef_id"`
	Payment string `json:"payment_method"`
	Food    string `json:"food"`
}

func NewOrder(food string) *Order {

	o := Order{}
	o.TimeReceived = time.Now()
	o.Food = food

	// I override this function to add custom id for this type.
	// This will be added to all logged messages, allowing easy grep/filtering of logs.
	//
	//     2023/10/15 13:06:26 INFO +1.0s [20.pineapple pizza] ready
	//
	o.idPrefix = func() string { return fmt.Sprintf("%d.%s", o.ChefId, o.Food) }

	return &o
}
