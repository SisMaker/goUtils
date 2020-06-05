package eventWay

type Event struct {
	Name  string
	CfgId int
}

type HandlerErr struct {
	Event   Event
	Handler Handler
	ErrInfo error
}

type Handler func(msg interface{}) error

var eventMgr = map[Event][]Handler{}

func Register(event Event, handler Handler) {
	if handlers, ok := eventMgr[event]; ok {
		eventMgr[event] = append(handlers, handler)
	} else {
		eventMgr[event] = []Handler{handler}
	}
}

func Publish(event Event, msg interface{}) (errs []HandlerErr) {
	if handlers, ok := eventMgr[event]; ok {
		for _, handler := range handlers {
			if err := handler(msg); err != nil {
				errs = append(errs, HandlerErr{Event: event, Handler: handler, ErrInfo: err})
			}
		}
	}
	return errs
}
