package main

import (
	"io"
	"os"
	"strings"
)

// ROT13 とはアルファベットの単純な入れ替えの暗号文のこと
// 26 文字を半分の13文字で区切って、対応する文字と入れ替える
// アルファベット以外の '!' などはそのまま
func rot13(b byte) byte {
	var base_byte_code byte

	switch {
	case b >= 'A' && b <= 'Z':
		base_byte_code = 'A'
	case b >= 'a' && b <= 'z':
		base_byte_code = 'a'
	default:
		return b
	}

	// 元の文字に半分(13 文字分)の値を足して
	// 26 文字で割った余りを求めて、アルファベットコードに戻す
	return (((b - base_byte_code) + 13) % 26) + base_byte_code
}

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	size, err := r.r.Read(b)

	for i, v := range b {
		b[i] = rot13(v)
	}

	return size, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
