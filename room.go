package main

type Room struct {
	id    int64
	users map[int64]*Client
}
