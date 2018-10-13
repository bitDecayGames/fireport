package routing

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	apiv1 = "/api/v1"
)

// RegisterAll will register all needed routes on the given router
func RegisterAll(r *mux.Router) {
	lobby := &LobbyRoutes{}
	lobby.AddRoutes(r)

	game := &GameRoutes{}
	game.AddRoutes(r)

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
				info = append(info, fmt.Sprint("Methods:", strings.Join(methods, ",")))
			}
			return nil
		})
		w.Write([]byte(strings.Join(info, "\n") + "\n"))
	}
}
