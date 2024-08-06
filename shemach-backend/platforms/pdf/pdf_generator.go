package pdf

import (
	"os"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/samuael/shemach/shemach-backend/platforms/helper"
)

// GetThePdf function
func GetThePdf(fileDirectory string) string {
	pdfg, erra := wkhtmltopdf.NewPDFGenerator()
	if erra != nil {
		return ""
	}
	pdfg.Dpi.Set(30)
	pdfg.ImageDpi.Set(30)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	homepath := os.Getenv("CAR_INSPECTION_ASSETS_DIRECTORY") + "pdf/" + helper.GenerateRandomString(5, helper.CHARACTERS) + ".pdf"
	page := wkhtmltopdf.NewPage(fileDirectory)
	// Set options for this page
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	page.Zoom.Set(0.95)
	pdfg.AddPage(page)
	pdfCreationError := pdfg.Create()
	if pdfCreationError != nil {
		println(pdfCreationError.Error())
		return ""
	}
	// Generating Random Name to Be Output Name
	writingError := pdfg.WriteFile(homepath)
	if writingError != nil {
		return ""
	}
	os.Remove(fileDirectory)
	return homepath
}
