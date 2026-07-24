package dispatcher

import (
	"github.com/kuldeephsc/channel"
	"github.com/kuldeephsc/model"
)

type Dispatcher struct {
	channels []channel.Channel
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) Register(ch channel.Channel) {
	d.channels = append(d.channels, ch)
}

func (d *Dispatcher) Dispatch(alert model.Alert) map[string]string {
	results := make(map[string]string)
	for _, ch := range d.channels {
		err := ch.Send(alert)
		if err == nil {
			results[ch.Name()] = "Success"
			continue
		}
		//retry once
		err = ch.Send(alert)
		if err == nil {
			results[ch.Name()] = "Success after retry"
		} else {
			results[ch.Name()] = "Failed"
		}
	}
	return results
}
