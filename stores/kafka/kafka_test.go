package kafka

import "testing"

func TestNew(t *testing.T) {
	c := Config{
		Broker:    "localhost:9092",
		Topic:     "test",
		Partition: 0,
		Timeout:   10,
	}
	k, err := New(&c)

	if err != nil {
		t.Fatal(err)
	}
	t.Log(k)
}
