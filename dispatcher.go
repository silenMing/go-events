package main

import(
	"fmt"
)

type dispatcher struct{
	listeners map[string] *EventChain
}

type EventChain struct{
	chs         [] chan *Event
	callbacks   [] *EventCallback
}

type Event struct{
	eventName string
	Params    map[string] interface{}
}

func Create() {
	
}