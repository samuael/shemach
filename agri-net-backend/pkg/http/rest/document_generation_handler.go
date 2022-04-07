package rest

type IDocumentGenerationHandler interface {
	// Generate students certificate
	// generate Daily Registration Result
	// generate rounds registration result
	// generate monthly registration result
	// backup category inforamtion
	// backup round
	// Generate Student ID
	// Generate Id for all students of a round
	// generate certificate for all students of a round. who has paid this much money.
	// Generate certificate for a student.
	//
}

type DocumentGenerationHandler struct {
}

func NewDocumentGenerationHandler() IDocumentGenerationHandler {
	return DocumentGenerationHandler{}
}
