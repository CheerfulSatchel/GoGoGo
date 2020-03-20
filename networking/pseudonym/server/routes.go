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
		Route{"all", "GET", "/all", handlers.AllPseudonyms},
		Route{"add", "POST", "/add", handlers.AddPseudonym},
		Route{"get", "GET", "/pseudonym/:id", handlers.GetPseudonym},
	}

	return routes
}
