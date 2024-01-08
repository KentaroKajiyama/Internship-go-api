package database

import (
	"context"

	"github.com/KentaroKajiyama/internship-go-api/domain/entity"
	"github.com/KentaroKajiyama/internship-go-api/infrastructure/database/model"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type SampleRepository struct {
	db *gorm.DB
}

func (s SampleRepository) Save(ctx context.Context, sample entity.Sample) (*entity.Sample, error) {
	//db.WithContext(ctx)で返ってくるのはdb, ctx内のリクエスト情報が込められる
	conn := s.db.WithContext(ctx)
	sampleModel := model.NewSampleFromEntity(sample)
	//Saveメソッドで primary key があるかないかでCREATEかUPDATEを実行する、その後エラーをハンドリングする。ここではエラーはうまくやってくれると信じる。
	if err := conn.Save(&sampleModel).Error; err != nil {
		return nil, err
	}
	// sampleModelをpointerに変換する
	return lo.ToPtr(sampleModel.ToEntity()), nil
}

func (s SampleRepository) FindById(ctx context.Context, id string) (*entity.Sample, error) {
	conn := s.db.WithContext(ctx)
	var sampleModel model.Sample
	if err := conn.Where("id = ?", id).First(&sampleModel).Error; err != nil {
		return nil, err
	}
	return lo.ToPtr(sampleModel.ToEntity()), nil
}

// func NewSampleRepository(db *gorm.DB) domain.ISampleRepository {
// 	return &SampleRepository{db: db}
// }
