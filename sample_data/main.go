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
	NUM_OF_USERS           = 100
	NUM_OF_CATEGORIES      = 25
	NUM_OF_BOOKS           = 1000
	MAX_RATE_PER_USER      = NUM_OF_BOOKS / 4
	MAX_BOOKMARKS_PER_USER = NUM_OF_BOOKS / 4
	MAX_COMMENTS_PER_BOOK  = 20
)

func Gen(store *db.Store) {
	addCategory(store)

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
		go func(s *db.Store) {
			defer wg.Done()
			fn(s)
		}(store)
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
		numberOfCategories := util.RandomInt(5, 10)
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

func addCategory(store *db.Store) {
	createCategoryParams := []db.CreateCategoryParams{
		{
			Name:        "Action",
			Description: "Traditionally (but no longer exclusively) aimed at male readers, features physical action and violence, usually around a quest or mission set in exotic locations such as deserts, jungles, or mountains. Stories often include elements of technology, weapons, and other hardware. The genre is still largely aimed towards males (for example, the James Bond series is often categorized as action), but many later works feature women in roles more traditionally reserved for males, such as heroic bodyguards.",
		},
		{
			Name:        "Adventure",
			Description: "Books whose characters undergo some kind of quest or task to accomplish a goal, often a journey, and usually dangerous.",
		},
		{
			Name:        "Alternate History",
			Description: "Discussion of what might have been, given different decisions and actions at key points in history -- if Hitler had won WWII, or if Hitler hadn't ever come to power, for example.",
		},
		{
			Name:        "Androids, Cyborgs & Robots",
			Description: "Stories in which the main story, a major sub-theme, or one (or more) of the central characters, involves cybernetic limb replacement, anthropomorphic (i.e., human-like) androids, or autonomous robots.",
		},
		{
			Name:        "Angels & Demons",
			Description: "Stories that involve immortal spirits interacting with mortals, and possibly even assuming physical form. Usually, but not always, involves a religious theme, or morality play. Also included here are lesser spirits given mortally created form, such as enchanted constructs (golems), incubated homunculi, and other types of magical mortal familiars inhabited by immortal spirits.",
		},
		{
			Name:        "Animal",
			Description: "Books which have animals as the main characters, or whose plots revolve around animals. For example, the Redwall series by Brian Jacques.",
		},
		{
			Name:        "Arthurian Legends",
			Description: "Stories relating in some way to the mythology surrounding King Arthur and his Knights of the Round Table. May be historical, may be fantasy (or neither), but has Arthur as a central theme.",
		},
		{
			Name:        "Artificial Intelligence",
			Description: "Stories in which the main story, a major sub-theme, or one (or more) of the central characters, involves or explores the role and/or nature of computer intelligence. Examples: 2001 (HAL), Number of the Beast (Gay Deceiver), etc.",
		},
		{
			Name:        "Autobiography/ Memoirs",
			Description: "Non-Fiction books about the lives of the author that wrote them.",
		},
		{
			Name:        "Avante Garde & Surreal",
			Description: "Books having a deliberately non-traditional experimental story structure, and/or highly unusual or surreal writing style.",
		},
		{
			Name:        "Biographical",
			Description: "Nonfiction books about the lives of real people.",
		},
		{
			Name:        "Children",
			Description: "Books intended either to be read to children, or to be read by children who are still learning to read well, up to perhaps age 8 or so.",
		},
		{
			Name:        "Christmas",
			Description: "Works about the traditional Christian holy day.",
		},
		{
			Name:        "Classic",
			Description: "Books that are recognizable by a typical high-school grad, whether they've read them or not. Examples include War & Peace, Great Expectations, Of Mice & Men, The Pearl.",
		},
		{
			Name:        "Coming of Age",
			Description: "Works that have an emphasis on the journey from youth to adulthood.",
		},
		{
			Name:        "Computers",
			Description: "Books about computer software packages, computer hardware, computer languages and operating systems, various computer applications, the internet, and stories about same.",
		},
		{
			Name:        "Contemporary",
			Description: "The setting is in the modern world, filled with the hustle and bustle of our everyday lives.",
		},
		{
			Name:        "Crime & Prison",
			Description: "Books in which the primary focus involves crime, criminals, gangs, the mafia, etc.",
		},
		{
			Name:        "Criticism & Commentary",
			Description: "A book whose main purpose is to interpret, analyze, or otherwise opine about the work of another author, real life events (past, present or projected), a body of thought, a belief system, etc",
		},
		{
			Name:        "Death, Dying and Grieving",
			Description: "Literature which explores the process of death, dying, and grieving from the perspective of the person experiencing it and/or the perspective of those close to him/her.",
		},
		{
			Name:        "Decision Tree",
			Description: "Includes those \"Choose Your Own Adventure\" books you may remember. The reader chooses which parts of the book to read to form a story based on the decisions they make for the characters.",
		},
		{
			Name:        "Educational, Instructional, Reference",
			Description: "Books that focus on educating/instructing the reader on a particular subject (i.e., how-to books), or which are intended to serve as an information reference (i.e., encyclopedias, dictionaries, charts & tables, etc).",
		},
		{
			Name:        "Epic Saga",
			Description: "A book, or often series of books, with far-reaching storylines that involve an entire region or world (or more). Usually, characters in such sagas have great effect on their surroundings. The books or series themselves tend to be long, and often dense reading.",
		},
		{
			Name:        "Erotic",
			Description: "Stories in which sex and/or sensuality play a major role. Not necessarily romance, as love is not always included. Often for mature audiences only.",
		},
		{
			Name:        "Espionage",
			Description: "Books in which the primary focus involves spying and spycraft, of any flavor (military, corporate, international, etc.)",
		},
		{
			Name:        "Ethics and Morality",
			Description: "Books that deal with Ethics (the discipline dealing with what is good and bad) or Morality (referring to a code of conduct put forward by a society or, some other group, such as a religion, or accepted by an individual for their own behavior or given specified conditions, would be put forward by all rational persons. ) or have those as major themes.",
		},
		{
			Name:        "Ethnic: Native American",
			Description: "Works which come from the history and legends of Native Americans, and works that explain and expound upon the same, as well as work written by Native Americans.",
		},
		{
			Name:        "Fairy Tales",
			Description: "Books which deal heavily with well-known children's wonder tales such as Cinderella, Snow White, etc. Also includes stories which feature fairies, faeries, and fae as primary characters or plot devices.",
		},
	}

	for i := 0; i < len(createCategoryParams); i++ {
		store.CreateCategory(gContext, createCategoryParams[i])
		log.Println("created a new category")
	}

}
