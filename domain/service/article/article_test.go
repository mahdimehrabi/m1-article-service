package article

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"m1-article-service/domain/entity"
	infraMock "m1-article-service/mock/infrastructure"
	mock_article "m1-article-service/mock/repository"
	"testing"
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

func BenchmarkService_Store(b *testing.B) {
	ctrl := gomock.NewController(b)
	articleRepoMock := mock_article.NewMockArticle(ctrl)
	articleRepoMock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(int64(1), nil)
	loggerMock := infraMock.NewMockLog(ctrl)

	service := NewService(loggerMock, articleRepoMock)
	service.Create(context.Background(), entity.NewArticle("title", "slug", []string{"tag1", "tag2", "tag3"}))

	loggerMock.EXPECT()
	articleRepoMock.EXPECT()
}
