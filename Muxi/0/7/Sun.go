package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	objectfile, err := os.OpenFile("muxi-backend/paper/Academician Sun's papers.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0123)
	if err != nil {
		fmt.Println(err)
	}
	defer objectfile.Close()
	objectfile.WriteString("Here is muxi-backend/paper/Academician Sun's papers.txt ")

}
