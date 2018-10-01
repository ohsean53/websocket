package main

type Event struct {
	client   *Client
	mainType string
	subType  string
	data     interface{}
}

func NewEvent(c *Client, mt string, st string, d interface{}) *Event {
	return &Event{
		client:   c,
		mainType: mt,
		subType:  st,
		data:     d,
	}
}
