package server

import (
	"github.com/CheerfulSatchel/GoGoGo/networking/pseudonym/database_service/handlers"
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
		Route{"create_tables", "PUT", "/tables", handlers.CreateTables},
		Route{"add_pseudonym", "PUT", "/pseudonym", handlers.AddPseudonym},
	}

	return routes
}
