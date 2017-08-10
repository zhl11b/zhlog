package zhlog

import (
	"os"
	"testing"
)

func BenchmarkDebug(b *testing.B) {
	f, err := os.OpenFile("./my.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		b.Error(err)
	}
	defer f.Close()
	SetOutPut(f)
	SetLevel(LevelError)
	for i := 0; i < b.N; i++ {
		Fatal("Goood:%d", i)
	}
	f.Close()
	if err := os.Remove("./my.log"); err != nil {
		b.Error(err)
	}
}
