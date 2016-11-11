package storageservice

// StorageContext storage context
type StorageContext struct {
	Project string
	Bucket  string
	Object  string
}

// SetCurrentProject set current project
func (context *StorageContext) SetCurrentProject(project string) *StorageContext {
	context.Project = project	
	return context
}

// SetCurrentBucket set current bucket
func (context *StorageContext) SetCurrentBucket(bucket string) *StorageContext {
	context.Bucket = bucket	
	return context
}

// SetCurrentObject set current object
func (context *StorageContext) SetCurrentObject(object string) *StorageContext {
	context.Object = object	
	return context
}