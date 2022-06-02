package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"web-bootcamp-dh/internal/products"
	"web-bootcamp-dh/pkg/web"
)

type Controller interface {
	GetAll() gin.HandlerFunc
	Save() gin.HandlerFunc
	Update() gin.HandlerFunc
	Delete() gin.HandlerFunc
	UpdateSobrenomeIdade() gin.HandlerFunc
}
type userController struct {
	service products.Service
}

func NewUserController(service products.Service) Controller {
	return &userController{service}
}

// GetAll
// @Summary Get all Users
// @Tags Users
// @Description Get Users list
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /user [get]
func (c *userController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, web.NewErrorResp(http.StatusUnauthorized, err.Error()))
			return
		}

		users, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewErrorResp(http.StatusNotFound, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewSuccesResp(http.StatusOK, users))
	}
}

func (c *userController) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req products.Request
		err = ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = validParams(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, err := c.service.Save(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, user)
	}
}

func (c *userController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req products.Request
		err = ctx.Bind(&req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = validParams(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, err := c.service.Update(int(id), req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}

func (c *userController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Apagado com sucesso",
		})
	}
}

func (c *userController) UpdateSobrenomeIdade() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validToken(ctx.GetHeader("token"))
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		var req products.Request
		err = ctx.Bind(&req) // nao esta recebendo os valores da requisicao
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		user, err := c.service.UpdateSobrenomeIdade(int(id), req.Sobrenome, req.Idade)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, user)
	}
}

func validToken(token string) error {
	if token != os.Getenv("TOKEN") {
		return fmt.Errorf("token incorreto")
	}
	return nil
}

func validParams(req products.Request) error {
	if req.Nome == "" {
		return fmt.Errorf("o campo nome não pode ser vazio")
	}
	if req.Sobrenome == "" {
		return fmt.Errorf("o campo sobrenome não pode ser vazio")
	}
	if req.Email == "" {
		return fmt.Errorf("o campo email não pode ser vazio")
	}
	if req.Idade == 0 {
		return fmt.Errorf("o campo idade não pode ser vazio")
	}
	if req.Altura == 0 {
		return fmt.Errorf("o campo altura não pode ser vazio")
	}
	if req.Ativo == false {
		return fmt.Errorf("o campo ativo não pode ser vazio")
	}
	if req.DataCriacao == "" {
		return fmt.Errorf("o campo data de criacao não pode ser vazio")
	}
	return nil
}
