package main

import (

	"fmt"
	"github.com/Nathan-Groves/img_mod/Colors"
	"github.com/Nathan-Groves/img_mod/Getpic"
	"github.com/Nathan-Groves/img_mod/Grayscale"
	"github.com/Nathan-Groves/img_mod/Text"

)

func main(){


	// run functions and print whitespace
	
	Getpic.GetImage()

	fmt.Print('\n')
	fmt.Print('\n')

	Colors.PrintColors()

	fmt.Print('\n')
	fmt.Print('\n')

	Grayscale.MakeGrayScale()

	fmt.Print('\n')
	fmt.Print('\n')

	Text.PrintImageText()

}
