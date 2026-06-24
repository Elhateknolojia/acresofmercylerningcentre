package handlers

import (
    "bytes"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func DownloadBrochureHandler(w http.ResponseWriter, r *http.Request) {
    log.Println("[DownloadBrochureHandler] Request received")

    // Load brochure HTML
    htmlContent, err := os.ReadFile("attachments/brochure.html")
    if err != nil {
        log.Printf("ERROR reading brochure.html: %v", err)
        http.Error(w, "Brochure not found", http.StatusInternalServerError)
        return
    }

    // Send HTML to Node.js PDF service
    log.Println("[DownloadBrochureHandler] Sending HTML to Node.js service...")
    resp, err := http.Post("https://pdf-service-rudh.onrender.com/generate-pdf", "text/html", bytes.NewReader(htmlContent))
    if err != nil {
        log.Printf("ERROR calling PDF service: %v", err)
        http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    pdfBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("ERROR reading PDF response: %v", err)
        http.Error(w, "Failed to read PDF", http.StatusInternalServerError)
        return
    }

    // Forward PDF to client
    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", `attachment; filename="Acres_of_Mercy_Prospectus.pdf"`)
    w.Header().Set("Content-Length", strconv.Itoa(len(pdfBytes)))
    w.Write(pdfBytes)

    log.Println("[DownloadBrochureHandler] PDF successfully sent to client")
}
