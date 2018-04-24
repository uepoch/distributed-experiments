package main

import (
	"crypto"

	_ "crypto/sha512"
	"fmt"
	"strconv"

	"github.com/uepoch/distributed-experiments/lib/utils"
	"github.com/uepoch/distributed-experiments/merkle-dir"
	"go.uber.org/zap"
	"flag"
	"time"
)

func merkleDump(path *merkle_dir.MerklePath, indent int) {
	p := func(args ...interface{}) {
		for i := 0; i < indent; i++ {
			fmt.Print(" ")
		}
		fmt.Println(args...)
	}
	for i, m := range path.Tree {
		p("-", i, ":")
		p("Name", m.Name())
		merkleDump(m, indent+2)
	}

}

func RandomUpdater(m1, m2 *merkle_dir.MerklePath, strCh chan string, errCh chan struct{}) {
	i := 0
	a := []*merkle_dir.MerklePath{m1, m2}
	for {
		select {
		case <-errCh:
			zap.L().Info("Generator finished", zap.Int("generated", i))
			return
		default:
			a[i%2].Update(<-strCh, merkle_dir.HashableStr(strconv.Itoa(i)))
			i++
		}
	}
}

func main() {
	verbose := flag.Bool("v", false, "Use to trigger more output")

	flag.Parse()

	var l *zap.Logger

	if !*verbose {
		l, _ = zap.NewProduction()
	} else {
		l, _ = zap.NewDevelopment()
	}

	zap.ReplaceGlobals(l)

	merkle, err := merkle_dir.NewMerkleRoot(crypto.SHA512)
	if err != nil {
		fmt.Println(err)
	}

	m2, err := merkle_dir.NewMerkleRoot(crypto.SHA512)
	if err != nil {
		fmt.Println(err)
	}

	if true {
		strCh := make(chan string)

		strCh, errCh := utils.StringGenerator(5, 50)
		for i := 0; i < 3 ; i++ {
			go RandomUpdater(merkle, m2, strCh, errCh)
			go func() {
				for first := false; ; first = !first {
					if first {
						merkle.Compare(m2)
					} else {
						m2.Compare(merkle)
					}
					time.Sleep(200 * time.Millisecond)
				}
			}()
		}

		time.Sleep(5 * time.Second)

		close(errCh)
	}

	merkle.Update("a/b/c/d", merkle_dir.HashableInt(30))
	m2.Update("a/b/c/e", merkle_dir.HashableInt(30))

	//merkleDump(merkle, 0 )

	changes, err := merkle.Compare(m2)
	mod := merkle_dir.Changes{}
	del := merkle_dir.Changes{}
	for _, c := range changes {
		switch c.Action {
		case merkle_dir.ActionCreate:
			fallthrough
		case merkle_dir.ActionUpdate:
			mod = append(mod, c)
		case merkle_dir.ActionDelete:
			del = append(del, c)
		}
	}

	zap.L().Info("Changes needed", zap.Int("number", len(changes)))
	zap.L().Info("Update needed", zap.Int("number", len(mod)))
	zap.L().Info("Delete needed", zap.Int("number", len(del)))
}
