package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
)

func main() {
	var a int64 = 4531

	buf := Int64ToBytes(a)

	//проход по массиву байт и замена значений на 1 или 0
	for k := range buf {
		buf[k] = byte(rand.Intn(2))
	}

	fmt.Println(buf)
	fmt.Println(BytesToInt64(buf))
}

// преобразование int64 в массив байт
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
