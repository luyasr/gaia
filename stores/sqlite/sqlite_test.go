package sqlite

import "testing"

func TestNewSqlite(t *testing.T) {
	c := &Config{}
	s, err := New(c)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(s)
}
