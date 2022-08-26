package router

import (
        "go-backend/controller"
        "github.com/gorilla/mux"
)

func Router() *mux.Router {
        router := mux.NewRouter()

        /* Post Controller */
        router.HandleFunc("/api/getPost", controller.GetPost).Methods("GET", "OPTIONS")
        router.HandleFunc("/api/createPost", controller.CreatePost).Methods("POST", "OPTIONS")

        /* Admin Controller */
        router.HandleFunc("/api/login", controller.Login).Methods("POST", "OPTIONS")

        return router
}
