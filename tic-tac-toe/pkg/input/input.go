package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputReader struct {
}

func NewInputReader() *InputReader {

	return &InputReader{}
}
func (r *InputReader) GetInput() (x int, y int, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter X: ")
	xString, _ := reader.ReadString('\n')
	xString = strings.TrimSpace(xString)
	x, err = strconv.Atoi(xString)
	if err != nil {
		return 0, 0, err
	}
	fmt.Print("Enter Y: ")
	yString, _ := reader.ReadString('\n')
	yString = strings.TrimSpace(yString)
	y, err = strconv.Atoi(yString)
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}
