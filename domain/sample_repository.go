package domain

import (
	"context"

	"github.com/KentaroKajiyama/Internship-go-api/domain/entity"
)

type ISampleRepository interface {
	FindById(ctx context.Context, id string) (*entity.Sample, error)
	Save(ctx context.Context, sample entity.Sample) (*entity.Sample, error)
}
