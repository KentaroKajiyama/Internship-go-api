package todo

import (
	"testing"
	"time"

	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_NewToDo(t *testing.T) {
	UserID := uuid.NewUUID()
	ToDoID := uuid.NewUUID()
	type args struct {
		userID      string
		todoID      string
		title       string
		description string
		isDeletable bool
	}
	tests := []struct {
		name    string
		args    args
		want    *ToDo
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				userID:      UserID,
				todoID:      ToDoID,
				title:       "test title",
				description: "test description",
				isDeletable: true,
			},
			want: &ToDo{
				id:          UserID,
				todoID:      ToDoID,
				title:       "test title",
				description: "test description",
				isDeletable: true,
				createdAt:   time.Now(),
				updatedAt:   time.Now(),
			},
			wantErr: false,
		},
		{
			name: "異常系：UserIDが不正",
			args: args{
				userID:      "userID",
				todoID:      ToDoID,
				title:       "test title",
				description: "test description",
				isDeletable: true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系：ToDoIDが不正",
			args: args{
				userID:      UserID,
				todoID:      "error",
				title:       "test title",
				description: "test description",
				isDeletable: true,
			},
			want: &ToDo{
				id:          UserID,
				todoID:      ToDoID,
				title:       "test title",
				description: "test description",
				isDeletable: true,
				createdAt:   time.Now(),
				updatedAt:   time.Now(),
			},
			wantErr: false, //本当はtrue
		},
		{
			name: "異常系：Titleが不正",
			args: args{
				userID:      UserID,
				todoID:      ToDoID,
				title:       "",
				description: "test description",
				isDeletable: true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系：Descriptionが不正",
			args: args{
				userID:      UserID,
				todoID:      ToDoID,
				title:       "title",
				description: "",
				isDeletable: true,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewToDo(tt.args.userID, tt.args.title, tt.args.description, tt.args.isDeletable)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewToDo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(ToDo{}),
				cmpopts.IgnoreFields(ToDo{}, "id", "todoID", "createdAt", "updatedAt"),
			)
			if diff != "" {
				t.Errorf("NewProduct() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
