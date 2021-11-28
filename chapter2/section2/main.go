package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Comment struct {
	CommentID int
	ArticleID int
	Message   string
	CreatedAt time.Time
}

type Article struct {
	ID          int
	Title       string
	Contents    string
	UserName    string
	NiceNum     int
	CommentList []Comment
	CreatedAt   time.Time
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}
	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "second comment",
		CreatedAt: time.Now(),
	}
	article := Article{
		ID:          1,
		Title:       "first article",
		Contents:    "This is the test article.",
		UserName:    "saki",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s\n", jsonData)
}
