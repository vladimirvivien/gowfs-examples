package main

import "fmt"
import "flag"
import "log"
import "github.com/vladimirvivien/gowfs"

var nn   = flag.String("namenode", "localhost:50070", "Namenode address")
var path = flag.String("path", "/", "File path")
var user = flag.String("user", "", "HDFS user.")

func main() {
	flag.Parse()

	fs, err := gowfs.NewFileSystem(gowfs.Configuration{Addr: *nn, User: *user})
	if err != nil{
		log.Fatal(err)
	}

	checksum, err := fs.GetFileChecksum(gowfs.Path{Path: *path})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println ("Checksum for ", *path)
	fmt.Println (checksum)
}