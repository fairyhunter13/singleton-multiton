package multiton

import "testing"

func BenchmarkSendEmail(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sendEmail("benchmark@example.com", "Let's benchmark this function!")
		}
	})
}
