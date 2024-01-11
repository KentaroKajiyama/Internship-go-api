package model

import (
	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
)

func (s *Tag) ToDomainTag() (*tagDomain.Tag, error) {
	if s.CreatedAt.IsZero() {
		return tagDomain.NewTag(s.ID, s.Name)
	} else {
		return tagDomain.ReconstructTag(s.ID, s.TagID, s.Name, s.CreatedAt)
	}
}

func NewTagFromDomainTag(tag *tagDomain.Tag) Tag {
	return Tag{
		ID:        tag.ID(),
		TagID:     tag.TagID(),
		Name:      tag.Name(),
		CreatedAt: tag.CreatedAt(),
		UpdatedAt: tag.UpdatedAt(),
	}
}
