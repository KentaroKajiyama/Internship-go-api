package repository

import (
	"context"

	userDomain "github.com/KentaroKajiyama/Internship-go-api/domain/user"
	"github.com/KentaroKajiyama/Internship-go-api/infrastructure/database/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) userDomain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Find(ctx context.Context, id string) (*userDomain.User, error) {
	conn := r.db.WithContext(ctx)
	var userModel model.User
	var userDomainPtr *userDomain.User
	var errDom error
	if err := conn.Where("id = ?", id).Find(&userModel).Error; err != nil {
		return nil, err
	}
	//infra層からdomain層へ
	userDomainPtr, errDom = userModel.ToDomainUser()
	if errDom != nil {
		return nil, errDom
	}
	//上手くいったら取得したuser(domain)を返す。
	return userDomainPtr, nil
}

func (r *userRepository) FindByUid(ctx context.Context, FirebaseUid string) (*userDomain.User, error) {
	conn := r.db.WithContext(ctx)
	var userModel model.User
	var userDomainPtr *userDomain.User
	var errDom error
	if err := conn.Where("firebase_uid = ?", FirebaseUid).Find(&userModel).Error; err != nil {
		return nil, err
	}
	//infra層からdomain層へ
	userDomainPtr, errDom = userModel.ToDomainUser()
	if errDom != nil {
		return nil, errDom
	}
	//上手くいったら取得したuser(domain)を返す。
	return userDomainPtr, nil
}

func (r *userRepository) Create(ctx context.Context, user *userDomain.User) (*userDomain.User, error) {
	conn := r.db.WithContext(ctx)
	//domain層からinfra層へ
	userModel := model.NewUserFromDomainUser(user)
	if err := conn.Create(&userModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったuser(domain)を返す。
	return userModel.ToDomainUser()
}

func (r *userRepository) Update(ctx context.Context, user *userDomain.User) (*userDomain.User, error) {
	conn := r.db.WithContext(ctx)
	//domain層からinfra層へ
	userModel := model.NewUserFromDomainUser(user)
	if err := conn.Model(&model.User{}).Where("id = ?", user.Id()).Updates(&userModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったuser(domain)を返す。
	return userModel.ToDomainUser()
}

func (r *userRepository) Delete(ctx context.Context, user *userDomain.User) (*userDomain.User, error) {
	conn := r.db.WithContext(ctx)
	//domain層からinfra層へ
	userModel := model.NewUserFromDomainUser(user)
	if err := conn.Where("id = ?", user.Id()).Delete(&userModel).Error; err != nil {
		return nil, err
	}
	//データベース処理に問題がなければそのまま受け取ったuser(domain)を返す。
	return user, nil
}
