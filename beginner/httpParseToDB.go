package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type post struct {
	UserId int    `json:"user_id"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type comment struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Body  string `json:"body"`
}

func main() {
	host, user, password, dbname, port, sslmode := "localhost", "postgres", "mysql", "test", "5432", "disable"
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode))
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	request, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=7")
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err.Error())
	}

	var posts []post
	err = json.Unmarshal(res, &posts)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, post := range posts {
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err.Error())
		}
		var postId int

		row := tx.QueryRow(fmt.Sprintf("INSERT INTO posts(user_id,title,body) VALUES (%d,'%s','%s') RETURNING id", 7, post.Title, post.Body))
		if err := row.Scan(&postId); err != nil {
			fmt.Println(postId)
			tx.Rollback()
			log.Fatal(err.Error())
		}

		commentRequest, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=" + strconv.Itoa(post.Id))
		if err != nil {
			log.Fatal(err.Error())
		}

		res, err := ioutil.ReadAll(commentRequest.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		var comments []comment
		err = json.Unmarshal(res, &comments)
		if err != nil {
			log.Fatal(err.Error())
		}

		for _, comment := range comments {
			func() {
				_, err := tx.Exec(fmt.Sprintf("INSERT INTO comments(post_id, name, email, body) VALUES (%d,'%s','%s','%s')", postId, comment.Name, comment.Email, comment.Body))
				if err != nil {
					tx.Rollback()
					log.Fatal(err.Error())
				}
			}()
		}

		tx.Commit()
	}

	var a string
	fmt.Scan(a)
}
