package routing

import (
	"fmt"
	"net"
	"net/http"

	"github.com/bitdecaygames/fireport/server/services"

	"github.com/gorilla/mux"
)

// ServeGame will load all routes and begin serving the game on addr.
// This call will block until the server has stopped
func ServeGame(port int, svcs *services.MasterList) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}

	return serveInternal(listener, svcs)
}

func serveInternal(l net.Listener, svcs *services.MasterList) error {
	r := mux.NewRouter()
	RegisterAll(r, svcs)

	fmt.Printf("Starting server on %v\n", l.Addr())
	return http.Serve(l, r)
}
