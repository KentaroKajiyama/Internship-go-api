package model

import (
	tagDomain "github.com/KentaroKajiyama/internship-go-api/domain/tag"
)

func (s *Tag) ToDomainTag() tagDomain.Tag {
	return tagDomain.NewTag(s.Id, s.Name, s.CreatedAt, s.UpdatedAt)
}

func (s *Tag) NewTagFromDomainTag(tag *tagDomain.Tag) Tag {
	return Tag{
		Id:        tag.Id(),
		TagId:     tag.TagId(),
		Name:      tag.Name(),
		CreatedAt: tag.CreatedAt(),
		UpdatedAt: tag.UpdatedAt(),
	}
}
