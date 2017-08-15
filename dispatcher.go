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

func CreateEvent(eventName string,Params map[string] interface{}) *Event {
	return &Event{eventName: eventName,Params:Params}
}

func CreateEventChain() *EventChain {
    return &EventChain{chs: []chan *Event{},callbacks:[] *EventCallback{}}
}


type EventCallback func (*Event)

func SharedDispatcher() *Dispatcher {
    var _instance *Dispatcher
    if _instance == nil{
        _instance = &Dispatcher{}
        _instance.Init()
    }
    return _instance;
}

func (this *Dispatcher) Init() {
    this.listeners = make(map[string] *EventChain)
}

func (this *Dispatcher) AddEventListener (eventName string,callback *EventCallback) {
    EventChain,ok := this.listeners[eventName]
    if !ok {
        EventChain = CreateEventChain()
        this.listeners[eventName] = eventChain
    }
}