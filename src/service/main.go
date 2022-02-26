package main

import (
	"log"
	"os"

	"github.com/hirac1220/go/go-monorepo/utils/file"
)

func main() {
	// ファイルサイズチェック
	f, _ := os.Open("./sample/image.jpg") // ファイルオープン
	defer func() {
		_ = f.Close()
	}()

	fi, _ := f.Stat()
	size, err := file.CheckFileSize(fi.Size()) // ファイルサイズチェック
	if err != nil {
		log.Println(err)
	}
	log.Printf("size: %d", size)

}
