package crop

// ICropService interface representing the main crop type
type ICropService interface {
}

// CropService ...
type CropService struct {
	Repo ICropService
}
