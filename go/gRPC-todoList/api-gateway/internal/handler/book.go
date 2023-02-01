package handler

import (
	"api-gateway/internal/service"
	"api-gateway/pkg/res"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

// GetBookList 获取所有书籍
func GetBookList(ginCtx *gin.Context) {
	var bookReq service.BookListRequest
	PanicIfBookError(ginCtx.Bind(&bookReq))
	bookService := ginCtx.Keys["book"].(service.BookControllerClient)
	bookStream, err := bookService.List(context.Background(), &bookReq)
	fmt.Println(bookStream)
	PanicIfBookError(err)
	var bookList []*service.Book
	for {
		resp, err := bookStream.Recv()
		if err == nil {
			bookList = append(bookList, resp)
		}
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		if err != nil {
			log.Println(err)
		}
	}
	r := res.Response{
		Data:   bookList,
		Status: 200,
		Msg:    "成功",
	}
	ginCtx.JSON(http.StatusOK, r)
}
