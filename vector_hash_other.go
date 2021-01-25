// +build !amd64

package xxh3_seed

const (
	avx2 = false
	sse2 = false
)

func hashVector(_ ptr, _ u64, _ ptr) u64 { return 0 }
