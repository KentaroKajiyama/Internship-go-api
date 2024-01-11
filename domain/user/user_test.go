package user

import (
	"testing"
	"time"

	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func Test_NewProduct(t *testing.T) {
	UserID := uuid.NewUUID()
	type args struct {
		id    string
		name  string
		email string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id:    UserID,
				name:  "test",
				email: "test@gmail.com",
			},
			want: &User{
				id:        UserID,
				name:      "test",
				email:     "test@gmail.com",
				createdAt: time.Now(),
				updatedAt: time.Now(),
			},
			wantErr: false,
		},
		{
			name: "異常系：Nameが不正",
			args: args{
				id:    UserID,
				name:  "",
				email: "test@gmail.com",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "異常系：Emailが不正",
			args: args{
				id:    UserID,
				name:  "test",
				email: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, want %v", err, tt.wantErr)
				return
			}
			diff := cmp.Diff(
				got, tt.want,
				cmp.AllowUnexported(User{}),
				cmpopts.IgnoreFields(User{}, "id", "createdAt", "updatedAt"),
			)
			if diff != "" {
				t.Errorf("NewUser() = %v, want %v. error is %s", got, tt.want, err)
			}
		})
	}
}
