package repository

import (
	"context"

	"github.com/KentaroKajiyama/internship-go-api/domain/tag"
)

type tagRepository struct {
}

func NewTagRepository() tag.TagRepository {
	return &tagRepository{}
}

func (r *tagRepository) Create(ctx context.Context, tag *tag.Tag) error {
}

func (r *tagRepository) Update(ctx context.Context, tag *tag.Tag) error {
}

func (r *tagRepository) Delete(ctx context.Context, tag *tag.Tag) error {
}
