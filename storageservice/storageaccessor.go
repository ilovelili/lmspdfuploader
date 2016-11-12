package storageservice

import (
	"fmt"
	"log"
	"os"

	storage "google.golang.org/api/storage/v1"
)

// StorageAccessor storage accessor
type StorageAccessor struct {
	StorageContext *StorageContext
	LocalDirectory string
}

func fatalf(errorMessage string, args ...interface{}) {
	log.Fatalf("Dying with error:\n"+errorMessage, args...)
}

// SetCurrentStorageContext inject storagecontext
func (accessor *StorageAccessor) SetCurrentStorageContext(context *StorageContext) *StorageAccessor {
	accessor.StorageContext = context
	return accessor
}

// SetLocalDirectory set local directory
func (accessor *StorageAccessor) SetLocalDirectory(localdirectory string) *StorageAccessor {
	accessor.LocalDirectory = localdirectory
	return accessor
}

// Move move local object to storage
func (accessor *StorageAccessor) Move(service *storage.Service, shortfilename string) bool {
	_storagecontext := *(accessor.StorageContext)
	// all empty
	if (StorageContext{}) == _storagecontext {
		fatalf("No storage context")
		return false
	}

	longfilename := accessor.LocalDirectory + shortfilename
	if len(_storagecontext.Project) == 0 {
		fmt.Println("project not assigned")
		return false
	}

	if len(_storagecontext.Bucket) == 0 {
		fmt.Println("bucket not assigned")
		return false
	}

	object := &storage.Object{Name: shortfilename}
	file, err := os.Open(longfilename)
	if err != nil {
		fatalf("Error opening %q: %v", longfilename, err)
	}
	if res, err := service.Objects.Insert(_storagecontext.Bucket, object).Media(file).Do(); err == nil {
		fmt.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
		// great, unlink the localfile
		os.Remove(longfilename)
		return true
	} else {
		fatalf("Objects.Insert failed: %v", err)
		return false
	}
}

// Copy copy localfile to storage
func (accessor *StorageAccessor) Copy(service *storage.Service, shortfilename string) bool {
	_storagecontext := *(accessor.StorageContext)
	// all empty
	if (StorageContext{}) == _storagecontext {
		fatalf("No storage context")
		return false
	}

	longfilename := accessor.LocalDirectory + shortfilename
	if len(_storagecontext.Project) == 0 {
		fmt.Println("project not assigned")
		return false
	}

	if len(_storagecontext.Bucket) == 0 {
		fmt.Println("bucket not assigned")
		return false
	}

	object := &storage.Object{Name: shortfilename}
	file, err := os.Open(longfilename)
	if err != nil {
		fatalf("Error opening %q: %v", longfilename, err)
	}
	if res, err := service.Objects.Insert(_storagecontext.Bucket, object).Media(file).Do(); err == nil {
		fmt.Printf("Created object %v at location %v\n\n", res.Name, res.SelfLink)
		return true
	}

	fatalf("Objects.Insert failed: %v", err)
	return false
}
