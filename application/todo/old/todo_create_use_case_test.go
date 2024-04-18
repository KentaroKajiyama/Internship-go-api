package todo

// import (
// 	"context"
// 	"testing"

// 	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
// 	"go.uber.org/mock/gomock"

// 	todoDomoin "github.com/KentaroKajiyama/Internship-go-api/domain/todo"
// )

// func TestCreateToDoUseCase_Run(t *testing.T) {
// 	//usecase準備
// 	ctrl := gomock.NewController(t)
// 	mockToDoRepo := todoDomoin.NewMockToDoRepository(ctrl)
// 	uc := NewCreateToDoUseCase(mockToDoRepo)
// 	UserID := uuid.NewUUID()

// 	//各種テストデータ準備
// 	dto := CreateToDoUseCaseInputDto{
// 		ID:          UserID,
// 		Title:       "Test ToDo 1",
// 		Description: "Create Test ToDo 1",
// 		IsDeletable: false,
// 	}
// 	todo, _ := todoDomoin.NewToDo(dto.ID, dto.Title, dto.Description, dto.IsDeletable)

// 	tests := []struct {
// 		name     string
// 		dto      CreateToDoUseCaseInputDto
// 		mockFunc func()
// 		wantErr  bool
// 	}{
// 		{
// 			name: "正常系",
// 			dto:  dto,
// 			mockFunc: func() {
// 				gomock.InOrder(
// 					// mockされた関数の期待値を検証
// 					mockToDoRepo.EXPECT().Create(gomock.Any(), todo).Return(nil),
// 				)
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt := tt
// 			t.Parallel()
// 			tt.mockFunc()
// 			_, err := uc.Create(context.Background(), tt.dto)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("CreateToDoUseCase.Create() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
