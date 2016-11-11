package main

import (
	"flag"
	"log"
	"storageservice"

	"fmt"
	"io/ioutil"
	"sync"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	storage "google.golang.org/api/storage/v1"
)

const (
	// This scope allows the application full control over resources in Google Cloud Storage
	scope = storage.DevstorageFullControlScope
)

var (
	projectID  = flag.String("project", "artificial-intelligence-ade", "define project")
	bucketName = flag.String("bucket", "lms_archieve", "define bucket")
	localDir   = flag.String("localdirectory", "c:\\Apache24\\htdocs\\lms\\public\\pdf", "local directory")
)

func main() {
	flag.Parse()
	if *bucketName == "" {
		log.Fatalf("Bucket argument is required. See --help.")
	}
	if *projectID == "" {
		log.Fatalf("Project argument is required. See --help.")
	}
	if *localDir == "" {
		log.Fatalf("File argument is required. See --help.")
	}

	// Authentication is provided by the gcloud tool when running locally, and
	// by the associated service account when running on Compute Engine.
	client, err := google.DefaultClient(context.Background(), scope)
	if err != nil {
		log.Fatalf("Unable to get default client: %v", err)
	}

	service, err := storage.New(client)
	if err != nil {
		log.Fatalf("Unable to create storage service: %v", err)
	}

	context := &storageservice.StorageContext{}
	context.SetCurrentProject(*projectID)
	context.SetCurrentBucket(*bucketName)

	accessor := &storageservice.StorageAccessor{}
	accessor.SetCurrentStorageContext(context)
	accessor.SetLocalDirectory(*localDir)

	storageservice := &storageservice.StorageService{}
	storageservice.SetCurrentStorageAccessor(accessor)

	// read the pdfs in pdf directory
	files, err := ioutil.ReadDir(*localDir)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(files))

	fmt.Printf("%d files to move", len(files))

	for _, file := range files {
		// go routine to move
		go func(file string) {
			defer wg.Done()
			if !storageservice.Move(service, file) {
				fmt.Println("Move failed!")
			} else {
				fmt.Printf("%q Moved to storage", file)
			}
		}(file.Name())
	}

	wg.Wait();
	fmt.Println("done");
}
