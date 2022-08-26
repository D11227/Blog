package models

type POST struct {
        Id              string          `json: "Id"`
        Title           string          `json: "Title"`
        Writer          string          `json: "Writer"`
        Image           string          `json: "Image"`
        Content         string          `json: "Content"`
        Date            string          `json: "Date"`
        Tags            []string        `json: "Tags"`
}
