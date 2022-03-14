package postgresql

import (
	"context"
	"errors"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"restful_go_project/internal/author/model"
	"restful_go_project/internal/author/storage"
	model2 "restful_go_project/internal/author/storage/model"
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

func (r *repository) Create(ctx context.Context, author *model.Author) error {
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

func (r *repository) FindAll(ctx context.Context, sortOption model2.SortOptions) (u []model.Author, err error) {
	qb := sq.Select("id, name, age, is_alive, created_at").From("public.author")
	if sortOption.Field != "" && sortOption.Order != "" {
		qb = qb.OrderBy(fmt.Sprintf("%s %s", sortOption.Field, sortOption.Order))
	}
	sql, i, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(sql)))

	rows, err := r.client.Query(ctx, sql, i...)
	if err != nil {
		return nil, err
	}

	authors := make([]model.Author, 0)

	for rows.Next() {
		var au model.Author

		err = rows.Scan(&au.ID, &au.Name, &au.Age, &au.IsAlive, &au.CreatedAt)
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

func (r *repository) FindOne(ctx context.Context, id string) (model.Author, error) {
	q := `
		SELECT id, name FROM public.author WHERE id=$1
	`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var au model.Author
	err := r.client.QueryRow(ctx, q, id).Scan(&au.ID, &au.Name)
	if err != nil {
		return model.Author{}, err
	}

	return au, nil
}

func (r *repository) Update(ctx context.Context, user model.Author) error {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) storage.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
