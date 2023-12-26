package model

import "github.com/edgedb/edgedb-go"

type User struct {
	Id        edgedb.UUID `edgedb:"id"`
	Name      string      `edgedb:"name"`
	Email     string      `edgedb:"email"`
	Password  string      `edgedb:"password"`
	Following []User      `edgedb:"following"`
	Followers []User      `edgedb:"followers"`
}

type Post struct {
	Id      edgedb.UUID `edgedb:"id"`
	Title   string      `edgedb:"title"`
	Content string      `edgedb:"content"`
	Author  User        `edgedb:"author"`
}

type Notification struct {
	Id      edgedb.UUID `edgedb:"id"`
	Type    string      `edgedb:"type"`
	Message string      `edgedb:"message"`
	User    User        `edgedb:"user"`
}
