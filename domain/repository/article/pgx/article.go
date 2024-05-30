package pgx

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"m1-article-service/domain/entity"
	"m1-article-service/infrastructure/godotenv"
)

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

func (r ArticleRepository) Store(ctx context.Context, log *entity.Article) error {
	if _, err := r.conn.Exec(ctx, `INSERT INTO articles (created_at,error) VALUES($1,$2)  `,
		log.CreatedAt, log.Error); err != nil {
		return err
	}
	return nil
}
