package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/qpliu/qrencode-go/qrencode"
)

func main() {
	var d string

	if len(os.Args) > 1 {
		d = strings.Join(os.Args[1:], " ")
	} else {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			d += sc.Text()
		}
	}

	qr, err := qrencode.Encode(d, qrencode.ECLevelL)
	if err != nil {
		log.Fatal(err)
	}

	black := ansi.ColorCode("default:black")
	white := ansi.ColorCode("default:white")
	reset := ansi.ColorCode("reset")

	xmax := qr.Width()
	ymax := qr.Height()
	for x := 0; x < xmax; x++ {
		ln := ""

		for y := 0; y < ymax; y++ {
			if qr.Get(x, y) {
				ln += black
			} else {
				ln += white
			}

			ln += "  "
		}

		fmt.Println(ln)
	}

	fmt.Println(reset)
}
