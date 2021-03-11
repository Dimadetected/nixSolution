package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Post struct {
	UserId int    `json:"user_id"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	var i int = 1
	for i = 1; i <= 100; i++ {
		func() {
			var post Post
			req, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))
			if err != nil {
				log.Fatal(err.Error())
			}

			res, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err.Error())
			}

			err = json.Unmarshal(res, &post)
			if err != nil {
				log.Fatal(err.Error())
			}

			filename := "storage/posts/" + strconv.Itoa(i) + ".json"
			file, err := os.Create(filename)
			if err != nil {
				log.Fatal(err.Error())
			}
			file.Write([]byte(fmt.Sprintf("%v\n", post)))

		}()
	}
	var a string
	fmt.Scan(&a)
}
