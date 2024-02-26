package decrypt

func Nimbus(str string) string{ // fn name first char as Capital to visible outside
								// if lowercase then in can only accessible inside the package
	encrypt := ""
	for _,c := range str{
		ascii := int(c)
		ascii-=3
		encrypt += string(rune(ascii))
	}

	return encrypt
}