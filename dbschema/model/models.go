package model

import "github.com/edgedb/edgedb-go"

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Following []User `json:"following"`
	Followers []User `json:"followers"`
}

type Post struct {
	Id      edgedb.UUID `edgedb:"id"`
	Title   string      `edgedb:"title"`
	Content string      `edgedb:"content"`
	Author  User        `edgedb:"author"`
}

type Notification struct {
	Type    NotificationType `json:"type"`
	Message string           `json:"message"`
}

type NotificationType string

const (
	FOLLOWER NotificationType = "FOLLOWER"
)
