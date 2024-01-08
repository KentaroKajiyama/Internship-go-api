package tag

import (
	"context"

	"github.com/google/uuid"
)

type TagRepository interface {
	Find(ctx context.Context, id uuid.UUID, tag_id int) (*Tag, error)
	Create(ctx context.Context, Tag *Tag) error
	Update(ctx context.Context, Tag *Tag) error
	Delete(ctx context.Context, Tag *Tag) error
}
