package main

import (
        "fmt"
        "log"
        "net/http"
        "go-backend/router"
)

const port string = "8080"

func main() {
        r := router.Router()

        fmt.Printf("Server is running at localhost:%v\n", port)
        log.Fatal(http.ListenAndServe(":" + port, r))
}
