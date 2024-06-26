package todo

import (
	"context"

	todoDomain "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
)

// コード内では実装済みのtodoRepositoryが送られてくるが、UseCaseのプログラム上ではDomain層のinterface：Repositoryとして扱われる。ここで渡される時点ですでに構造体todoRepositoryoは具体的なメソッドを実装している。そのためもちろんinterface:Repositoryでもある。struct:Repositoryでもある。ってことで間違いないと思う。interfaceって不思議だし、柔軟性がある。
type CreateTodoUseCase struct {
	todoRepository todoDomain.TodoRepository
}

func NewCreateTodoUseCase(todoRepository todoDomain.TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{todoRepository: todoRepository}
}

// todo項目新規作成
type CreateTodoUseCaseInputDto struct {
	Id          string
	Title       string
	Description string
	IsDeletable bool
}

// 新規項目を作成してリポジトリに登録する、userはどうする？ todoIDを新規に作るドメインサービスが必要？
func (uc *CreateTodoUseCase) Create(ctx context.Context, dto CreateTodoUseCaseInputDto) (*todoDomain.Todo, error) {
	todo, err := todoDomain.NewTodoWithoutTodoIdAndTime(dto.Id, dto.Title, dto.Description, dto.IsDeletable)
	if err != nil {
		return nil, err
	}
	return uc.todoRepository.Create(ctx, todo)
}
