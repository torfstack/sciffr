package cryptservice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cryptService_Decrypt(t *testing.T) {
	type args struct {
		encrypted Encrypted
	}
	tests := []struct {
		name          string
		argsCreation  func(*cryptService) args
		wantAssertion func(*testing.T, []byte)
	}{
		{
			name: "encryption can be decrypted",
			argsCreation: func(service *cryptService) args {
				encrypted := service.Encrypt([]byte("hello world!"))
				return args{
					encrypted: Encrypted{
						Ciphertext: encrypted.Ciphertext,
						Key:        encrypted.Key,
					},
				}
			},
			wantAssertion: func(t *testing.T, bytes []byte) {
				assert.Equal(t, "hello world!", string(bytes))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cryptService{
				masterCipher: initCipher(generateKey()),
			}
			got := c.Decrypt(tt.argsCreation(c).encrypted)
			tt.wantAssertion(t, got)
		})
	}
}

func Test_cryptService_Encrypt(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name          string
		args          args
		wantAssertion func(*testing.T, Encrypted)
	}{
		{
			name: "encrypt empty string has correct length",
			args: args{
				bytes: []byte{},
			},
			wantAssertion: func(t *testing.T, encrypted Encrypted) {
				assert.Equal(t, 28, len(encrypted.Ciphertext))
			},
		},
		{
			name: "encrypt 'hello world!' string has correct length",
			args: args{
				bytes: []byte("hello world!"),
			},
			wantAssertion: func(t *testing.T, encrypted Encrypted) {
				assert.Equal(t, 40, len(encrypted.Ciphertext))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cryptService{
				masterCipher: initCipher(generateKey()),
			}
			got := c.Encrypt(tt.args.bytes)
			tt.wantAssertion(t, got)
		})
	}
}
