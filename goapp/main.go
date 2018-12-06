package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	minio "github.com/minio/minio-go"
)

func getData(w http.ResponseWriter, r *http.Request) {
	var data = map[string]map[string]string{}
	minioClient, err := minio.New("minio:9000", "accesskey", "secretkey", false)
	if err != nil {
		log.Fatalln(err)
	}
	buckets, err := minioClient.ListBuckets()
	if err != nil {
		log.Fatalln(err)
	}
	for _, bucket := range buckets {
		objects := make(map[string]string)
		name := bucket.Name
		// Create a done channel to control 'ListObjects' go routine.
		doneCh := make(chan struct{})

		// Indicate to our routine to exit cleanly upon return.
		defer close(doneCh)

		isRecursive := true
		objectCh := minioClient.ListObjects(name, "", isRecursive, doneCh)
		reqParams := make(url.Values)
		data[name] = make(map[string]string)
		for object := range objectCh {
			if object.Err != nil {
				log.Fatal(object.Err)
			}
			url, err := minioClient.PresignedGetObject(
				name, object.Key, time.Second*24*60*60, reqParams,
			)
			if err != nil {
				log.Fatalln(err)
			}
			urlString := strings.Replace(url.String(), "http://minio:9000/", "/minio/", -1)
			objects[object.Key] = urlString
		}
		data[name] = objects
	}
	jsonString, _ := json.Marshal(data)
	w.Write([]byte(jsonString))
}

func main() {
	http.HandleFunc("/", getData)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
