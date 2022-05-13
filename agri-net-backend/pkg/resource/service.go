package resource

// IResourceService interface representing the main crop type
type IResourceService interface {
}

// ResourceService ...
type ResourceService struct {
	Repo IResourceService
}
