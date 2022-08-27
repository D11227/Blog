package controllers

import (
        "net/http"
        "io/ioutil"
        "encoding/json"
        "go-backend/models"
)

func GetPost(w http.ResponseWriter, r *http.Request)  {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "*")

        data := GetPostsList()
        json.NewEncoder(w).Encode(data)
}

func CreatePost(w http.ResponseWriter, r *http.Request)  {
        w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
                return
        }

        post := models.POST{}

        json.NewDecoder(r.Body).Decode(&post)

        AddToDatabase(&post)

        json.NewEncoder(w).Encode([]string { "Created post successfully!" })
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

func AddToDatabase(post *models.POST)  {
        data := GetPostsList()
        data = append(data, *post)

        file, err := json.MarshalIndent(data, "", " ")

        if err != nil {
                panic(err)
        }

        err = ioutil.WriteFile("./databases/posts.database.json", file, 0644)

        if err != nil {
                panic(err)
        }
}
