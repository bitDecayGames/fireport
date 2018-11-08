package routing

import (
	"reflect"

	"github.com/bitdecaygames/fireport/server/pogo"

	"github.com/gorilla/websocket"
)

// PlayerConnWrapper is a simple wrapper that does some magical type setting before sending messages
type PlayerConnWrapper struct {
	con *websocket.Conn
}

// WriteJSON will set the type of the msg before sending it over the websocket
func (pc *PlayerConnWrapper) WriteJSON(msg pogo.Typer) error {
	trueType := reflect.TypeOf(msg)
	msg.SetType(trueType.String())
	return pc.con.WriteJSON(msg)
}

// Close closes the connection
func (pc *PlayerConnWrapper) Close() error {
	return pc.con.Close()
}
