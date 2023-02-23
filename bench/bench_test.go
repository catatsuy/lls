package bench

import (
	"encoding/base64"
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

var rd = rand.New(rand.NewSource(time.Now().UnixNano()))

func BenchmarkUseStringCast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bb, _ := genByteArray256()
		_ = string(bb[0:256])
	}
}

func BenchmarkUseUnsafeString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bb, _ := genByteArray256()
		_ = unsafe.String(&bb[0:256][0], 256)
	}
}

func genByteArray256() (bb [256]byte, err error) {
	b := make([]byte, 191)
	_, err = rd.Read(b)
	if err != nil {
		return bb, err
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(dst, b)

	return [256]byte(dst), nil
}
