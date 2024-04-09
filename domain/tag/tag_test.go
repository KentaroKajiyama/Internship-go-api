package tag

// import (
// 	"testing"

// 	"github.com/KentaroKajiyama/Internship-go-api/pkg/uuid"
// 	"github.com/google/go-cmp/cmp"
// 	"github.com/google/go-cmp/cmp/cmpopts"
// )

// func Test_NewTag(t *testing.T) {
// 	UserID := uuid.NewUUID()
// 	TagID := 10
// 	type args struct {
// 		userID string
// 		tagID  int
// 		name   string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *Tag
// 		wantErr bool
// 	}{
// 		{
// 			name: "正常系",
// 			args: args{
// 				userID: UserID,
// 				tagID:  TagID,
// 				name:   "test name",
// 			},
// 			want: &Tag{
// 				id:    UserID,
// 				tagID: TagID,
// 				name:  "test name",
// 			},
// 			wantErr: false,
// 		},
// 		{
// 			name: "異常系：UserIDが不正",
// 			args: args{
// 				userID: "userID",
// 				tagID:  TagID,
// 				name:   "test name",
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name: "異常系：TagIDが不正", //NewTagで元々のTagIDを無視して新しくIDを作る仕様上ここでは必ずエラーが出る。
// 			args: args{
// 				userID: UserID,
// 				tagID:  11,
// 				name:   "test name",
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 		{
// 			name: "異常系：Nameが不正",
// 			args: args{
// 				userID: UserID,
// 				tagID:  TagID,
// 				name:   "",
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := NewTag(tt.args.userID, tt.args.name)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewTag() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			diff := cmp.Diff(
// 				got, tt.want,
// 				cmp.AllowUnexported(Tag{}),
// 				cmpopts.IgnoreFields(Tag{}, "id", "createdAt", "updatedAt"),
// 			)
// 			if diff != "" {
// 				t.Errorf("NewProduct() = %v, want %v. error is %s", got, tt.want, err)
// 			}
// 		})
// 	}
// }
