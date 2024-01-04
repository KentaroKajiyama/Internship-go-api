package user

/*
import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
)
func Test_NewProduct(t *testing.T) {
	id := uuid.New()
	type args struct {
		id			uuid.UUID
		name        string
		email	    string
		createdAt   time.Time
		updatedAt   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *Product
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id: id,
				name: "test",
				email: "test@gmail.com",
				createdAT: time.Now(),
				updatedAt: time.Now(),
			},
			want: &Product{
				id: id,
				name: "test",
				email: "test@gmail.com",
				createdAT: time.Now(),
				updatedAt: time.Now(),
			},
			wantErr: false,
		},
	}
// その他テストケース
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
					got, err := NewProduct(tt.args.ownerID, tt.args.name, tt.args.descript
ion, tt.args.price, tt.args.stock)
					if (err != nil) != tt.wantErr {
r)
nt, err)
} }
	t.Errorf("NewProduct() error = %v, wantErr %v", err, tt.wantEr
return
}
diff := cmp.Diff(
	got, tt.want,
	cmp.AllowUnexported(Product{}),
	cmpopts.IgnoreFields(Product{}, "id"),
)
if diff != "" {
} })
*/
