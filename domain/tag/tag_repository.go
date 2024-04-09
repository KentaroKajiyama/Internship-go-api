package tag

import (
	"context"
)

type TagRepository interface {
	// Find でCreatedAtをどうするか？
	Find(ctx context.Context, id string, tag_id uint) (*Tag, error)
	FindMultiple(ctx context.Context, id string, tag_id uint, name string) ([]*Tag, error)
	Create(ctx context.Context, Tag *Tag) (*Tag, error)
	Update(ctx context.Context, Tag *Tag) (*Tag, error)
	Delete(ctx context.Context, Tag *Tag) (*Tag, error)
}
