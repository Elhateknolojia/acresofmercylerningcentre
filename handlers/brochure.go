package handlers

import (
    "context"
    "github.com/chromedp/chromedp"
    "github.com/chromedp/cdproto/page"
    "log"
    "net/http"
    "os"
    "strconv"
    "time"
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
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    ctx, cancelTimeout := context.WithTimeout(ctx, 15*time.Second)
    defer cancelTimeout()

    var pdfBuffer []byte

    err := chromedp.Run(ctx,
        chromedp.Navigate("about:blank"),
        chromedp.ActionFunc(func(ctx context.Context) error {
            frameTree, err := page.GetFrameTree().Do(ctx)
            if err != nil {
                return err
            }
            return page.SetDocumentContent(frameTree.Frame.ID, brochureHTML).Do(ctx)
        }),
        chromedp.WaitReady("body", chromedp.ByQuery),
        chromedp.ActionFunc(func(ctx context.Context) error {
            buf, _, err := page.PrintToPDF().
                WithPrintBackground(true).
                WithPaperWidth(8.27).
                WithPaperHeight(11.69).
                Do(ctx)
            if err != nil {
                return err
            }
            pdfBuffer = buf
            return nil
        }),
    )

    if err != nil {
        log.Printf("Failed to generate PDF: %v", err)
        http.Error(w, "Failed to generate PDF", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/pdf")
    w.Header().Set("Content-Disposition", `attachment; filename="Acres_of_Mercy_Prospectus.pdf"`)
    w.Header().Set("Content-Length", strconv.Itoa(len(pdfBuffer)))
    w.Write(pdfBuffer)
}
