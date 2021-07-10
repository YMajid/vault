package main

import (
	"fmt"

	"github.com/YMajid/vault"
)

func main() {
	v := vault.File("my-fake-key", ".secrets")
	_ = v.Set("demo_key_1", "123some crazy value")
	_ = v.Set("demo_key_2", "some 123crazy value")
	_ = v.Set("demo_key_3", "some crazy 123value")

	plain, _ := v.Get("demo_key_1")
	fmt.Println("Plain: ", plain)
	plain, _ = v.Get("demo_key_2")
	fmt.Println("Plain: ", plain)
	plain, _ = v.Get("demo_key_3")
	fmt.Println("Plain: ", plain)
}
