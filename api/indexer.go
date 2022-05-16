package api

import (
	"context"
	"log"
	"math"
	"strconv"
	"strings"
	"sync"

	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
)

type BookBrief struct {
	ID            int32    `json:"id"`
	Title         string   `json:"title"`
	Author        string   `json:"author"`
	Publisher     string   `json:"publisher"`
	CoverUrl      string   `json:"cover_url"`
	Categories    []string `json:"categories"`
	CommentCount  int64    `json:"comment_count"`
	BookmarkCount int64    `json:"bookmark_count"`
	RateCount     int64    `json:"rate_count"`
	RateSum       int64    `json:"rate_sum"`
	Pages         int64    `json:"pages"`
}

func toBookBrief(server *Server, book db.ListBooksRow) BookBrief {
	categories := []string{}
	for _, categoryID := range strings.Split(book.Categories, ",") {
		id, _ := strconv.Atoi(categoryID)
		if id == 0 {
			continue
		}
		categories = append(categories, server.categoryIndex[id])
	}

	return BookBrief{
		ID:            book.ID,
		Title:         book.Title,
		Author:        book.Author,
		Publisher:     book.Publisher,
		CoverUrl:      "/covers/" + book.CoverUrl,
		Categories:    categories,
		CommentCount:  book.CommentCount,
		BookmarkCount: book.BookmarkCount,
		RateCount:     book.RateCount,
		RateSum:       book.RateSum,
		Pages:         int64(book.Pages),
	}
}

func (server *Server) indexBooks() {
	server.bookIndex.DeleteAllDocuments()
	var wg sync.WaitGroup
	lastID := math.MaxInt32
	sem := make(chan int, 30)
	for {
		books, err := server.store.ListBooks(context.Background(), db.ListBooksParams{
			Limit:  1000,
			LastID: int32(lastID - 1),
		})
		if err != nil {
			log.Println("db error")
			break
		}
		if len(books) == 0 {
			break
		}

		bookBriefs := []BookBrief{}

		for i := range books {
			bookBriefs = append(bookBriefs, toBookBrief(server, books[i]))
		}
		wg.Add(1)
		sem <- 1
		go func(doc interface{}) {
			defer wg.Done()
			if _, err := server.bookIndex.AddDocuments(doc); err != nil {
				panic(err)
			}
			<-sem
		}(bookBriefs)

		lastID = int(books[len(books)-1].ID)
	}
	wg.Wait()
}

func (server *Server) updateBookIndex(id int32) {
	book, err := server.store.GetBookBrief(context.Background(), id)
	if err != nil {
		log.Fatal("couldn't read bookbrief from db")
		return
	}
	server.bookIndex.UpdateDocuments(toBookBrief(server, db.ListBooksRow(book)))
}
