package routers

import (
        "go-backend/controllers"
        "github.com/gorilla/mux"
)

func Router() *mux.Router {
        router := mux.NewRouter()

        /* Post Controller */
        router.HandleFunc("/api/getPost", controllers.GetPost).Methods("GET", "OPTIONS")
        router.HandleFunc("/api/createPost", controllers.CreatePost).Methods("POST", "OPTIONS")

        /* Admin Controller */
        router.HandleFunc("/api/login", controllers.Login).Methods("POST", "OPTIONS")

        return router
}
