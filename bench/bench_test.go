package bench

import (
	"encoding/base64"
	"math/rand"
	"testing"
	"time"
	"unsafe"
)

func BenchmarkUseStringCast(b *testing.B) {
	src := generatePayload(b)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = string(src[:])
	}
}

func BenchmarkUseUnsafeString(b *testing.B) {
	src := generatePayload(b)
	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		_ = unsafe.String(&src[0], len(src))
	}
}

func generatePayload(b *testing.B) (dst [256]byte) {
	b.Helper()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	var raw [191]byte
	if _, err := rng.Read(raw[:]); err != nil {
		b.Fatalf("rng.Read: %v", err)
	}
	base64.StdEncoding.Encode(dst[:], raw[:])
	return
}
