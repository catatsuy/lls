package cli

import (
	"bytes"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRun_Success(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cl := NewCLI(outStream, errStream)
	target := "testdata"
	status := cl.run(target, false, 0)
	if status != ExitCodeOK {
		t.Errorf("ExitStatus=%d, want %d", status, ExitCodeOK)
	}
	actualList := strings.Split(outStream.String(), "\n")
	// last element is unused
	actualList = actualList[:len(actualList)-1]
	expectedList := dirList(t, target)
	sort.Slice(actualList, func(i, j int) bool { return actualList[i] < actualList[j] })

	if diff := cmp.Diff(actualList, expectedList); diff != "" {
		t.Errorf("file list mismatch (-actual +expected):\n%s", diff)
	}
}

func dirList(t *testing.T, target string) []string {
	files, err := os.ReadDir(target)
	if err != nil {
		t.Error(err)
	}

	names := make([]string, 0, len(files))

	for _, file := range files {
		names = append(names, file.Name())
	}

	return names
}
