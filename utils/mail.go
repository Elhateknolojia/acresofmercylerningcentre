package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

func RelayMail(endpoint string, payload interface{}) error {
    data, _ := json.Marshal(payload)
    resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(data))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("relay error: %s", resp.Status)
    }
    return nil
}
