package entity

import "time"

type Article struct {
	ID        int64    `json:"ID"`
	Title     string   `json:"title"`
	Slug      string   `json:"slug"`
	Tags      []string `json:"tags"`
	CreatedAt uint64   `json:"createdAt"`
}

func NewArticle(title string, slug string, tags []string) *Article {
	return &Article{Title: title, Slug: slug, Tags: tags,
		CreatedAt: uint64(time.Now().Unix()),
	}
}
