package main

import (
	"exampleMask/mask"
	"exampleMask/search"
	"fmt"
)

const page = "Here's my spammy page: http://hehefouls.netHAHAHA see you."
const key = "http://"
const m, s = '*', ' '

func main() {
	fmt.Println(page)
	p := []byte(page)
	k := []byte(key)
	n := search.Search(p, k)
	fmt.Println(mask.Mask(n, p, s, m))
}
