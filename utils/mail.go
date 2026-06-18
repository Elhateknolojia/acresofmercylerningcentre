package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

func RelayMail(endpoint string, payload interface{}) error {
    data, _ := json.Marshal(payload)

    req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
    if err != nil {
        return err
    }

    // ✅ Tell PHP this is JSON
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("relay error: %s", resp.Status)
    }
    return nil
}
