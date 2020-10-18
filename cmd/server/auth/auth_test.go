package auth_test

// u need create toml file with test user from ur DB in this folder

import (
	. "server/auth"
	"testing"
)

var u User = User{
	Login:    "test",
	Username: "test",
	Password: "test",
}
var token interface{} = ""

func Test_token(t *testing.T) {
	token, key, err := CreatToken(u)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	t.Logf("token: %s, key: %d", token, key)

}

func Test_newSession(t *testing.T) {
	token, key, err := CreatToken(u)
	if err != nil {
		t.Errorf("error: %d", err)
		return
	}

	t.Logf("token created %d, key: %d", token, key)
	err = NewSession(u, token.(string), key)
	if err != nil {
		t.Errorf("error: %d", err)
	}

}

func Test_sessionValid(t *testing.T) {
	_, key, err := CreatToken(u)
	if err != nil {
		t.Errorf("error: %d", err)
		return
	}
	s := Session{
		User:  u.Login,
		Token: token.(string),
		Key:   key,
	}
	err = s.Valid()
	if err != nil {
		t.Logf("User: %s, key: %d", s.User, key)
		t.Errorf("error: %d", err)
	}
}
