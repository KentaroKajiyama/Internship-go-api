package tag

import (
	"context"
)

type TagRepository interface {
	// Find でCreatedAtをどうするか？
	Find(ctx context.Context, id string, tag_id uint64) (*Tag, error)
	FindByTodoId(ctx context.Context, id string, todo_id string, name string) ([]*Tag, error)
	Create(ctx context.Context, Tag *Tag) (*Tag, error)
	Update(ctx context.Context, Tag *Tag) (*Tag, error)
	Delete(ctx context.Context, Tag *Tag) (*Tag, error)
	DeleteMultiple(ctx context.Context, id string, tag_ids []uint64) (*Tags, error)
}
