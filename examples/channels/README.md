## Channels

Official Docs: [Effective Go: Concurrency](https://go.dev/doc/effective_go#concurrency)

### Goals

To process requests, events, data, etc, in real-time.

Easy troubleshooting and replaying of events.

### Strategy

Use channels between _go routines_ to complete tasks in real-time.
When data is ready on a channel, it is processed.

Avoid queues and polling, because it the timings are unpredictable and testing is more complex.

### Design Notes

#### Debug

My goal is to be able to debug like this, enabling debug via a protected header in this case:

```
$ curl -d '{"food": "mushroom pizza","payment_method": "cash"}'
  -H "Content-Type: application/json"
  --header "X-MyApp: someAuthData,debug"
  -X POST http://localhost:8080/myapp/order/
```

To support this granularity I add a debug flag to the type being passed through channels.
I also make consistently formatted log entries easy to use in code.

#### Channel Messages

I create a base [ChannelMessage](https://github.com/thisdougb/practicalgo/blob/develop/examples/channels/message.go) type, and hang some logging functions from it.
When I [embed](https://github.com/thisdougb/practicalgo/blob/develop/examples/channels/order.go) this in the type I want to send on a channel, it brings simple but effective logging.

```
type ChannelMessage struct {
  TimeReceived time.Time     `json:"time_received"` // used to log processing duration
  Debug        bool          `json:"debug"`         // to enable per-message debug
  Synthetic    bool          `json:"synthetic"`     // to enable per-message synthetic testing
  idPrefix     func() string // use this to add an id into the log message
}
func (cm *ChannelMessage) LogInfo(msg)
func (cm *ChannelMessage) LogDebug(msg)
func (cm *ChannelMessage) LogError(msg)

type Order struct {
  ChannelMessage
  ChefId  int64  `json:"chef_id"`
  Payment string `json:"payment_method"`
  Food    string `json:"food"`
}

func NewOrder(food string) *Order {
  o := Order{}
  o.TimeReceived = time.Now()
  o.idPrefix = func() string { return fmt.Sprintf("%d.%s", o.ChefId, o.Food) }

  return &o
}
```

A requirement of this approach is a New function, in order to setup the TimeReceived and idPrefix values.
A small price to pay for easy logging anywhere along a pipeline.

This:

```
order.LogInfo("new customer order")
order.LogDebug("the customer wanted pineapple")
```

gives this:

```
2023/10/15 13:25:18 INFO +0.0s [0.pineapple pizza] new customer order
2023/10/15 13:25:18 DEBUG +0.0s [20.pineapple pizza] the customer wanted pineapple
```

#### Channels

I prefix channel names with _ch_, _inch_, _ouch_.
Signifying a bidirectional, input, or output channel repsectively.

I defer the closing of the channel as close as possible to the making of the channel.
This is simply more readable.

```
chOrders := make(chan *Order) // bidirectional channel name
defer close(chOrders)
```

When accepting or returning channels, functions should enforce a direction and type.
The return channel type is often an in-channel for the caller.
And this function (the channel owner) writes to it, and closes it.

```
kitchen := func(done chan interface{}, inchOrders <-chan *Order) <-chan *Order {}
```

I favour the readability of the for-select pattern when consuming messages from channels.

```
for {
  select {
  case order := <-inchOrders:
    // do work
  case <-done:
    return
  }
}
```
