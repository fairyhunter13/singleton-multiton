package singleton

import (
	"testing"
)

func BenchmarkSendEmail(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wrap.sendEmail("benchmark@example.com", "Let's benchmark this function!")
		}
	})
}
