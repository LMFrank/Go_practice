package main

import (
	"fmt"
	"time"
)
import "course/mooc/retriever/mock"
import "course/mooc/retriever/real"

const url = "http://www.baidu.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{"name": "Tom", "course": "golang"})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake baidu.com",
	})
	return s.Get(url)
}

func main() {
	var r Retriever

	retriever := mock.Retriever{Contents: "this is a fake baidu.com"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	inspect(r)

	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.Timeout)

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
