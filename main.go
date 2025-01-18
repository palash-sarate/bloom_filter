package bloomFilter

import (
	"fmt"
)

func main() {
	bf := NewBloomFilter(100, 3, nil)
	init_strings := []string{"hello", "world", "foo", "bar", "baz"}
	check_strings := []string{"hello", "world", "foo", "bar", "palash"}

	for _, s := range init_strings {
		bf.Add([]byte(s))
	}
	for _, s := range check_strings {
		fmt.Println(bf.Contains([]byte(s)))
	}
}