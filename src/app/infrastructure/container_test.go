package infrastructure

import (
	"testing"
)


func Benchmark_Execute(b *testing.B) {
	type args struct {
		cmd string
	}
	benchmarks := map[string]struct {
		args args
	}{
		"whoami": {
			args: args {
				cmd: "whoami",
			},
		},
		"ls": {
			args: args {
				cmd: "ls",
			},
		},
		"find / -name passwd": {
			args: args {
				cmd: "find / -name passwd",
			},
		},
	}
	sut := NewContainerHandler()
	for bName, bm := range benchmarks {
		b.Run(bName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sut.Execute(bm.args.cmd, "5c5b87abfc21")
			}
		})
	}
}