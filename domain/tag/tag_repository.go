package tag

type TagRepository interface {
	Create(tag *Tag) (*Tag, error)
	Update(tag *Tag) (*Tag, error)
}
