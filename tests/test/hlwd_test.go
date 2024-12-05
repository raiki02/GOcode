package test

import (
	"t/helloworld"
	"testing"
)

func TestHlwd(t *testing.T) {
	helloworld.HelloWorld()
}

func BenchMark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		helloworld.HelloWorld()
	}

}
