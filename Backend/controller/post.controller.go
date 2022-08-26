package controller

import (
        "net/http"
        "io/ioutil"
        "encoding/json"
        "go-backend/models"
        "github.com/gorilla/mux"
)

func GetPost(w http.ResponseWriter, r *http.Request)  {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "*")

        data := GetPostsList()
        json.NewEncoder(w).Encode(data)
}

func Create(w http.ResponseWriter, r *http.Request)  {
        w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
        w.Header().Set("Access-Control-Allow-Origin", "*")

        params := mux.Vars(r)
        json.NewEncoder(w).Encode(params["id"])
}

func GetPostsList() []models.POST {
        content, err := ioutil.ReadFile("./databases/posts.database.json")
        if err != nil {
                panic(err)
        }

        data := make([]models.POST, 0)
        err = json.Unmarshal(content, &data)
        if err != nil {
                panic(err)
        }

        return data
}
