package authors

import (
	"context"
	"errors"
	"net/http"
	sqlcGen "root/db/generated"
	idGen "root/src/id_generator"
	"root/src/middleware"
	"root/src/pool"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreateController(r *gin.Engine) {
	c := r.Group("/authors")
	{
		c.GET("", findAll)
		c.POST("", middleware.ValidateBody[sqlcGen.CreateAuthorParams](), createOne)
	}
}

func findAll(ctx *gin.Context) {
	authors, err := FindAll()
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
	} else {
		ctx.JSON(http.StatusOK, authors)
	}
}

func createOne(ctx *gin.Context) {
	conn, err := pool.GetConnection()
	defer conn.Release()
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	q := sqlcGen.New(conn)

	createAuthor, ok := ctx.Value("body").(sqlcGen.CreateAuthorParams)
	if !ok {
		panic("error convert to type sqlcGen.CreateAuthorParams")
	}
	createAuthor.ID = int64(idGen.GenId())
	result, err := q.CreateAuthor(context.Background(), createAuthor)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"message": "Author already in exist",
				})

			}

		}
		ctx.AbortWithStatus(http.StatusBadGateway)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}
