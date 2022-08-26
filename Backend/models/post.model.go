package models

type POST struct {
        Id              string          `json: "id"`
        Title           string          `json: "title"`
        Writer          string          `json: "writer"`
        Image           string          `json: "image"`
        Content         string          `json: "content"`
        Date            string          `json: "date"`
        Tags            []string        `json: "tags"`
}
