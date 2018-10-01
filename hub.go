package main

import (
	"go.uber.org/zap"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	rooms map[int64]*Room

	// Registered clients.
	clients map[*Client]bool

	// 서버 전체 브로드 캐스트
	broadcastAll chan []byte

	eventQ chan *Event

	logger *zap.SugaredLogger

	userSeq int64
}

func newHub() *Hub {

	rawlogger, _ := zap.NewDevelopment() // TODO build tag 써서 분리할 수 있도록 (상용은 NewProduction())
	logger := rawlogger.Sugar()

	return &Hub{
		broadcastAll: make(chan []byte),
		eventQ:       make(chan *Event),
		clients:      make(map[*Client]bool),
		logger:       logger,
		userSeq:      1000,
	}
}

func (h *Hub) run() {
	for {
		// select case 처리되기 전엔 block 되며 한번에 한 case 만 실행됨
		select {
		//case client := <-h.register:
		//	h.clients[client] = true
		//	client.id = h.userSeq
		//	h.userSeq++
		//case client := <-h.unregister:
		//	if _, ok := h.clients[client]; ok {
		//		delete(h.clients, client)
		//		close(client.send)
		//	}
		//case message := <-h.broadcastAll:
		//	for client := range h.clients {
		//		select {
		//		case client.send <- message:
		//		default:
		//			close(client.send)
		//			delete(h.clients, client)
		//		}
		//	}

		// todo 이벤트 핸들링
		case e := <-h.eventQ:

			switch e.mainType {
			case "packet":
				h.logger.Debug("packet processing sub type :", e.subType)
				handler := msgHandler[e.subType]
				handler(e)
				break
			case "internal":
				break
			}

		}

	}
}
