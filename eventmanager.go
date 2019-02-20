package winc

type EventHandler func(arg *Event)

type EventManager struct {
	handler EventHandler
}

func (evm *EventManager) Fire(arg *Event) {
	if evm.handler != nil {
		evm.handler(arg)
	}
}

func (evm *EventManager) Bind(handler EventHandler) {
	evm.handler = handler
}
