package models

type USER struct {
        Id              string  `json: "id"`
        Username        string  `json: "username"`
        Password        string  `json: "password"`
}
