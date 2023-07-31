package authors

import (
	"errors"
	"net/http"
	sqlcGen "root/db/generated"
	"root/src/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

type QueryAuthorDTO struct {
	ID int64 `uri:"id" binding:"required,gt=0"`
}

type UpdateAuthorDTO struct {
	Name string `binding:"required" json:"name"`
}

func CreateController(r *gin.Engine) {
	c := r.Group("/authors")
	{
		c.GET("", findAll)
		c.GET("/:id", middleware.ValidatePathParam[QueryAuthorDTO](), findOne)
		c.POST("", middleware.ValidateBody[sqlcGen.CreateAuthorParams](), createOne)
		c.PATCH("/:id", middleware.ValidatePathParam[QueryAuthorDTO](), middleware.ValidateBody[UpdateAuthorDTO](), updateOne)
		c.DELETE("/:id", middleware.ValidatePathParam[QueryAuthorDTO](), deleteOne)
	}
}

func findAll(ctx *gin.Context) {
	authors, err := FindAllAuthors()
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	} else {
		ctx.JSON(http.StatusOK, authors)
	}
}

func createOne(ctx *gin.Context) {
	createAuthorDTO, ok := ctx.Value("body").(sqlcGen.CreateAuthorParams)
	if !ok {
		panic("error convert to type sqlcGen.CreateAuthorParams")
	}
	result, err := CreateAuthor(createAuthorDTO.Name)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				ctx.JSON(http.StatusConflict, gin.H{
					"message": "Author already in exist",
				})

			}

		}
		ctx.AbortWithStatus(http.StatusBadGateway)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func findOne(ctx *gin.Context) {
	queryAuthorDTO, ok := ctx.Value("params").(QueryAuthorDTO)
	if !ok {
		panic("error convert to type sqlcGen.CreateAuthorParams")
	}
	author, err := FindAuthorById(queryAuthorDTO.ID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"msg": "author not found",
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, author)
}

func updateOne(ctx *gin.Context) {
	queryAuthorDTO, ok := ctx.Value("params").(QueryAuthorDTO)
	if !ok {
		panic("error convert to type QueryAuthorDTO")
	}
	updateAuthorDTO, ok := ctx.Value("body").(UpdateAuthorDTO)
	if !ok {
		panic("error convert to type UpdateAuthorParams")
	}

	author, err := UpdateAuthor(queryAuthorDTO.ID, updateAuthorDTO.Name)
	if err != nil {
		if err.Error() == "no rows in result set" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"msg": "author not found",
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, author)
}

func deleteOne(ctx *gin.Context) {
	queryAuthorDTO, ok := ctx.Value("params").(QueryAuthorDTO)
	if !ok {
		panic("error convert to type QueryAuthorDTO")
	}
	err := DeleteAuthor(queryAuthorDTO.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, err.Error())
		return
	}
	ctx.AbortWithStatus(http.StatusOK)
}
