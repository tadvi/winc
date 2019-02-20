package winc

type Event struct {
	Sender Controller
	Data   interface{}
}

func NewEvent(sender Controller, data interface{}) *Event {
	return &Event{Sender: sender, Data: data}
}
