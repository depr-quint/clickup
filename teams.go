package clickup

import (
	"fmt"
	"net/http"
)

type TeamsService service

func (s *TeamsService) Get() (*Teams, *http.Response, error) {
	u := fmt.Sprintf("team")
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	teams := new(Teams)
	resp, err := s.client.Do(req, teams)
	if err != nil {
		return nil, resp, err
	}

	return teams, resp, nil
}

type Teams struct {
	Teams []Team `json:"teams,omitempty"`
}

type Team struct {
	ID      *string  `json:"id,omitempty"`
	Name    *string  `json:"name,omitempty"`
	Color   *string  `json:"color,omitempty"`
	Avatar  *string  `json:"avatar,omitempty"`
	Members []Member `json:"members,omitempty"`
}

func (t *Team) GetID() string {
	if t == nil || t.ID == nil {
		return ""
	}
	return *t.ID
}

func (t *Team) GetName() string {
	if t == nil || t.Name == nil {
		return ""
	}
	return *t.Name
}

func (t *Team) GetColor() string {
	if t == nil || t.Color == nil {
		return ""
	}
	return *t.Color
}

func (t *Team) GetAvatar() string {
	if t == nil || t.Avatar == nil {
		return ""
	}
	return *t.Avatar
}
