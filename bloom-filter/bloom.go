package main

import (
	"fmt"
	"math"
	"math/rand"
)

type hashFunc func([]byte) int32

type bloomFilter struct {
	m      int
	k      int
	bits   []byte
	hashes []hashFunc
}

func newBloomFilter(m, k int) *bloomFilter {
	return &bloomFilter{
		m:      m,
		k:      k,
		bits:   make([]byte, int(math.Ceil(float64(m)/8.0))),
		hashes: generateHashes(k),
	}

}

func (f *bloomFilter) add(x string) {
	for _, hash := range f.hashes {
		res := hash([]byte(x))
		res = res % int32(f.m)
		if res < 0 {
			res += int32(f.m)
		}

		indx := res / 8
		shift := 7 - (res - indx*8)
		mask := 1 << uint(shift)

		f.bits[indx] = f.bits[indx] | byte(mask)
	}
}

func (f *bloomFilter) check(x string) bool {
	for _, hash := range f.hashes {
		res := hash([]byte(x))
		res = res % int32(f.m)
		if res < 0 {
			res += int32(f.m)
		}

		indx := res / 8
		shift := 7 - (res - indx*8)
		mask := 1 << uint(shift)

		present := f.bits[indx] & byte(mask)
		if present == 0 {
			return false
		}
	}
	return true
}

func generateHashes(k int) []hashFunc {
	var hh []hashFunc
	for i := 0; i < k; i++ {
		seed := rand.Int31()
		hh = append(hh, func(bb []byte) int32 {
			res := int64(1)
			for _, b := range bb {
				res += (int64(seed)*int64(res) + int64(b)) & 0xFFFFFFFF
			}

			return int32(res)
		})
	}
	return hh
}

func main() {
	filter := newBloomFilter(60, 2)
	toInsert := []string{"apple", "orange", "melon", "strawberry"}
	toTest := []string{"apple", "pineapple", "watermelon", "orange", "raspberry", "kiwi"}

	for _, insert := range toInsert {
		filter.add(insert)
		fmt.Printf("Inserted %q to set\n", insert)
	}

	for _, test := range toTest {
		fmt.Printf("Check %q: %t\n", test, filter.check(test))
	}
}
