package backend

import (
	"github.com/gin-gonic/gin"
	"sciffr/backend/cryptservice"
	"testing"
)

func TestBackend_encrypt(t *testing.T) {
	type fields struct {
		crypt cryptservice.CryptService
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Backend{
				crypt: tt.fields.crypt,
			}
			b.encrypt(tt.args.c)
		})
	}
}

func TestBackend_decrypt(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Backend{
				crypt: cryptservice.New(),
			}
			b.decrypt(tt.args.c)
		})
	}
}
