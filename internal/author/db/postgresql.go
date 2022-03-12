package author

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"restful_go_project/internal/author"
	"restful_go_project/pkg/client/postgresql"
	"restful_go_project/pkg/logging"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, author *author.Author) error {
	q := `INSERT INTO author (name) 
		  VALUES ($1)
    	  RETURNING id
    	  `

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))
	err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.Is(err, pgErr) {
			pgErr = err.(*pgconn.PgError)
			sqlErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(sqlErr)
			return sqlErr
		}
		return err
	}
	return nil
}

func (r *repository) FindAll(ctx context.Context) (u []author.Author, err error) {
	q := `
		SELECT id, name FROM public.author
	`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	authors := make([]author.Author, 0)

	for rows.Next() {
		var au author.Author

		err = rows.Scan(&au.ID, &au.Name)
		if err != nil {
			return nil, err
		}
		authors = append(authors, au)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (author.Author, error) {
	q := `
		SELECT id, name FROM public.author WHERE id=$1
	`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var au author.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&au.ID, &au.Name)
	if err != nil {
		return author.Author{}, err
	}

	return au, nil
}

func (r *repository) Update(ctx context.Context, user author.Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) author.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
