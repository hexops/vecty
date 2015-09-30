package dispatcher

type ID int

var idCounter ID
var callbacks = make(map[ID]func(action interface{}))

func Dispatch(action interface{}) {
	for _, c := range callbacks {
		c(action)
	}
}

func Register(callback func(action interface{})) ID {
	idCounter++
	id := idCounter
	callbacks[id] = callback
	return id
}

func Unregister(id ID) {
	delete(callbacks, id)
}
