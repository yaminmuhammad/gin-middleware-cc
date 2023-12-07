package controller

import (
	"gin/config"
	"gin/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorUC usecase.AuthorUseCase
	rg       *gin.RouterGroup
}

func (a *AuthorController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	author, err := a.authorUC.FindAuthorByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok", "data": author})
}

func (a *AuthorController) listHandler(ctx *gin.Context) {
	authors, err := a.authorUC.FindAllAuthor()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Ok", "data": authors})
}

func (a *AuthorController) Route() {
	a.rg.GET(config.AuthorGetById, a.getHandler)
	a.rg.GET(config.AuthorGetList, a.listHandler)
}

func NewAuthorController(authorUC usecase.AuthorUseCase, rg *gin.RouterGroup) *AuthorController {
	return &AuthorController{
		authorUC: authorUC,
		rg:       rg,
	}
}
