package api

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/meilisearch/meilisearch-go"
	db "github.com/ndtai772/MyBookListBackend/db/sqlc"
)

func (server *Server) listBooks(ctx *gin.Context) {
	var req struct {
		Limit  int    `form:"page_size,default=50"`
		Offset int    `form:"offset,default=0"`
		Query  string `form:"q"`
	}

	if err := ctx.ShouldBindWith(&req, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	q, _ := url.QueryUnescape(req.Query)

	res, _ := server.bookIndex.Search(q, &meilisearch.SearchRequest{
		Limit:  int64(req.Limit),
		Offset: int64(req.Offset),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"data":   res.Hits,
		"offset": req.Offset,
		"page_size":  req.Limit,
	})
}

func (server *Server) searchBooks(ctx *gin.Context) {
	var req struct {
		Limit  int    `form:"limit" binding:"required"`
		Offset int    `form:"offset" binding:"required"`
		Query  string `form:"q" binding:"required"`
	}

	if err := ctx.ShouldBindWith(&req, binding.Query); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	q, _ := url.QueryUnescape(req.Query)

	res, _ := server.bookIndex.Search(q, &meilisearch.SearchRequest{
		Limit:  int64(req.Limit),
		Offset: int64(req.Offset),
	})

	ctx.JSON(http.StatusOK, gin.H{
		"data":   res.Hits,
		"offset": req.Offset,
		"limit":  req.Limit,
	})

}

// func (server *Server) createBook(ctx *gin.Context) {
// 	var createBookRequest struct {
// 		Title       string  `form:"title" binding:"required"`
// 		Author      string  `form:"author" binding:"required"`
// 		Description string  `form:"description" binding:"required"`
// 		Categories  []int32 `from:"categories" binding:"required"`
// 	}

// 	if err := ctx.ShouldBindWith(&createBookRequest, binding.Form); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	params := db.CreateBookParams{
// 		Title:       createBookRequest.Title,
// 		Author:      createBookRequest.Author,
// 		Description: createBookRequest.Description,
// 	}

// 	book, err := server.store.CreateBook(ctx, params)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	// TODO: bad practice
// 	for i := 0; i < len(createBookRequest.Categories); i++ {
// 		params := db.CreateBookCategoryParams{
// 			BookID:     book.ID,
// 			CategoryID: createBookRequest.Categories[i],
// 		}

// 		_, err := server.store.CreateBookCategory(ctx, params)
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 			return
// 		}
// 	}

// 	ctx.JSON(http.StatusOK, book)
// }

func (server *Server) getBook(ctx *gin.Context) {
	id, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	book, err := server.store.GetBook(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// func (server *Server) updateBook(ctx *gin.Context) {
// 	id, err := parseIdUri(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	var updateBookRequest struct {
// 		Title       string  `form:"title,omitempty"`
// 		Author      string  `form:"author,omitempty"`
// 		Description string  `form:"description,omitempty"`
// 		Categories  []int32 `form:"categories,omitempty"`
// 	}

// 	if err := ctx.ShouldBindWith(updateBookRequest, binding.Form); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	book, err := server.store.UpdateBook(ctx, db.UpdateBookParams{
// 		ID:          id,
// 		Title:       updateBookRequest.Title,
// 		Author:      updateBookRequest.Author,
// 		Description: updateBookRequest.Description,
// 	})

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, book)
// }

// func (server *Server) deleteBook(ctx *gin.Context) {
// 	id, err := parseIdUri(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	if err := server.store.DeleteBook(ctx, id); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, nil)
// }

// func (server *Server) listBookRates(ctx *gin.Context) {
// 	bookId, err := parseIdUri(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	page_size, last_id, err := parsePaginateQuery(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	rates, err := server.store.ListRatesByBookId(ctx, db.ListRatesByBookIdParams{
// 		Limit:  limit,
// 		Offset: offset,
// 		BookID: bookId,
// 	})

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data":       rates,
// 		"next_index": offset + int32(len(rates)),
// 	})
// }

func (server *Server) listBookComments(ctx *gin.Context) {
	bookId, err := parseIdUri(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	page_size, last_id, err := parsePaginateQuery(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	comments, err := server.store.ListCommentsByBookId(ctx, db.ListCommentsByBookIdParams{
		PageSize: page_size,
		LastID: last_id,
		BookID: bookId,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId := -1
	if len(comments) > 0 {
		lastId = int(comments[len(comments)-1].ID)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":      comments,
		"page_size": page_size,
		"last_id":   lastId,
	})
}
