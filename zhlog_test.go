package zhlog

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func BenchmarkDebug(b *testing.B) {
	f, err := os.OpenFile("./my.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		b.Error(err)
	}
	defer f.Close()
	SetOutPut(f)
	SetLevel(LevelDebug)
	for i := 0; i < b.N; i++ {
		Debug("Goood:%d", i)
	}
	f.Close()

	if err := os.Remove("./my.log"); err != nil {
		b.Error(err)
	}
}

func TestDebug(t *testing.T) {
	var cache []byte
	buffer := bytes.NewBuffer(cache)
	SetOutPut(buffer)
	SetLevel(LevelDebug)
	now := time.Now().Format(time.RFC3339)
	now = strings.Replace(now, "T", " ", 1)
	spiStr := strings.Split(now, "+")
	now = spiStr[0]
	now = strings.Replace(now, "-", "/", 2)

	Debug("Debug out puting %s", "yes")
	str, err := buffer.ReadString('\n')
	if err != nil && err != io.EOF {
		t.Error(err)
	}
	want := "[D]" + now + " zhlog_test.go:41:" + " Debug out puting yes\n"
	if str != want {
		t.Errorf("get(len:%d) %s, want(len:%d) %s", len(str), str, len(want), want)
	}
}
