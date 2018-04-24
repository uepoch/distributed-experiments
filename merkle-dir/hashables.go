package merkle_dir

import (
	"crypto"
	"fmt"
)
type HashableStr string
type HashableInt int

func weakHashing(v interface{}, hash crypto.Hash) ([]byte){
	h := hash.New()
	h.Write([]byte(fmt.Sprint(v)))
	return h.Sum(nil)
}

func (s HashableStr) Hash(hash crypto.Hash) []byte {
	return weakHashing(s, hash)
}

func (i HashableInt) Hash(hash crypto.Hash) []byte{
	return weakHashing(i, hash)
}