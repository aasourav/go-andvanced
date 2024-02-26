package main

import (
	"fmt"

	//======== module path ====/packagename
	"github.com/aasourav/crypto/decrypt"
	"github.com/aasourav/crypto/encrypt"
)

func main(){
	fmt.Println(encrypt.Nimbus("hello"))
	fmt.Println(decrypt.Nimbus(encrypt.Nimbus("hello")))
}