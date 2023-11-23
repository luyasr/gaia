package reflection

import (
	"fmt"
	"testing"
)

func TestSetDefaultTag(t *testing.T) {
	type Person struct {
		Name  string `json:"name" default:"alex"`
		Age   int    `json:"age" default:"18"`
		Hobby struct {
			Swim      bool `json:"swim" default:"true"`
			PlayGames bool `json:"playGames"`
		}
	}

	p := Person{Name: "tom", Age: 20}
	SetUp(&p)
	fmt.Println(p)
}
