package repository

import (
	"context"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure/database/model"
	"gorm.io/gorm"
)

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) tagDomain.TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) Find(ctx context.Context, id string, tag_id uint64) (*tagDomain.Tag, error) {
	conn := r.db.WithContext(ctx)
	var tagModel model.Tag
	var tagDomainPtr *tagDomain.Tag
	var errDom error
	if err := conn.Where("id = ? AND tag_id = ?", id, tag_id).Find(&tagModel).Error; err != nil {
		return nil, err
	}
	//infra層からdomain層へ
	tagDomainPtr, errDom = tagModel.ToDomainTag()
	if errDom != nil {
		return nil, errDom
	}
	return tagDomainPtr, nil
}

func (r *tagRepository) FindByTodoId(ctx context.Context, id string, todo_id string, name string) ([]*tagDomain.Tag, error) {
	var tagsModel []model.Tag
	var tags []*tagDomain.Tag
	var tag *tagDomain.Tag
	var errDom error
	conn := r.db.WithContext(ctx)
	if id != "" {
		conn = conn.Where("id = ?", id)
	}
	if todo_id != "" {
		conn = conn.Joins("JOIN todo_tag ON todo_tag.tag_id = tags.tag_id").Where("todo_id = ?", todo_id)
	}
	if name != "" {
		conn = conn.Where("name = ?", name)
	}
	if err := conn.Find(&tagsModel).Error; err != nil {
		return nil, err
	}
	//from infra to domain
	for _, tm := range tagsModel {
		tag, errDom = tm.ToDomainTag()
		if errDom != nil {
			return nil, errDom
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r *tagRepository) Create(ctx context.Context, tag *tagDomain.Tag) (*tagDomain.Tag, error) {
	conn := r.db.WithContext(ctx)
	// domain層からinfra層へ
	tagModel := model.NewTagFromDomainTag(tag)
	if err := conn.Create(&tagModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったtodo(domain)を返す。
	return tagModel.ToDomainTag()
}

// 修正が必要かも
func (r *tagRepository) Update(ctx context.Context, tag *tagDomain.Tag) (*tagDomain.Tag, error) {
	conn := r.db.WithContext(ctx)
	tagModel := model.NewTagFromDomainTag(tag)
	if err := conn.Model(&model.Tag{}).Where("id = ? AND tag_id = ?", tag.Id(), tag.TagId()).Updates(&tagModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったtodo(domain)を返す。
	return tagModel.ToDomainTag()
}

func (r *tagRepository) Delete(ctx context.Context, tag *tagDomain.Tag) (*tagDomain.Tag, error) {
	conn := r.db.WithContext(ctx)
	tagModel := model.NewTagFromDomainTag(tag)
	if err := conn.Where("id = ? AND tag_id = ?", tag.Id(), tag.TagId()).Delete(&tagModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったtodo(domain)を返す。
	return tagModel.ToDomainTag()
}

func (r *tagRepository) DeleteMultiple(ctx context.Context, id string, tag_id_s []uint64) (*tagDomain.Tags, error) {
	conn := r.db.WithContext(ctx)

	// transaction?
	// In gorm, you have to pass pointer to gorm methods.
	tx := conn.Begin()
	if err := conn.Where("id = ? AND tag_id IN(?)", id, tag_id_s).Delete(&model.Tag{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return tagDomain.NewTags(id, tag_id_s)
}
