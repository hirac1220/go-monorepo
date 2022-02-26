package file

import (
	"errors"
	"fmt"
)

const (
	MaxSize = 10 * 1024 * 1024 // 10MB
)

var (
	ErrFileSize = errors.New("file size error")
)

func CheckFileSize(size int64) (int64, error) {
	if size > MaxSize { // MaxSizeより大きい場合はエラー
		return size, fmt.Errorf("size: %v error: %w", size, ErrFileSize)
	}
	return size, nil
}
