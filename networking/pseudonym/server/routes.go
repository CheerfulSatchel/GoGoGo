package server

import (
	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/handlers"
	"github.com/julienschmidt/httprouter"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes {
	routes := Routes{
		Route{"index", "GET", "/", handlers.AllPseudonyms},
		Route{"create", "POST", "/pseudonym", handlers.CreatePseudonyms},
		Route{"read", "GET", "/pseudonym/:id", handlers.GetPseudonym},
	}

	return routes
}
