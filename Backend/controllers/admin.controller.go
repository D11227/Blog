package controllers

import (
        "net/http"
        "io/ioutil"
        "encoding/json"
        "go-backend/models"
)

func Search(users []models.USER, condition func(name string) bool) int {
	for index, user := range users {
		if condition(user.Username) {
			return index
		}
	}
	return -1
}

func Login(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "*")

        user := models.USER{}

        json.NewDecoder(r.Body).Decode(&user)
        data := GetUsersList()

        index := Search(data, func(name string) bool {
                return name == user.Username
        })

        if index == -1 || data[index].Password != user.Password {
                json.NewEncoder(w).Encode([]string { "Username or password is wrong!" })
                return
        }

        json.NewEncoder(w).Encode([]string { "Login successfully", `
        <div class="input-group mb-3 mt-3">
                <input class="form-control" type="file" id="img">
                </div>
                <hr />
                <div class="align-items-center mt-2">
                    <div class="input-group mb-3">
                        <div class="input-group-prepend">
                            <span class="input-group-text" id="basic-addon1">Title</span>
                        </div>
                        <input
                            type="text"
                            class="form-control"
                            placeholder="Title"
                            aria-label="Title"
                            aria-describedby="basic-addon1"
                            name="title"
                            id="title"
                        />
                    </div>
                    <div class="input-group mb-3">
                        <div class="input-group-prepend">
                            <span class="input-group-text" id="basic-addon1">Writer</span>
                        </div>
                        <input
                            type="text"
                            class="form-control"
                            placeholder="writer"
                            aria-label="writer"
                            aria-describedby="basic-addon1"
                            name="writer"
                            id="writer"
                        />
                    </div>
                    <div class="input-group mb-3">
                        <div class="input-group-prepend">
                            <span class="input-group-text" id="basic-addon1">Tags</span>
                        </div>
                        <input
                            type="text"
                            class="form-control"
                            placeholder="tags"
                            aria-label="tags"
                            aria-describedby="basic-addon1"
                            name="tags"
                            id="tags"
                        />
                    </div>
                    <textarea id="content"></textarea>
                    <hr />
                    <button class="btn btn-primary cus-btn-primary mb-3 spinner-btn" onclick="app.uploadPost()">
                        Upload Book
                    </button>
                </div>` })
}

func GetUsersList() []models.USER {
        content, err := ioutil.ReadFile("./databases/users.database.json")
        if err != nil {
                panic(err)
        }

        data := make([]models.USER, 0)
        err = json.Unmarshal(content, &data)
        if err != nil {
                panic(err)
        }

        return data
}
