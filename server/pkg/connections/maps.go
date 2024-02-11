package connections

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/thoratvinod/vi-chat/server/pkg/common/message"
)

// MessageChannelMap
type MessageChannelMap struct {
	syncMap sync.Map
}

func (mp *MessageChannelMap) Put(userID uint, userMsgChan chan *message.Message) {
	mp.syncMap.Store(userID, userMsgChan)
}

func (mp *MessageChannelMap) Get(userID uint) (chan *message.Message, error) {
	val, ok := mp.syncMap.Load(userID)
	if !ok {
		return nil, fmt.Errorf("can't find user's message channel in MessageChannelMap, userID=%v", userID)
	}
	msgUserChan, ok := val.(chan *message.Message)
	if !ok {
		return nil, fmt.Errorf("type assertion failed in MessageChannelMap, userID=%v", userID)
	}
	return msgUserChan, nil
}

func (mp *MessageChannelMap) Delete(userID uint) {
	mp.syncMap.Delete(userID)
}



type ConnectionMap struct {
	syncMap sync.Map
}

func (mp *ConnectionMap) Put(userID uint, conn *websocket.Conn) {
	mp.syncMap.Store(userID, conn)
}

func (mp *ConnectionMap) Get(userID uint) (*websocket.Conn, error) {
	val, ok := mp.syncMap.Load(userID)
	if !ok {
		return nil, fmt.Errorf("can't find user's connection in ConnectionMap, userID=%v", userID)
	}
	conn, ok := val.(*websocket.Conn)
	if !ok {
		return nil, fmt.Errorf("type assertion failed in ConnectionMap, userID=%v", userID)
	}
	return conn, nil
}

func (mp *ConnectionMap) Delete(userID uint) {
	mp.syncMap.Delete(userID)
}


func (mp *ConnectionMap) Range(f func(key, value any) bool) {
	mp.syncMap.Range(f)
}