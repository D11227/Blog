package main

import (
        "fmt"
        "log"
        "net/http"
        "go-backend/routers"
)

const port string = "8080"

func main() {
        r := routers.Router()

        fmt.Printf("Server is running at http://localhost:%v\n", port)
        log.Fatal(http.ListenAndServe(":" + port, r))
}
