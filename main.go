package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	makeFile("1.txt")

	if err := os.Mkdir("/tmp/test-dir", 0666); err != nil {
		fmt.Println(err)
	}
	makeFile("/tmp/test-dir/2.txt")
}

func makeFile(name string) {
	fp, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	fp.WriteString("hello")

	// 直近で書き込んだ内容をReadするにはSeekでファイルの先頭に戻る必要がある
	fp.Seek(0, 0)

	b := make([]byte, 256)
	for {
		_, err := fp.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
				return
			}
			break
		}
		fmt.Println(string(b))
	}
}
