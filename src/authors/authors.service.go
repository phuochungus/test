package authors

import (
	"context"
	sqlcGen "root/db/generated"
	idGen "root/src/id_generator"
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

func FindOne(id int64) (author sqlcGen.Author, err error) {
	conn, err := pool.GetConnection()
	defer conn.Release()
	if err != nil {
		return
	}

	q := sqlcGen.New(conn)

	author, err = q.GetAuthor(context.Background(), id)
	return
}

func CreateOne(name string) (author sqlcGen.Author, err error) {
	conn, err := pool.GetConnection()
	defer conn.Release()
	if err != nil {
		return
	}

	id := idGen.GenId()
	q := sqlcGen.New(conn)

	author, err = q.CreateAuthor(context.Background(), sqlcGen.CreateAuthorParams{ID: id, Name: name})
	return
}

func UpdateAuthor(id int64, name string) (author sqlcGen.Author, err error) {
	conn, err := pool.GetConnection()
	defer conn.Release()
	if err != nil {
		return
	}
	q := sqlcGen.New(conn)
	author, err = q.UpdateAuthor(context.Background(), sqlcGen.UpdateAuthorParams{ID: id, Name: name})
	return
}

func DeleteAuthor(id int64) (err error) {
	conn, err := pool.GetConnection()
	defer conn.Release()
	if err != nil {
		return
	}
	q := sqlcGen.New(conn)
	err = q.DeleteAuthor(context.Background(), id)
	return
}
