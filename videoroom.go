// Janus code specifically relating to the videoroom plugin

package videoroom

import (
	"fmt"
	"time"
	"github.com/csparrell/janus-go"
)

// janus.go provides KeepAlive() to send a heartbeat, but
// no routine to send them repeatedly.  This function
// sends them (currently at hardwired 30 second intervals)
// until signalled to stop on the arrest channel
//
func HeartBeat(s *janus.Session, arrest chan bool) {

	ticker := time.NewTicker (30 * time.Second)

	for {
		select {
		case <-arrest:
			fmt.Println("Heartbeat Stopped")
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("KeepAlive heartbeat sent at", t)
			_, erm := s.KeepAlive()
			if erm != nil {
				fmt.Println(erm)
			}
		}
	}
}

// Given a janus.EventMsg, will return the string tag representing
// the name of the video room event type found in that message
// (see EventKeys in the videoroom types.go file)
//
func GetEventType (em *janus.EventMsg) (string) {
	//var testevent VideoroomTestEvent
	emData := em.Plugindata.Data

	// Note "videoroom:destroyed" may require special handling

	for k, _ := range EventTypes {
		if _, ok := emData[k]; ok {
			return k
		}
	}
	fmt.Println("Event is not a known video room event message type.")
	return ""
}

// Creates a function that will create the appropriate type struct
// for the actual message type stored in the passed janus.EventMsg
// ...i.e., it disambiguates janus.EventMsg into a more specific
//    message type
func GetEventTypeFunc (em *janus.EventMsg) (func() interface{}) {

	emData := em.Plugindata.Data

	// Note "videoroom:destroyed" may require special handling

	// Hmph: "The iteration order over maps is not specified and
	// is not guaranteed to be the same from one iteration to the next."
	// Hence, we use a slice for range, to operate on our map

	for _, k := range EventKeys {
		v := EventTypes[k]
		if _, ok := emData[k]; ok {
			//fmt.Println("Event received is of type: ", k)
			return v
		}
	}
	
	fmt.Println("Event is not a known video room event message type.")
	// probably fatal, but lets at least try this
	return EventTypes["videoroom"]
}
