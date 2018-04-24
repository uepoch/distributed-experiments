package utils

import (
	"math/rand"
)

const defaultCharset = "abcdefghijklmnopqrstuvwxyz"

func RandStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		rd := rand.Intn(len(charset))
		b[i] = charset[rd]
		if i%2 == 0 && i < length-2 {
			b[i+1] = '/'
			i++
		}

	}
	return string(b)
}

func RandString(length int) string {
	return RandStringWithCharset(length, defaultCharset)
}

func StringGenerator(length int, bufSize int) (chan string, chan struct{}){
	return StringGeneratorWithCharset(length, bufSize, defaultCharset)
}

func StringGeneratorWithCharset(length int, bufSize int, charset string) (chan string, chan struct{}) {
	out := make(chan string, bufSize)
	err := make(chan struct{})

	go func() {
		for {
			select {
			case <-err:
				return
			case out <- RandStringWithCharset(length, charset):
			}
		}
	}()

	return out, err
}

//func RandomUpdater(m1, m2 *merkle_dir.MerklePath, strCh chan string, errCh chan struct{}) {
//	i := 0
//	a := []*merkle_dir.MerklePath{m1, m2}
//	for {
//		select {
//		case <-errCh:
//			return
//		default:
//			a[i%2].Update(<-strCh, merkle_dir.HashableStr(strconv.Itoa(i)))
//			i++
//		}
//	}
//}
