package sampledata

import (
	"context"
	"fmt"
	"log"
	"sync"

	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
	"github.com/ndtai772/MyBookListBackend/util"
)

var gContext context.Context = context.Background()

const (
	NUM_OF_USERS           = 20
	NUM_OF_CATEGORIES      = 25
	NUM_OF_BOOKS           = 1000
	MAX_RATE_PER_USER      = NUM_OF_BOOKS / 4
	MAX_BOOKMARKS_PER_USER = NUM_OF_BOOKS / 4
	MAX_COMMENTS_PER_BOOK  = 10
)

func Gen(store *db.Store) {
	// addCategory(store)

	randomUsers(store)

	var wg sync.WaitGroup

	funcs := []func(store *db.Store){
		randomBookCategories,
		randomRates,
		randomBookmarks,
		randomComments,
	}

	wg.Add(len(funcs))

	for _, fn := range funcs {
		go func(f func(store *db.Store)) {
			defer wg.Done()
			f(store)
		}(fn)
	}

	wg.Wait()
}

func randomUsers(store *db.Store) {
	for i := 0; i < NUM_OF_USERS; i++ {
		createAccountParams := db.CreateAccountParams{
			Name:           util.RandomString(10),
			Email:          fmt.Sprintf("%s@gmail.com", util.RandomString(10)),
			HashedPassword: util.RandomString(100),
		}
		store.CreateAccount(gContext, createAccountParams)

		log.Println("created a new user")
	}
}

func randomBookCategories(store *db.Store) {
	for bookId := 1; bookId <= NUM_OF_BOOKS; bookId++ {
		numberOfCategories := util.RandomInt(3, 5)
		categories := map[int64]bool{}

		for {
			if len(categories) > int(numberOfCategories) {
				break
			}

			newCategory := util.RandomInt(1, NUM_OF_CATEGORIES)
			categories[newCategory] = true
		}

		for i, found := range categories {
			if found {
				params := db.CreateBookCategoryParams{
					BookID:     int32(bookId),
					CategoryID: int32(i),
				}
				store.CreateBookCategory(gContext, params)
				log.Println("created a new book-category map")
			}
		}
	}
}

func randomRates(store *db.Store) {
	for userId := 1; userId <= NUM_OF_USERS; userId++ {
		numberOfRates := util.RandomInt(0, MAX_RATE_PER_USER)
		bookIds := map[int64]bool{}

		for {
			if len(bookIds) > int(numberOfRates) {
				break
			}

			newBook := util.RandomInt(1, NUM_OF_BOOKS)
			bookIds[newBook] = true
		}

		for i, found := range bookIds {
			if found {
				params := db.CreateRateParams{
					BookID:    int32(i),
					CreatedBy: int32(userId),
					RateValue: int32(util.RandomInt(1, 10)),
				}
				store.CreateRate(gContext, params)
				log.Println("created a new rate")
			}
		}
	}
}

func randomBookmarks(store *db.Store) {
	for userId := 1; userId <= NUM_OF_USERS; userId++ {
		numberOfBookmarks := util.RandomInt(0, MAX_BOOKMARKS_PER_USER)
		bookIds := map[int64]bool{}

		for {
			if len(bookIds) > int(numberOfBookmarks) {
				break
			}

			newBook := util.RandomInt(1, NUM_OF_BOOKS)
			bookIds[newBook] = true
		}

		for i, found := range bookIds {
			if found {
				params := db.CreateBookmarkParams{
					BookID:    int32(i),
					CreatedBy: int32(userId),
				}
				store.CreateBookmark(gContext, params)
				log.Println("created a new bookmark")
			}
		}
	}
}

func randomComments(store *db.Store) {
	for bookId := 1; bookId <= NUM_OF_BOOKS; bookId++ {
		numberOfComments := util.RandomInt(0, MAX_COMMENTS_PER_BOOK)

		for i := 0; i < int(numberOfComments); i++ {
			params := db.CreateCommentParams{
				BookID:    int32(bookId),
				CreatedBy: int32(util.RandomInt(1, NUM_OF_USERS)),
				Content:   util.RandomString(100),
			}

			store.CreateComment(gContext, params)
			log.Println("created a new comment")
		}
	}
}

// func addCategory(store *db.Store) {
// 	buff, err := ioutil.ReadFile("sample_data/categories.txt")
// 	if err != nil {
// 		panic("couldn't read data from file")
// 	}
// 	data := string(buff)
// 	lines := strings.Split(data, "\n")

// 	createCategoryParams := []db.CreateCategoryParams{}

// 	for i := 0; i + 1 < len(lines); i += 2 {
// 		name := strings.TrimSpace(lines[i])
// 		description := strings.TrimSpace(lines[i + 1])
		
// 		createCategoryParams = append(createCategoryParams, db.CreateCategoryParams{
// 			Name: name,
// 			Description: description,
// 		})
// 	}
// 	for i := 0; i < len(createCategoryParams); i++ {
// 		store.CreateCategory(gContext, createCategoryParams[i])
// 		log.Println("created a new category")
// 	}

// }
