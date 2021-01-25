package some

import (
	"fmt"
	"testing"
)

func TestLruCache_Hash(t *testing.T) {

	fmt.Println(hash(600))

	fmt.Println(1 << 3)
	fmt.Println(8 >> 3)

	x := 10
	fmt.Println(x&1, x%2)
}