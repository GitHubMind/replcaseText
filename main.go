package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"replaceText/lib"
)

var (
	reg   lib.RegFileAddes
	tra   lib.Transale
	roots []string
)

func init() {
	reg = lib.RegFileAddes{lib.Config.NotIgnoreAddress, nil, nil}
	tra = lib.Transale{lib.Config.TransaleAddresJson, nil, nil}
	//有两种读取的方式
	if len(lib.Config.TransaleAddres) > 0 {
		//tra.OpenFileTxt(lib.Config.TransaleAddres, lib.Config.Language)
	}
	tra.ReadJsonFile(lib.Config.Language)
	reg.ReadJsonFile()
	roots = lib.Config.Catalog_address //要查询的地址
}
func main() {
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes, root)
		}
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}
func printDiskUsage(nfiles, nbytes int64) {
	///1e9
	fmt.Printf("%d files  %.1f \n", nfiles, float64(nbytes))
}
func walkDir(dir string, fileSizes chan<- int64, root string) {
	for _, entry := range dirents(dir) {
		// 这个是

		if entry.IsDir() {
			//if reg.RegName(dir, entry.Name()+"/", root) == false {
			//	continue
			//}
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes, root)
		} else {
			if reg.RegName(dir, entry.Name(), root) == false {
				continue
			}

			fileName := dir + "/" + entry.Name()
			lib.OpenFile(fileName, tra)
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
