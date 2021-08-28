package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	target := "."
	if len(os.Args) > 1 {
		target = os.Args[1]
	}

	f, err := os.Open(target)
	if err != nil {
		log.Fatal(err)
	}

	finfo, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// about actual size: 20 + filename
	// ls -dl
	buf := make([]byte, finfo.Size())
	n, err := syscall.Getdents(int(f.Fd()), buf)
	if err != nil {
		log.Fatal(err)
	}

	for bufp := 0; bufp < n; {
		dirent := (*syscall.Dirent)(unsafe.Pointer(&buf[bufp]))
		bufp += int(dirent.Reclen)

		// deleted file
		if dirent.Ino == 0 {
			continue
		}

		bb := (*[256]byte)(unsafe.Pointer(&dirent.Name[0]))
		name := string(bb[0:blen(*bb)])
		fmt.Println(name)
	}
}

func blen(b [256]byte) int {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return i
		}
	}
	return len(b)
}
