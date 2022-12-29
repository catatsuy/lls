package cli

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"syscall"
	"unsafe"
)

var (
	Version string
)

const (
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
	ExitCodeFail           = 1

	// 5MB
	DefaultBufSize = 5 * 1024 * 1024
)

type CLI struct {
	outStream, errStream io.Writer

	appVersion string
}

func NewCLI(outStream, errStream io.Writer) *CLI {
	return &CLI{appVersion: version(), outStream: outStream, errStream: errStream}
}

func version() string {
	if Version != "" {
		return Version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "(devel)"
	}
	return info.Main.Version
}

func (c *CLI) Run(args []string) int {
	var (
		version bool
		debug   bool
		bufSize int
	)

	flags := flag.NewFlagSet("lls", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	flags.IntVar(&bufSize, "buf-size", DefaultBufSize, "specify buf size")

	flags.BoolVar(&version, "version", false, "print version information and quit")
	flags.BoolVar(&debug, "debug", false, "debug mode")

	err := flags.Parse(args[1:])
	if err != nil {
		return ExitCodeParseFlagError
	}

	if version {
		fmt.Fprintf(c.errStream, "lls version %s; %s\n", c.appVersion, runtime.Version())
		return ExitCodeOK
	}

	argv := flags.Args()
	target := "."
	if len(argv) == 1 {
		target = argv[0]
	} else if len(argv) > 1 {
		target = argv[0]
		err = flags.Parse(argv[1:])
		if err != nil {
			return ExitCodeParseFlagError
		}

		argv = flags.Args()
		if len(argv) > 0 {
			return ExitCodeParseFlagError
		}
	}

	return c.run(target, debug, bufSize)
}

func (c *CLI) run(target string, debug bool, bufSize int) int {
	f, err := os.Open(target)
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeFail
	}
	defer f.Close()

	finfo, err := f.Stat()
	if err != nil {
		fmt.Fprintln(c.errStream, err)
		return ExitCodeFail
	}

	if !finfo.IsDir() {
		fmt.Fprintln(c.errStream, "not a directory")
		return ExitCodeFail
	}

	buf := make([]byte, bufSize)

	for {
		n, err := syscall.Getdents(int(f.Fd()), buf)
		if err != nil {
			fmt.Fprintln(c.errStream, err)
			return ExitCodeFail
		}

		if n == 0 {
			break
		}

		if debug {
			fmt.Fprintf(c.errStream, "bufSize: %d; getdents ret: %d\n", bufSize, n)
		}

		for bufp := 0; bufp < n; {
			dirent := (*syscall.Dirent)(unsafe.Pointer(&buf[bufp]))
			bufp += int(dirent.Reclen)

			// deleted file
			if dirent.Ino == 0 {
				continue
			}

			bb := (*[256]byte)(unsafe.Pointer(&dirent.Name))
			name := string(bb[0:blen(*bb)])

			if name == "." || name == ".." {
				// ignore
				continue
			}
			fmt.Fprintln(c.outStream, name)
		}
	}

	return ExitCodeOK
}

func blen(b [256]byte) int {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			return i
		}
	}
	return len(b)
}
