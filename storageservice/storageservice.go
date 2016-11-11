package storageservice

import (
    storage "google.golang.org/api/storage/v1"
)

// StorageService storage service
type StorageService struct {
    StorageAccessor *StorageAccessor 
}

// SetCurrentStorageAccessor inject storageaccessor
func (service *StorageService) SetCurrentStorageAccessor(accessor *StorageAccessor) *StorageService {
	service.StorageAccessor = accessor;
	return service
}

// Move move file to storage
func (service *StorageService) Move(storageservice *storage.Service, shortfilename string) bool {
    return service.StorageAccessor.Move(storageservice, shortfilename)
}