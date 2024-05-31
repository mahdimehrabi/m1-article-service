package article

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"m1-article-service/domain/entity"
	infraMock "m1-article-service/mock/infrastructure"
	mock_article "m1-article-service/mock/repository"
	"testing"
	"time"
)

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name            string
		article         *entity.Article
		loggerMock      func() *infraMock.MockLog
		articleRepoMock func() *mock_article.MockArticle
		error           error
		ctx             context.Context
	}{
		{
			name: "success",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				return loggerInfra
			},
			articleRepoMock: func() *mock_article.MockArticle {
				repoLogMock := mock_article.NewMockArticle(ctrl)
				repoLogMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
				return repoLogMock
			},
			article: entity.NewArticle("title", "slug", []string{"tag1", "tag2", "tag3"}),
			error:   nil,
			ctx:     context.Background(),
		},
		{
			name: "RepoError",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			articleRepoMock: func() *mock_article.MockArticle {
				repoLogMock := mock_article.NewMockArticle(ctrl)
				repoLogMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(0), err)
				return repoLogMock
			},
			article: entity.NewArticle("title", "slug", []string{"tag1", "tag2", "tag3"}),
			error:   err,
			ctx:     context.Background(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.articleRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			_, err := service.Create(test.ctx, test.article)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_Create(b *testing.B) {
	ctrl := gomock.NewController(b)
	articleRepoMock := mock_article.NewMockArticle(ctrl)
	articleRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
	loggerMock := infraMock.NewMockLog(ctrl)
	b.ResetTimer()
	service := NewService(loggerMock, articleRepoMock)
	service.Create(context.Background(), entity.NewArticle("title", "slug", []string{"tag1", "tag2", "tag3"}))
	fmt.Println(b.Elapsed())
	if b.Elapsed() > 100*time.Microsecond {
		b.Error("article service-create takes too long to run")
	}
	loggerMock.EXPECT()
	articleRepoMock.EXPECT()

}

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")

	var tests = []struct {
		name            string
		article         *entity.Article
		loggerMock      func() *infraMock.MockLog
		articleRepoMock func() *mock_article.MockArticle
		error           error
		ctx             context.Context
	}{
		{
			name: "success",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				return loggerInfra
			},
			articleRepoMock: func() *mock_article.MockArticle {
				repoLogMock := mock_article.NewMockArticle(ctrl)
				repoLogMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
				return repoLogMock
			},
			article: entity.NewArticle("title", "slug", []string{"tag1", "tag2", "tag3"}),
			error:   nil,
			ctx:     context.Background(),
		},
		{
			name: "RepoError",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			articleRepoMock: func() *mock_article.MockArticle {
				repoLogMock := mock_article.NewMockArticle(ctrl)
				repoLogMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(err)
				return repoLogMock
			},
			article: entity.NewArticle("title", "slug", []string{"tag1", "tag2", "tag3"}),
			error:   err,
			ctx:     context.Background(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.articleRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			err := service.Update(test.ctx, test.article)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}

}

func BenchmarkService_Update(b *testing.B) {
	ctrl := gomock.NewController(b)
	articleRepoMock := mock_article.NewMockArticle(ctrl)
	articleRepoMock.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
	loggerMock := infraMock.NewMockLog(ctrl)
	b.ResetTimer()

	service := NewService(loggerMock, articleRepoMock)
	service.Update(context.Background(), entity.NewArticle("title", "slug", []string{"tag1", "tag2", "tag3"}))
	if b.Elapsed() > 100*time.Microsecond {
		b.Error("article service-update takes too long to run")
	}
	loggerMock.EXPECT()
	articleRepoMock.EXPECT()
}

func TestService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")
	var tests = []struct {
		name            string
		id              int64
		loggerMock      func() *infraMock.MockLog
		articleRepoMock func() *mock_article.MockArticle
		error           error
		ctx             context.Context
	}{
		{
			name: "success",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				return loggerInfra
			},
			articleRepoMock: func() *mock_article.MockArticle {
				repoLogMock := mock_article.NewMockArticle(ctrl)
				repoLogMock.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil)
				return repoLogMock
			},
			id:    1,
			error: nil,
			ctx:   context.Background(),
		},
		{
			name: "RepoError",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			articleRepoMock: func() *mock_article.MockArticle {
				repoLogMock := mock_article.NewMockArticle(ctrl)
				repoLogMock.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(err)
				return repoLogMock
			},
			id:    1,
			error: err,
			ctx:   context.Background(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.articleRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			err := service.Delete(test.ctx, test.id)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_Delete(b *testing.B) {
	ctrl := gomock.NewController(b)
	articleRepoMock := mock_article.NewMockArticle(ctrl)
	articleRepoMock.EXPECT().Delete(gomock.Any(), int64(1)).Return(nil)
	loggerMock := infraMock.NewMockLog(ctrl)
	b.ResetTimer()
	service := NewService(loggerMock, articleRepoMock)
	service.Delete(context.Background(), int64(1))
	if b.Elapsed() > 100*time.Microsecond {
		b.Error("article service-delete takes too long to run")
	}
	loggerMock.EXPECT()
	articleRepoMock.EXPECT()

}

func TestService_Detail(t *testing.T) {
	ctrl := gomock.NewController(t)
	t.Cleanup(func() {
		ctrl.Finish()
	})
	err := errors.New("error")
	article := entity.NewArticle("title", "slug", []string{"tag1", "tag2", "tag3"})

	var tests = []struct {
		name            string
		id              int64
		loggerMock      func() *infraMock.MockLog
		articleRepoMock func() *mock_article.MockArticle
		error           error
		ctx             context.Context
		returnedArticle *entity.Article
	}{
		{
			name: "success",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				return loggerInfra
			},
			articleRepoMock: func() *mock_article.MockArticle {
				repoLogMock := mock_article.NewMockArticle(ctrl)
				repoLogMock.EXPECT().Detail(gomock.Any(), gomock.Any()).Return(article, nil)
				return repoLogMock
			},
			id:              1,
			error:           nil,
			ctx:             context.Background(),
			returnedArticle: article,
		},
		{
			name: "RepoError",
			loggerMock: func() *infraMock.MockLog {
				loggerInfra := infraMock.NewMockLog(ctrl)
				loggerInfra.EXPECT().Error(err).Return()
				return loggerInfra
			},
			articleRepoMock: func() *mock_article.MockArticle {
				repoLogMock := mock_article.NewMockArticle(ctrl)
				repoLogMock.EXPECT().Detail(gomock.Any(), gomock.Any()).Return(nil, err)
				return repoLogMock
			},
			id:              1,
			error:           err,
			ctx:             context.Background(),
			returnedArticle: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logRepoMock := test.articleRepoMock()
			loggerMock := test.loggerMock()
			service := NewService(loggerMock, logRepoMock)
			resArticle, err := service.Detail(test.ctx, test.id)
			if !errors.Is(err, test.error) {
				t.Error("error is not equal")
			}

			if !gomock.Eq(resArticle).Matches(test.returnedArticle) {
				t.Error("returned article is not right")
			}
			loggerMock.EXPECT()
			logRepoMock.EXPECT()
		})
	}
}

func BenchmarkService_Detail(b *testing.B) {
	ctrl := gomock.NewController(b)
	articleRepoMock := mock_article.NewMockArticle(ctrl)
	article := entity.NewArticle("title", "slug", []string{"tag1", "tag2", "tag3"})
	articleRepoMock.EXPECT().Detail(gomock.Any(), int64(1)).Return(article, nil)
	loggerMock := infraMock.NewMockLog(ctrl)
	b.ResetTimer()
	service := NewService(loggerMock, articleRepoMock)

	service.Detail(context.Background(), int64(1))
	if b.Elapsed() > 100*time.Microsecond {
		b.Error("article service-detail takes too long to run")
	}
	loggerMock.EXPECT()
	articleRepoMock.EXPECT()
}
