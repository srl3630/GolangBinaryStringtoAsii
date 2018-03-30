package main

import (
	"os"
	"image/png"
	"log"
	"strings"
	"fmt"
	"math"
	//"container/list"
)




func imagetostring(){
	bin := decryptimage("out.png")
	fmt.Println(bintoasc(bin))


}

func decryptimage(pngtoread string)(binStr string){
	binStr = ""
	myimage, _ := os.Open(pngtoread)
	img, err := png.Decode(myimage)
	if err != nil {
		log.Fatal(err)
	}
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++{
			_,_,_,a := img.At(x, y).RGBA()
			if a == 65278 {
				binStr = joinstrings(binStr, "0")
			} else if a == 65535 {
				binStr = joinstrings(binStr,"1")
			} else {
				return
			}

		}
	}
	myimage.Close()
	return
}

func bintoasc(bin string)(ascii string) {
	asciimap := map[string]string{
		"00100000": " ",
		"00100001": "!",
		"00100010": "\"",
		"00100011": "#",
		"00100100": "$",
		"00100101": "%",
		"00100110": "&",
		"00100111": "'",
		"00101000": "(",
		"00101001": ")",
		"00101010": "*",
		"00101011": "+",
		"00101100": ",",
		"00101101": "-",
		"00101110": ".",
		"00101111": "/",
		"00110000": "0",
		"00110001": "1",
		"00110010": "2",
		"00110011": "3",
		"00110100": "4",
		"00110101": "5",
		"00110110": "6",
		"00110111": "7",
		"00111000": "8",
		"00111001": "9",
		"00111010": ":",
		"00111011": ";",
		"00111100": "<",
		"00111101": "=",
		"00111110": ">",
		"00111111": "?",
		"01000000": "@",
		"01000001": "A",
		"01000010": "B",
		"01000011": "C",
		"01000100": "D",
		"01000101": "E",
		"01000110": "F",
		"01000111": "G",
		"01001000": "H",
		"01001001": "I",
		"01001010": "J",
		"01001011": "K",
		"01001100": "L",
		"01001101": "M",
		"01001110": "N",
		"01001111": "O",
		"01010000": "P",
		"01010001": "Q",
		"01010010": "R",
		"01010011": "S",
		"01010100": "T",
		"01010101": "U",
		"01010110": "V",
		"01010111": "W",
		"01011000": "X",
		"01011001": "Y",
		"01011010": "Z",
		"01011011": "[",
		"01011100": "\\",
		"01011101": "]",
		"01011110": "^",
		"01011111": "_",
		"01100000": "`",
		"01100001": "a",
		"01100010": "b",
		"01100011": "c",
		"01100100": "d",
		"01100101": "e",
		"01100110": "f",
		"01100111": "g",
		"01101000": "h",
		"01101001": "i",
		"01101010": "j",
		"01101011": "k",
		"01101100": "l",
		"01101101": "m",
		"01101110": "n",
		"01101111": "o",
		"01110000": "p",
		"01110001": "q",
		"01110010": "r",
		"01110011": "s",
		"01110100": "t",
		"01110101": "u",
		"01110110": "v",
		"01110111": "w",
		"01111000": "x",
		"01111001": "y",
		"01111010": "z",
		"01111011": "{",
		"01111100": "|",
		"01111101": "}",
		"01111110": "~",
	}

	ascii = ""
	tempstring := ""
	bytebin := []byte(bin)
	for index, element := range bytebin{
		if math.Mod(float64(index), 8) == 7 {
			tempstring = joinstrings(tempstring, string(element))
			ascii = joinstrings(ascii,asciimap[tempstring])
			tempstring = ""
		} else {
			tempstring = joinstrings(tempstring, string(element))
		}
	}
	return


}



func joinstrings(string1, string2 string) (mashstring string){
	var strs []string
	strs = append(strs, string1)
	strs = append(strs, string2)
	mashstring = strings.Join(strs, "")
	return
}
