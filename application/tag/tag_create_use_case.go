package tag

import (
	"context"

	tagDomain "github.com/KentaroKajiyama/Internship-go-api/domain/tag"
)

type CreateTagUseCase struct {
	tagRepository tagDomain.TagRepository
}

func NewCreateTagUseCase(tagRepository tagDomain.TagRepository) *CreateTagUseCase {
	return &CreateTagUseCase{tagRepository: tagRepository}
}

// tag項目新規作成
type CreateTagUseCaseInputDto struct {
	Id   string
	Name string
}

// 新規項目を作成してリポジトリに登録する。作成して永続化って感じだからIDの生成はドメイン層でもいいかも
func (uc *CreateTagUseCase) Create(ctx context.Context, dto CreateTagUseCaseInputDto) (*tagDomain.Tag, error) {
	// 自動インクリメントの時にここのTagIDをどうするか？→ゼロ値の0に設定
	user, err := tagDomain.NewTagFirst(dto.Id, 0, dto.Name)
	if err != nil {
		return nil, err
	}
	return uc.tagRepository.Create(ctx, user)
}
