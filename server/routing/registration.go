package routing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bitdecaygames/fireport/server/services"
	"github.com/gorilla/mux"
)

const (
	// APIv1 is the version 1 api prefix
	APIv1 = "/api/v1"
)

// RegisterAll will register all needed routes on the given router
func RegisterAll(r *mux.Router, svcs *services.MasterList) {
	lobby := &LobbyRoutes{
		Services: svcs,
	}
	lobby.AddRoutes(r)

	game := &GameRoutes{
		Services: svcs,
	}
	game.AddRoutes(r)

	pubsub := &Subscriber{
		Services: svcs,
	}
	pubsub.AddRoutes(r)

	r.HandleFunc("/", showAllRoutesHandler(r))
}

// TODO: Make a route that dumps out all routes
func showAllRoutesHandler(router *mux.Router) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		info := make([]string, 0)
		router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			pathTemplate, err := route.GetPathTemplate()
			if err == nil {
				info = append(info, fmt.Sprint("ROUTE:", pathTemplate))
			}
			methods, err := route.GetMethods()
			if err == nil {
				info = append(info, fmt.Sprintf("Methods: %v\n", strings.Join(methods, ",")))
			}
			return nil
		})
		w.Write([]byte(strings.Join(info, "\n") + "\n"))
	}
}
