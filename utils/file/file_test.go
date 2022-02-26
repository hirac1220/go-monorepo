package file

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// テストの開始・終了
func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

// テスト初期設定
func setup() {
	log.Println("test start!")
}

// テスト終了処理
func teardown() {
	log.Println("test end!")
}

func TestCheckFileSize(t *testing.T) {
	// arrange
	f, _ := os.Open("../../sample/image.jpg")
	defer func() {
		_ = f.Close()
	}()

	// action
	fi, _ := f.Stat()
	_, actual := CheckFileSize(fi.Size())

	// assert
	assert.ErrorIs(t, nil, actual)
}
