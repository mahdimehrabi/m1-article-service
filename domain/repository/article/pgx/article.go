package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"m1-article-service/domain/entity"
	"m1-article-service/infrastructure/godotenv"
)

const pageSize = 10

type ArticleRepository struct {
	env  *godotenv.Env
	conn *pgxpool.Pool
}

func NewArticleRepository(env *godotenv.Env, conn *pgxpool.Pool) *ArticleRepository {
	lr := &ArticleRepository{
		env:  env,
		conn: conn,
	}
	return lr
}

func (r ArticleRepository) Create(ctx context.Context, article *entity.Article) (int64, error) {
	sql := `INSERT INTO articles (title,slug,tags,created_at) VALUES($1,$2,$3,$4) RETURNING id`
	err := r.conn.
		QueryRow(ctx, sql,
			article.Title, article.Slug, article.Tags, article.CreatedAt).Scan(&article.ID)
	if err != nil {
		return 0, err
	}
	return article.ID, err
}

func (r ArticleRepository) Update(ctx context.Context, article *entity.Article) error {
	if _, err := r.conn.Exec(ctx, `UPDATE articles SET title=$1,slug=$2,tags=$3,created_at=$4 WHERE id=$5`,
		article.Title, article.Slug, article.Tags, article.CreatedAt, article.ID); err != nil {
		return err
	}
	return nil
}

func (r ArticleRepository) Delete(ctx context.Context, id int64) error {
	if _, err := r.conn.Exec(ctx, `DELETE FROM articles WHERE id=$1`, id); err != nil {
		return err
	}
	return nil
}

func (r ArticleRepository) Detail(ctx context.Context, id int64) (article *entity.Article, err error) {
	err = r.conn.QueryRow(ctx, `SELECT * FROM articles WHERE id=$1`, id).Scan(&article)
	if err != nil {
		return nil, err
	}
	return
}

func (r ArticleRepository) List(ctx context.Context, pageNumber uint16) ([]*entity.Article, error) {
	articles := make([]*entity.Article, 0)
	offset := (pageNumber - 1) * pageSize
	rows, err := r.conn.Query(ctx, `SELECT * FROM articles LIMIT $1 OFFSET $2 `, pageSize, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		article := &entity.Article{}
		if err := rows.Scan(article); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	if rows.Err() != nil {
		return nil, err
	}
	return articles, nil
}
