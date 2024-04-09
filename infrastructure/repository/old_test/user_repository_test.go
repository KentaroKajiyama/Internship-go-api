package old

import (
	"context"
	"testing"
	"time"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure"
	userRepo "github.com/KentaroKajiyama/Internship-go-api/infrastructure/repository"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db := infrastructure.NewGormPostgres()
	defer func() {
		d, _ := db.DB()
		d.Close()
	}()

	repo := userRepo.NewUserRepository(db)
	user, err := userDomain.NewUser("a22487c1-52f7-a6b8-dcdf-184803461e3e", "test2", "test2@example.com", time.Now().UTC(), time.Now().UTC())
	if err != nil {
		t.Fatalf("Failed to create user domain instance: %v", err)
	}
	// テスト実行
	_, err = repo.Create(context.Background(), user)

	// エラーがないことを確認
	assert.NoError(t, err)
}
