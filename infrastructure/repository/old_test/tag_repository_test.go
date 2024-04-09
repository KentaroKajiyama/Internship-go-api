package old

import (
	"context"
	"testing"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure"
	tagRepo "github.com/KentaroKajiyama/Internship-go-api/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestDeleteTag(t *testing.T) {
	db := infrastructure.NewGormPostgres()
	defer func() {
		d, _ := db.DB()
		d.Close()
	}()

	repo := tagRepo.NewTagRepository(db)
	tag, err := tagDomain.NewTagWithoutTime("491e73f0-e93d-2111-33a0-ec82b4b7647c", 1, "test")
	if err != nil {
		t.Fatalf("Failed to create user domain instance: %v", err)
	}
	// テスト実行
	_, err = repo.Delete(context.Background(), tag)

	// エラーがないことを確認
	assert.NoError(t, err)
}
