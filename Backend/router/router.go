package router

import (
        "go-backend/controller"
        "github.com/gorilla/mux"
)

func Router() *mux.Router {
        router := mux.NewRouter()

        router.HandleFunc("/api/getPost", controller.GetPost).Methods("GET", "OPTIONS")
        router.HandleFunc("/api/create/{id}", controller.Create).Methods("GET", "OPTIONS")

        return router
}
