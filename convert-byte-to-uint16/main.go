package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	data := []byte{0x01, 0x02} // Example byte data

	var value uint16
	buf := bytes.NewReader(data)
	err := binary.Read(buf, binary.LittleEndian, &value)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Value:", value)
	// Output: Value: 513
}
