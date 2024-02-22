package gopt

import (
	"testing"
)

func Test(t *testing.T) {
	ga := NewGoogleAuthenticator()
	secret := ga.GenerateSecret()
	code, _ := ga.GenerateCode(secret)
	qr := ga.GenerateQRUrl("dd", "test@qq.com", secret)

	t.Log(secret, code, qr)
}
