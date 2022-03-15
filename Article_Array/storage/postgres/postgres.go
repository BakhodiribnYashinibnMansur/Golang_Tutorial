package postgres

import (
	"article/storage"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db      *sqlx.DB
	article storage.ArticleRepository
	author  storage.AuthorRepository
}

func NewPostgres(psqlConnString string) storage.Storage {
	db, err := sqlx.Connect("postgres", psqlConnString)
	if err != nil {
		log.Panic(err)
	}
	return &Store{
		db: db,
	}
}

func (s *Store) Article() storage.ArticleRepository {
	if s.article == nil {
		s.article = &articleRepo{db: s.db}
	}
	return s.article
}

func (s *Store) Author() storage.AuthorRepository {
	if s.author == nil {
		s.author = &authorRepo{db: s.db}
	}
	return s.author
}
