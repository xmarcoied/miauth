package auth

import (
	"context"
	"reflect"
	"testing"

	"github.com/xmarcoied/miauth/models"
	"github.com/xmarcoied/miauth/services/storage"
)

func TestService_CreateUser(t *testing.T) {
	type fields struct {
		storage storage.UsersInterface
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "successful create user",
			fields: fields{
				storage: &storage.MockedUser{},
			},
			args: args{
				ctx:      context.TODO(),
				username: "username",
				password: "password",
			},
			want: models.User{
				Username: "username",
				Password: "password",
			},
			wantErr: false,
		},
		{
			name: "already exist user",
			fields: fields{
				storage: storage.NewMockedUser([]models.User{
					models.User{
						Username: "username",
						Password: "password",
					},
				}),
			},
			args: args{
				ctx:      context.TODO(),
				username: "username",
				password: "password",
			},
			want:    models.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				storage: tt.fields.storage,
			}
			got, err := s.CreateUser(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
