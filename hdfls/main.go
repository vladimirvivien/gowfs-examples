package main


import "fmt"
import "flag"
import "log"
import "os"
import "strconv"
import "time"
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

	statuses, err := fs.ListStatus(gowfs.Path{Path: *path})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d item(s)\n", len(statuses))
	for _, v := range statuses {
		
		fmt.Printf(
			"%-11s %3s %s\t%s\t%11d %20v %s\n", 
			formatFileMode(v.Permission, v.Type), 
			formatReplication(v.Replication, v.Type),
			v.Owner,
			v.Group,
			v.Length,
			formatModTime(v.ModificationTime),
			v.PathSuffix)
	}
}

func formatFileMode(webfsPerm string, fileType string) string {
	perm, _ := strconv.ParseInt(webfsPerm, 8, 16)
	fm := os.FileMode(perm)
	if fileType == "DIRECTORY"{
		fm = fm | os.ModeDir
	}
	return fm.String()
}

func formatReplication(rep int64, fileType string) string {
	repStr := strconv.FormatInt(rep, 8)
	if fileType == "DIRECTORY" {
		repStr = "-"
	}
	return repStr
}

func formatModTime (modTime int64) string {
	modTimeAdj := time.Unix((modTime/1000),0) // adjusted for Java Calendar in millis.
	return modTimeAdj.Format("2006-01-02 15:04:05")
}