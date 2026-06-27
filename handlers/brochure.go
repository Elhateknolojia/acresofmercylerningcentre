package handlers

import (
    "log"
    "net/http"
    "os"
    "strconv"
)

func DownloadBrochureHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("[DownloadBrochureHandler] Request received")

    // Path to pre-compressed brochure PDF
    filePath := "attachments/school_brochure_compressed.pdf"

    // Open the file
    f, err := os.Open(filePath)
    if err != nil {
        log.Printf("ERROR opening brochure PDF: %v", err)
        http.Error(w, "Brochure not found", http.StatusInternalServerError)
        return
    }
    defer f.Close()

    // Get file info for size
    fi, err := f.Stat()
    if err != nil {
        log.Printf("ERROR getting file info: %v", err)
        http.Error(w, "Failed to read brochure file", http.StatusInternalServerError)
        return
    }

    // Set headers
    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", `attachment; filename="Acres_of_Mercy_Prospectus.pdf"`)
    w.Header().Set("Content-Length", strconv.FormatInt(fi.Size(), 10))

    // Stream file to client
    http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)

    log.Println("[DownloadBrochureHandler] PDF successfully sent to client")
}
