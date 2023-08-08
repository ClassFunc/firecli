package firecli

import (
	"log"
	"sync"
)

// firebase instance
var (
	once sync.Once
)

func New(projectId, serviceAccountFilePath string) {
	var err error

	singletonApp(projectId, serviceAccountFilePath)

	// App
	if App == nil {
		log.Fatal("App is nil")
	}

	// Auth
	Auth, err = GetAuth()
	if err != nil {
		log.Fatal(err)
	}

	// Firestore, DB
	Firestore, err = GetFirestoreClient()
	if err != nil {
		log.Fatal(err)
	}
	DB = Firestore

	// Storage, DefaultBucket
	Storage, err = GetStorageClient()
	if err != nil {
		log.Fatal(err)
	}
	DefaultBucket, err = Storage.DefaultBucket()
	if err != nil {
		log.Fatal(err)
	}
}

// singletonApp returns the singleton instance
func singletonApp(projectId, saFilePath string) {
	once.Do(func() { // <-- atomic, does not allow repeating
		App = newApp(projectId, saFilePath)
	})
}
