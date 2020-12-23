package clickup

type Role struct {
	ID     *int    `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Custom *bool   `json:"custom,omitempty"`
}

func (r *Role) GetID() int {
	if r == nil || r.ID == nil {
		return -1
	}
	return *r.ID
}

func (r *Role) GetName() string {
	if r == nil || r.Name == nil {
		return ""
	}
	return *r.Name
}

func (r *Role) IsCustom() bool {
	if r == nil || r.Custom == nil {
		return false
	}
	return *r.Custom
}
