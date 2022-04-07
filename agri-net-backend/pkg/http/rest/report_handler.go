package rest

// IReportHandler ...
type IReportHandler interface {
}

// ReportHandler ...
type ReportHandler struct{}

// NewReportHandler ...
func NewReportHandler() IReportHandler {
	return &ReportHandler{}
}
