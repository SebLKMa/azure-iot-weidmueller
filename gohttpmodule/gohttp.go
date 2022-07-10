package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	mux "github.com/gorilla/mux"
	"golang.org/x/sys/unix"
)

// defaultHandler is a http request handler for route / .
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	helloMsg := "Hello gohttp " + currentTime.Format("2006.01.02 15:04:05") + "\n"
	log.Printf(helloMsg)

	writeToFile(helloMsg)

	w.Write([]byte(helloMsg))
}

func writeToFile(msg string) {
	outdir := "/data" // docker volume
	outfilepath := outdir + "/output.txt"

	// check disk space, unix only, not windows
	var stat unix.Statfs_t
	wd, err := os.Getwd()
	unix.Statfs(wd, &stat)

	// Available blocks * size per block = available space in bytes
	// i.e. One can check How many bytes left before continuing...
	//
	bytesLeft := stat.Bavail * uint64(stat.Bsize)
	bytesLeftMsg := fmt.Sprintf("Disk has %d bytes left.\n", bytesLeft)
	log.Println(bytesLeftMsg)

	// append output to file
	f, err := os.OpenFile(outfilepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("Error opening file %s\n", outfilepath)
		log.Println(err)
		return
	}
	defer f.Close()

	if _, err = f.WriteString(bytesLeftMsg); err != nil {
		log.Printf("Error writing file %s\n", outfilepath)
		log.Println(err)
		return
	}
	if _, err = f.WriteString(msg); err != nil {
		log.Printf("Error writing file %s\n", outfilepath)
		log.Println(err)
		return
	}

	//err := ioutil.WriteFile(outfilepath, output, 0644)
	//if err != nil {
	//	log.Printf("Error writing to %s\n", outfilepath)
	//	return
	//}
	log.Printf("Successfully written to %s\n", outfilepath)
}

func main() {
	// Setting up a simple HTTP REST /ping request
	portPtr := flag.String("port", "8181", "port number")
	flag.Parse()

	httpPort := *portPtr
	httpURL := "0.0.0.0:" + httpPort
	log.Printf("HTTP %s up and listening...\n", httpURL)

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	//r.HandleFunc("/ping", pingHandler)
	r.HandleFunc("/", defaultHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(httpURL, r))
}
