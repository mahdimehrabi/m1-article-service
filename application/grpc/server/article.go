package server

import (
	"context"
	"errors"
	"github.com/mahdimehrabi/m1-article-proto/gen/go/article/article"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"m1-article-service/domain/entity"
	articleRepo "m1-article-service/domain/repository/article"
	"m1-article-service/domain/service/article"
	logger "m1-article-service/infrastructure/log"
)

type ArticleServer struct {
	logger         logger.Logger
	articleService *article.Service
	articlev1.UnimplementedArticleServiceServer
}

func NewArticleServer(logger logger.Logger, articleService *article.Service) *ArticleServer {
	return &ArticleServer{logger: logger, articleService: articleService}
}

func (a ArticleServer) Create(ctx context.Context, a2 *articlev1.Article) (*articlev1.ArticleCreateResponse, error) {
	article := entity.NewArticle(a2.Title, a2.Slug, a2.Tags)
	id, err := a.articleService.Create(ctx, article)

	if errors.Is(err, articleRepo.ErrAlreadyExist) {
		return nil, status.Errorf(codes.AlreadyExists, "article with this title already exists")
	} else if err != nil {
		a.logger.Error(err)
		return nil, err
	}

	return &articlev1.ArticleCreateResponse{
		ID: id,
	}, nil
}

func (a ArticleServer) Update(ctx context.Context, a2 *articlev1.Article) (*articlev1.ArticleUpdateResponse, error) {
	article := entity.NewArticle(a2.Title, a2.Slug, a2.Tags)
	err := a.articleService.Update(ctx, article)
	if errors.Is(err, articleRepo.ErrAlreadyExist) {
		return nil, status.Errorf(codes.AlreadyExists, "article with this title already exists")
	} else if errors.Is(err, articleRepo.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "article not found")
	} else if err != nil {
		a.logger.Error(err)
		return nil, err
	}

	return &articlev1.ArticleUpdateResponse{}, nil
}

func (a ArticleServer) Delete(ctx context.Context, id *articlev1.ArticleID) (*articlev1.Empty, error) {
	err := a.articleService.Delete(ctx, id.ID)
	if errors.Is(err, articleRepo.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "article not found")
	} else if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	return &articlev1.Empty{}, nil
}

func (a ArticleServer) Detail(ctx context.Context, id *articlev1.ArticleID) (*articlev1.ArticleDetailResponse, error) {
	article, err := a.articleService.Detail(ctx, id.ID)
	if errors.Is(err, articleRepo.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, "article not found")
	} else if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	return &articlev1.ArticleDetailResponse{
		Article: &articlev1.Article{
			ID:    article.ID,
			Title: article.Title,
			Tags:  article.Tags,
			Slug:  article.Slug,
		},
	}, nil
}

func (a ArticleServer) List(ctx context.Context, pagination *articlev1.Pagination) (*articlev1.ArticleListResponse, error) {
	articles, err := a.articleService.List(ctx, uint16(pagination.Page))
	if err != nil {
		a.logger.Error(err)
		return nil, err
	}
	articlesResObjs := make([]*articlev1.Article, len(articles))
	for i, article := range articles {
		articlesResObjs[i] = &articlev1.Article{
			ID:    article.ID,
			Title: article.Title,
			Tags:  article.Tags,
			Slug:  article.Slug,
		}
	}
	return &articlev1.ArticleListResponse{
		Article: articlesResObjs,
	}, nil
}
