package main

import "websocket/lib"

type MsgHandlerFunc func(*Event) error

var msgHandler = map[string]MsgHandlerFunc{
	"RequestLogin":    RequestLogin,
	"RequestRoomJoin": RequestRoomJoin,
	"RequestRoomExit": RequestRoomExit,
}

func RequestLogin(event *Event) error {
	client := event.client
	var req *MsgRequestLogin
	req = (event.data).(*MsgRequestLogin)
	client.hub.logger.Debug("request login, id:", req.Uid)

	var ret string
	ret = "login complete id " + lib.Itoa64(req.Uid)
	client.send <- []byte(ret)
	return nil
}

func RequestRoomJoin(event *Event) error {
	client := event.client
	client.hub.logger.Debug("request room join, id:", client.id)
	return nil
}

func RequestRoomExit(event *Event) error {
	client := event.client
	client.hub.logger.Debug("request room exit, id:", client.id)
	return nil
}
