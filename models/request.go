package models

type AdmissionRequest struct {
    ParentName string `json:"admParent"`
    Email      string `json:"admEmail"`
    Phone      string `json:"admPhone"`
    Level      string `json:"admLevel"`
    Student    string `json:"admStudent"`
    Message    string `json:"admComments"`
}
