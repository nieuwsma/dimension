package api

import "fmt"

// RFC 7807 compliant error payload.  All fields are optional except the 'type' field.
type Problem7807 struct {
	Type_    string `json:"type"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
	Status   int    `json:"status,omitempty"`
	Title    string `json:"title,omitempty"`
}

func (obj Problem7807) Equals(other Problem7807) (equal bool) {

	if obj.Type_ == other.Type_ &&
		obj.Detail == other.Detail &&
		obj.Instance == other.Instance &&
		obj.Status == other.Status &&
		obj.Title == other.Title {

		return true
	}
	return false
}

func (obj Problem7807) Error() string {
	return fmt.Sprintf("error - status: %v, detail: %v", obj.Status, obj.Detail)
}
