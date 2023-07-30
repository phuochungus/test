package authors

import (
	"context"
	sqlcGen "root/db/generated"
	"root/src/pool"
)

func FindAll() (authors []sqlcGen.Author, err error) {
	conn, err := pool.GetConnection()
	defer conn.Release()
	if err != nil {
		return
	}
	q := sqlcGen.New(conn)
	authors, err = q.ListAuthors(context.Background())
	return
}

func CreateOne(createDto sqlcGen.CreateAuthorParams) {

}
