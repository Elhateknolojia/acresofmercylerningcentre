package handlers

import (
    // "context"
    // "github.com/chromedp/chromedp"
    // "github.com/chromedp/cdproto/page"
    "log"
    "net/http"
    "os"
    "strconv"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
    // "time"
)

var brochureHTML string

func init() {
    data, err := os.ReadFile("attachments/brochure.html")
    if err != nil {
        log.Printf("Failed to load brochure: %v", err)
        brochureHTML = "<html><body><h1>Brochure not found</h1></body></html>"
    } else {
        brochureHTML = string(data)
    }
}


func DownloadBrochureHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("[DownloadBrochureHandler] Request received from:", r.RemoteAddr)

    // Load HTML file
    log.Println("[DownloadBrochureHandler] Attempting to open brochure.html...")
    f, err := os.Open("attachments/brochure.html")
    if err != nil {
        log.Printf("[DownloadBrochureHandler] ERROR opening brochure.html: %v", err)
        http.Error(w, "Brochure not found", http.StatusInternalServerError)
        return
    }
    defer func() {
        log.Println("[DownloadBrochureHandler] Closing brochure.html file handle")
        f.Close()
    }()

    // Create new PDF generator
    log.Println("[DownloadBrochureHandler] Initializing wkhtmltopdf generator...")
    pdfg, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        log.Printf("[DownloadBrochureHandler] ERROR initializing PDF generator: %v", err)
        http.Error(w, "Failed to init PDF generator", http.StatusInternalServerError)
        return
    }

    // Add input HTML file
    log.Println("[DownloadBrochureHandler] Adding brochure.html as input page...")
    pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

    // Set options
    log.Println("[DownloadBrochureHandler] Setting PDF options (DPI=300, A4, Portrait)...")
    pdfg.Dpi.Set(300)
    pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
    pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

    // Create PDF
    log.Println("[DownloadBrochureHandler] Generating PDF...")
    err = pdfg.Create()
    if err != nil {
        log.Printf("[DownloadBrochureHandler] ERROR generating PDF: %v", err)
        http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
        return
    }
    log.Println("[DownloadBrochureHandler] PDF generation successful")

    buf := pdfg.Bytes()
    log.Printf("[DownloadBrochureHandler] PDF size: %d bytes", len(buf))

    // Write response headers
    log.Println("[DownloadBrochureHandler] Writing response headers...")
    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", `attachment; filename="Acres_of_Mercy_Prospectus.pdf"`)
    w.Header().Set("Content-Length", strconv.Itoa(len(buf)))

    // Send PDF to client
    log.Println("[DownloadBrochureHandler] Sending PDF to client...")
    _, writeErr := w.Write(buf)
    if writeErr != nil {
        log.Printf("[DownloadBrochureHandler] ERROR writing PDF to response: %v", writeErr)
    } else {
        log.Println("[DownloadBrochureHandler] PDF successfully sent to client")
    }
}