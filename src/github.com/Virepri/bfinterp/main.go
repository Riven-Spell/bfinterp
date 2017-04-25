package main

import (
	"os"
	"fmt"
)

func main(){
	if len(os.Args) == 1 {
		fmt.Println("No file supplied")
		return
	}
	File := os.Args[1]
	fdo := ""
	fdid := []byte{}
	if fd, err := os.Open(File); err == nil {
		fdi,_ := fd.Stat()
		fdid = make([]byte,fdi.Size())
		fd.Read(fdid)
		fdo = string(fdid)

		prgmem := make([]byte,256) //256 bytes total in the memory of bf
		memptr := 0

		//fill brainfuck mem
		for k,_ := range prgmem {
			prgmem[k] = uint8(0)
		}


		for prgptr := 0; prgptr != len(fdo) || prgptr < 0; prgptr++ {
			switch string(fdo[prgptr]) {
			case "+":
				prgmem[memptr]++
			case "-":
				prgmem[memptr]--
			case "<":
				memptr--
				if memptr < 0 {
					memptr = 0
				}
			case ">":
				memptr++
				if memptr < len(prgmem) {
					prgmem = append(prgmem, uint8(0))
				}
			case "[":
				if prgmem[memptr] == 0 {
					for depth := 1; depth > 0; {
						prgptr++
						srcCharacter := fdo[prgptr]
						if srcCharacter == '[' {
							depth++
						} else if srcCharacter == ']' {
							depth--
						}
					}
				}
			case "]":
				for depth := 1; depth > 0; {
					prgptr--
					srcCharacter := fdo[prgptr]
					if srcCharacter == '[' {
						depth--
					} else if srcCharacter == ']' {
						depth++
					}
				}
				prgptr--
			case ".":
				fmt.Print(string(prgmem[memptr]))
			case ",":
				fmt.Scan(&prgmem[memptr])
			}
		}
	} else {
		fmt.Print("Error: brainfuck file not found")
	}
}