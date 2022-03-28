package identity

import (
	"reflect"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestParseToken(t *testing.T) {
	type args struct {
		tokenStr string
		key      []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				tokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjQ0MzM5NzksImlhdCI6MTYyMzgyOTE3OSwiVXNlcklEIjoxLCJEYXRhIjpudWxsfQ.ckfs-0N7RxT1Nc1xC9it9ykOjmVO8l-D8RFnDgWSUOc",
				key:      []byte("123"),
			},
			want: &Claims{
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: 1624433979,
					IssuedAt:  1623829179,
				},
				UserID: 1,
				Data:   nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseToken(tt.args.tokenStr, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRefresh(t *testing.T) {
	type args struct {
		tokenStr string
		key      []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				tokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjQ0MzM5NzksImlhdCI6MTYyMzgyOTE3OSwiVXNlcklEIjoxLCJEYXRhIjpudWxsfQ.ckfs-0N7RxT1Nc1xC9it9ykOjmVO8l-D8RFnDgWSUOc",
				key:      []byte("123"),
			},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjQ0MzYwMTUsImlhdCI6MTYyMzgyOTE3OSwiVXNlcklEIjoxLCJEYXRhIjpudWxsfQ.02c4btwBaIn-4dxNIkImG22lCilhrX8ZPtkc6o4T9uc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Refresh(tt.args.tokenStr, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Refresh() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Refresh() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyToken(t *testing.T) {
	type args struct {
		tokenStr string
		key      []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "yes",
			args: args{
				tokenStr: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MjQ0MzM5NzksImlhdCI6MTYyMzgyOTE3OSwiVXNlcklEIjoxLCJEYXRhIjpudWxsfQ.ckfs-0N7RxT1Nc1xC9it9ykOjmVO8l-D8RFnDgWSUOc",
				key:      []byte("123"),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidToken(tt.args.tokenStr, tt.args.key); got != tt.want {
				t.Errorf("VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
