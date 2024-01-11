package tag

import (
	"context"
)

type TagRepository interface {
	// Find でCreatedAtをどうするか？
	Find(ctx context.Context, id string, tag_id int) (*Tag, error)
	Create(ctx context.Context, Tag *Tag) error
	Update(ctx context.Context, Tag *Tag) error
	Delete(ctx context.Context, Tag *Tag) error
}
