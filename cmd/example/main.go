package main 

import (
	"fmt"
	"github.com/marcelluseasley/go-pii-masker/pkg/masker"
)

func main() {
	fmt.Println(masker.MaskEmail("marcellus.easley@gmail.com"))
	fmt.Println(masker.MaskPhone("404-510-5373"))
	fmt.Println(masker.MaskSSN("123-45-6789"))
}