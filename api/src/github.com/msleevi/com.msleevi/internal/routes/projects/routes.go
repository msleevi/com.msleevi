package projects

import (
	"github.com/msleevi/com.msleevi/internal/constants"
	"github.com/msleevi/com.msleevi/internal/routes/projects/handlers"

	"github.com/gorilla/mux"
)

/*
AddProjectsRoutes does
*/
func AddProjectsRoutes(r *mux.Router) {
	r.HandleFunc("/projects/new", handlers.NewProjectsHandler).Methods(constants.POST)
	r.HandleFunc("/projects", handlers.ProjectsHandler).Methods(constants.GET)
	r.HandleFunc("/projects/remove", handlers.RemoveProjectsHandler).Methods(constants.DELETE)
}
