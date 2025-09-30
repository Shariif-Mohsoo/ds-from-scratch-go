package readerwriter

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func normalRead() {
	reader := strings.NewReader("SAMPLE")
	var newString strings.Builder
	buffer := make([]byte, 4)
	for {
		numBytes, err := reader.Read(buffer)
		fmt.Println(numBytes)
		chunk := buffer[:numBytes]
		fmt.Println(chunk)
		newString.Write(chunk)
		fmt.Println(newString)
		fmt.Printf("Read %v bytes: %c\n", numBytes, chunk)
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("\n%v\n", newString.String())
}

func bufioRead() {
	source := strings.NewReader("Sample")
	buffered := bufio.NewReader(source)

	newString, err := buffered.ReadString('\n')
	if err == io.EOF {
		fmt.Println(newString)
	} else {
		fmt.Println("something went wrong")
	}

}

func Reader_Writer() {
	normalRead()
	bufioRead()
}
