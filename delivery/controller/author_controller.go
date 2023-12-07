package controller

import (
	"gin/config"
	"gin/shared/common"
	"gin/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorController struct {
	authorUC usecase.AuthorUseCase
	rg       *gin.RouterGroup
}

func (a *AuthorController) listHandler(ctx *gin.Context) {
	authors, err := a.authorUC.FindAllAuthor()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range authors {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, "Ok")
}

func (a *AuthorController) getHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	author, err := a.authorUC.FindAuthorByID(id)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "author with ID "+id+" not found")
		return
	}
	common.SendSingleResponse(ctx, author, "Ok")
}

func (a *AuthorController) Route() {
	a.rg.GET(config.AuthorGetList, a.listHandler)
	a.rg.GET(config.AuthorGetById, a.getHandler)
}

func NewAuthorController(authorUC usecase.AuthorUseCase, rg *gin.RouterGroup) *AuthorController {
	return &AuthorController{
		authorUC: authorUC,
		rg:       rg,
	}
}
