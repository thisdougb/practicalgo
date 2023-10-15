package main

import (
	"fmt"
	"time"
)

type ChannelMessage struct {
	TimeReceived time.Time     `json:"time_received"` // used to log processing duration
	Debug        bool          `json:"debug"`         // to enable per-message debug
	Synthetic    bool          `json:"synthetic"`     // to enable per-message synthetic testing
	idPrefix     func() string // use this to add an id into the log message
}

/*

A Message type suitable for send through channels.
It simplifies logging.

Logs messages like this:

2023/10/15 11:11:08 INFO +0.0s the customer wants a Diavolo
2023/10/15 11:11:08 ERROR +0.0s [23.Diavolo] the customer wants pineapple
*/

// Public
func (cm *ChannelMessage) LogInfo(msg string) {
	cm.writeToLog("INFO", msg)
}

func (cm *ChannelMessage) LogError(msg string) {
	cm.writeToLog("ERROR", msg)
}

func (cm *ChannelMessage) LogDebug(msg string) {
	if cm.Debug {
		cm.writeToLog("DEBUG", msg)
	}
}

// Private
func (cm *ChannelMessage) writeToLog(severity string, msg string) {

	if cm.idPrefix != nil {
		fmt.Printf("%s %s +%s [%s] %s\n",
			time.Now().UTC().Format("2006/01/02 15:04:05"),
			severity,
			cm.sinceReceived(),
			cm.idPrefix(),
			msg)
	} else {
		fmt.Printf("%s %s +%s %s\n",
			time.Now().UTC().Format("2006/01/02 15:04:05"),
			severity,
			cm.sinceReceived(),
			msg)
	}
}

func (cm *ChannelMessage) sinceReceived() string {
	t := time.Since(cm.TimeReceived).Seconds()
	return fmt.Sprintf("%.1fs", t)
}
