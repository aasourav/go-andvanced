package main

import (
	"fmt"

	"github.com/aasourav/crypto/decrypt"
	"github.com/aasourav/crypto/encrypt" // this is local module, so we have to use replace director
	"github.com/pborman/uuid"            //go mod tidy
)

func main(){
 uuid := uuid.NewRandom();
 fmt.Println(uuid)
 fmt.Println(encrypt.Nimbus("hello"))
 fmt.Println(decrypt.Nimbus(encrypt.Nimbus("hello")))
}


