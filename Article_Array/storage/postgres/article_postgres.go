package postgres

import (
	"article/models"
	repository "article/storage"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type articleRepo struct {
	db *sqlx.DB
}

func NewArticleRepository(db *sqlx.DB) repository.ArticleRepository {
	return articleRepo{
		db: db,
	}
}

func (repo articleRepo) Create(entity models.Article) (err error) {
	rows, err := repo.db.Exec(`WITH author AS (
		INSERT INTO author (
			firstname,
		lastname)
		VALUES($1,$2) RETURNING id)
		INSERT INTO article (	title,
		body,author_id) VALUES(
   $3,$4 , (SELECT id FROM author) )`, entity.Author.Firstname, entity.Author.Lastname, entity.Title, entity.Body)
	if err != nil {
		return err
	}
	created_id, err := rows.RowsAffected()
	if created_id == 0 {
		return err
	}
	return nil
}

func (repo articleRepo) GetList(query models.Query) (resp []models.Article, err error) {
	rows, err := repo.db.Query(
		`SELECT
		ar.id,
		ar.title,
		ar.body,
		au.firstname,
		au.lastname,
		ar.created_at FROM article AS ar JOIN author AS au ON ar.author_id = au.id  WHERE ar.deleted_at IS  NULL OFFSET $1 LIMIT $2`,
		query.Offset,
		query.Limit,
	)
	if err != nil {
		return resp, err
	}
	defer rows.Close()
	for rows.Next() {
		article := models.Article{}
		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Body,
			&article.Author.Firstname,
			&article.Author.Lastname,
			&article.CreatedAt)
		if err != nil {
			return resp, err
		}
		resp = append(resp, article)
	}
	return resp, err
}

func (repo articleRepo) GetByID(ID int) (resp models.Article, err error) {
	rows, err := repo.db.Query(`SELECT
			ar.id,
		ar.title,
		ar.body,
		au.firstname,
		au.lastname,
		ar.created_at FROM article AS ar JOIN author AS au ON ar.id = $1 WHERE  ar.deleted_at IS NULL `, ID)
	// getArticle, err := articleStorage.GetByID(id)
	if err != nil {
		return resp, err
	}
	defer rows.Close()
	if !rows.Next() {
		return resp, errors.New("not found id")
	} else {
		var article models.Article
		for rows.Next() {
			err = rows.Scan(
				&article.ID,
				&article.Title,
				&article.Body,
				&article.Author.Firstname,
				&article.Author.Lastname,
				&article.CreatedAt)
			if err != nil {
				return resp, err
			}
			resp = article
		}
	}
	return resp, nil
}

func (repo articleRepo) Update(entity models.Article) (effectedRowsNum int64, err error) {
	tm := time.Now()
	rows, err := repo.db.Exec(`WITH article AS (  UPDATE article SET
			title = $1 ,
			body = $2,
			updated_at = $3
			WHERE id = $4  RETURNING author_id)
			UPDATE author SET firstname = $5,lastname = $6 , updated_at = $7 WHERE id = (SELECT author_id  FROM article )`, entity.Title, entity.Body, tm, entity.ID, entity.Author.Firstname, entity.Author.Lastname, tm)
	if err != nil {
		return 0, err
	}
	effectedRowsNum, err = rows.RowsAffected()
	if err != nil {
		return 0, err
	}
	return effectedRowsNum, err
}

func (repo articleRepo) Delete(ID int) (effectedRowsNum int64, err error) {
	tm := time.Now()
	rows, err := repo.db.Exec(`WITH article AS (  UPDATE article SET
			deleted_at = $1
			WHERE id = $2  RETURNING author_id)
			UPDATE author SET deleted_at = $3 WHERE id = (SELECT author_id  FROM article )`, tm, ID, tm)
	if err != nil {
		return 0, err
	}
	effectedRowsNum, err = rows.RowsAffected()
	if err != nil {
		return 0, err
	}
	return effectedRowsNum, nil
}

func (repo articleRepo) Search(query models.Query) (resp []models.Article, err error) {
	rows, err := repo.db.Query(`
		SELECT
			ar.id,
		ar.title,
		ar.body,
		au.firstname,
		au.lastname,
		ar.created_at FROM article AS ar JOIN author AS au ON ar.author_id = au.id WHERE  ar.deleted_at IS  NULL AND ar.title LIKE  $1 OFFSET $2 LIMIT $3
		`, query.Search, query.Offset, query.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		article := models.Article{}
		err = rows.Scan(
			&article.ID,
			&article.Title,
			&article.Body,
			&article.Author.Firstname,
			&article.Author.Lastname,
			&article.CreatedAt)

		if err != nil {
			return resp, err
		}
		resp = append(resp, article)
	}
	return
}
