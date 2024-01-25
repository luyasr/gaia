package mongo

import "testing"

func TestNewMongo(t *testing.T) {
	c := Config{
		Host:     "localhost",
		Port:     27017,
		Username: "root",
		Password: "12345678",
	}
	m, err := New(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)
}
