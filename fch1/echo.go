package fch1

import (
	"fmt"
	"os"
)

func Echo() {
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}