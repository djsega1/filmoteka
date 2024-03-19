package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/djsega1/filmoteka/actors/internal/models"
	"github.com/djsega1/filmoteka/actors/internal/utils"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	Create(ctx context.Context, actor *models.Actor) error
	FindAll(ctx context.Context) (actors []models.Actor, err error)
	FindOne(ctx context.Context, id string) (models.Actor, error)
	Update(ctx context.Context, actor models.Actor) error
	Delete(ctx context.Context, id string) error
}

type ActorsRepository struct {
	client *pgxpool.Pool
}

func (r *ActorsRepository) Create(ctx context.Context, actor *models.Actor) error {
	q := `
		INSERT INTO actors 
		    (name, gender, birthdate) 
		VALUES 
			($1, $2, $3) 
		RETURNING id
	`

	log.Printf("SQL Query: %s", utils.FormatQuery(q))
	if err := r.client.QueryRow(ctx, q, actor.Name, actor.Gender, actor.Birthdate).Scan(&actor.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(
				"SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState(),
			)
			log.Println(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *ActorsRepository) FindAll(ctx context.Context) (actors []models.Actor, err error) {
	q := `
		SELECT id, name, gender, CAST(birthdate AS TEXT) FROM actors;
	`
	log.Printf("SQL Query: %s", utils.FormatQuery(q))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	actors = make([]models.Actor, 0)

	for rows.Next() {
		var actor models.Actor

		log.Println(rows.Values())
		err = rows.Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.Birthdate)
		if err != nil {
			return nil, err
		}

		actors = append(actors, actor)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return actors, nil
}

func (r *ActorsRepository) FindOne(ctx context.Context, id string) (models.Actor, error) {
	q := `
		SELECT id, name, gender, CAST(birthdate AS TEXT) FROM actors WHERE id = $1
	`
	log.Printf("SQL Query: %s", utils.FormatQuery(q))

	var actor models.Actor
	err := r.client.QueryRow(ctx, q, id).Scan(&actor.ID, &actor.Name, &actor.Gender, &actor.Birthdate)
	if err != nil {
		return models.Actor{}, err
	}

	return actor, nil
}

func (r *ActorsRepository) Update(ctx context.Context, actor models.Actor) error {
	//TODO implement me
	panic("implement me")
}

func (r *ActorsRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewActorsRepository(pool *pgxpool.Pool) Repository {
	return &ActorsRepository{
		client: pool,
	}
}
