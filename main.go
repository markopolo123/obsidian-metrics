package main

import (
	"flag"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var addr = flag.String("listen-address", ":8090", "The address to listen on for HTTP requests.")
var files = flag.String("vault-collection", "./", "location of your files to check")
var tag = flag.String("tag", "#unprocessed", "tag to check for")

var count int

// var c int

func main() {
	flag.Parse()
	var (
		unprocessedNotes = prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "total_unprocessed_notes",
			Help: "The total number of unprocessed notes",
		})
	)
	prometheus.MustRegister(unprocessedNotes)
	go func() {
		for {
			unprocessedNotes.Set(float64(returncount()))
		}
	}()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func returncount() int {
	count = 0
	filepath.WalkDir(*files, visit)
	return count
}

func visit(path string, di fs.DirEntry, err error) error {
	var tagged bool
	fileInfo, err := os.Lstat(path)
	if fileInfo.Mode().IsRegular() {

		b, err := os.ReadFile(path)

		if err != nil {
			// fmt.Println(err)
			panic(err)
		}
		s := string(b)
		tagged = (strings.Contains(s, *tag))
		if tagged == true {
			count++
		}
	}
	return nil
}
