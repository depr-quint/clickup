package clickup

import (
	"fmt"
	"net/http"
)

type SpacesService service

// ClickUp API docs: https://jsapi.apiary.io/apis/clickup20/reference/0/spaces/create-space.html
func (s *SpacesService) Create(teamID string, space *Space) (*Space, *http.Response, error) {
	if space == nil {
		return nil, nil, fmt.Errorf("space must be provided")
	}

	u := fmt.Sprintf("team/%v/space", teamID)
	req, err := s.client.NewRequest(http.MethodPost, u, space)
	if err != nil {
		return nil, nil, err
	}

	created := new(Space)
	resp, err := s.client.Do(req, created)
	if err != nil {
		return nil, resp, err
	}
	return created, resp, nil
}

// ClickUp API docs: https://jsapi.apiary.io/apis/clickup20/reference/0/spaces/update-space.html
func (s *SpacesService) Update(spaceID string, space *Space) (*Space, *http.Response, error) {
	if space == nil {
		return nil, nil, fmt.Errorf("space must be provided")
	}

	u := fmt.Sprintf("space/%v", spaceID)
	req, err := s.client.NewRequest(http.MethodPut, u, space)
	if err != nil {
		return nil, nil, err
	}

	created := new(Space)
	resp, err := s.client.Do(req, created)
	if err != nil {
		return nil, resp, err
	}
	return created, resp, nil
}

// ClickUp API docs: https://jsapi.apiary.io/apis/clickup20/reference/0/spaces/delete-space.html
func (s *SpacesService) Delete(spaceID string) (*http.Response, error) {
	u := fmt.Sprintf("space/%v", spaceID)
	req, err := s.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type Spaces struct {
	Spaces []Spaces `json:"spaces,omitempty"`
}

func (s *SpacesService) GetAll(teamID string, archived bool) (*Spaces, *http.Response, error) {
	u := fmt.Sprintf("team/%v/space?archived=%t", teamID, archived)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	spaces := new(Spaces)
	resp, err := s.client.Do(req, spaces)
	if err != nil {
		return nil, resp, err
	}
	return spaces, resp, nil
}

func (s *SpacesService) Get(spaceID string) (*Space, *http.Response, error) {
	u := fmt.Sprintf("space/%v", spaceID)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	space := new(Space)
	resp, err := s.client.Do(req, space)
	if err != nil {
		return nil, resp, err
	}
	return space, resp, nil
}

type Space struct {
	ID                *string   `json:"id,omitempty"`
	Name              *string   `json:"name,omitempty"`
	Private           *bool     `json:"private,omitempty"`
	Statuses          []Status  `json:"statuses,omitempty"`
	MultipleAssignees *bool     `json:"multiple_assignees,omitempty"`
	Features          *Features `json:"features,omitempty"`
}

func (s *Space) GetID() string {
	if s == nil || s.ID == nil {
		return ""
	}
	return *s.ID
}

func (s *Space) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}

func (s *Space) IsPrivate() bool {
	if s == nil || s.Private == nil {
		return false
	}
	return *s.Private
}

func (s *Space) HasMultipleAssignees() bool {
	if s == nil || s.MultipleAssignees == nil {
		return false
	}
	return *s.MultipleAssignees
}

func (s *Space) GetFeatures() *Features {
	if s == nil {
		return nil
	}
	return s.Features
}

type Status struct {
	Status     *string `json:"status,omitempty"`
	Type       *string `json:"type,omitempty"`
	OrderIndex *int    `json:"orderindex,omitempty"`
	Color      *string `json:"color,omitempty"`
}

func (s *Status) GetStatus() string {
	if s == nil || s.Status == nil {
		return ""
	}
	return *s.Status
}

func (s *Status) GetType() string {
	if s == nil || s.Type == nil {
		return ""
	}
	return *s.Type
}

func (s *Status) GetOrderIndex() int {
	if s == nil || s.OrderIndex == nil {
		return -1
	}
	return *s.OrderIndex
}

func (s *Status) GetColor() string {
	if s == nil || s.Color == nil {
		return ""
	}
	return *s.Color
}

type Features struct {
	DueDates *struct {
		Enabled            *bool `json:"enabled,omitempty"`
		StartDate          *bool `json:"start_date,omitempty"`
		RemapDueDates      *bool `json:"remap_due_dates,omitempty"`
		RemapClosedDueDate *bool `json:"remap_closed_due_date,omitempty"`
	} `json:"due_dates,omitempty"`
	TimeTracking *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"time_tracking,omitempty"`
	Tags *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"tags,omitempty"`
	TimeEstimates *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"time_estimates,omitempty"`
	Checklists *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"checklists,omitempty"`
	CustomFields *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"custom_fields,omitempty"`
	RemapDependencies *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"remap_dependencies,omitempty"`
	DependencyWarning *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"dependency_warning,omitempty"`
	Portfolios *struct {
		Enabled *bool `json:"enabled,omitempty"`
	} `json:"portfolios,omitempty"`
}

func (f *Features) DueDatesEnabled() bool {
	if f == nil || f.DueDates == nil || f.DueDates.Enabled == nil {
		return false
	}
	return *f.DueDates.Enabled
}

func (f *Features) StartDateEnabled() bool {
	if f == nil || f.DueDates == nil || f.DueDates.StartDate == nil {
		return false
	}
	return *f.DueDates.StartDate
}

func (f *Features) RemapDueDatesEnabled() bool {
	if f == nil || f.DueDates == nil || f.DueDates.RemapDueDates == nil {
		return false
	}
	return *f.DueDates.RemapDueDates
}

func (f *Features) RemapClosedDueDateEnabled() bool {
	if f == nil || f.DueDates == nil || f.DueDates.RemapClosedDueDate == nil {
		return false
	}
	return *f.DueDates.RemapClosedDueDate
}

func (f *Features) TimeTrackingEnabled() bool {
	if f == nil || f.TimeTracking == nil || f.TimeTracking.Enabled == nil {
		return false
	}
	return *f.TimeTracking.Enabled
}

func (f *Features) TagsEnabled() bool {
	if f == nil || f.Tags == nil || f.Tags.Enabled == nil {
		return false
	}
	return *f.Tags.Enabled
}

func (f *Features) TimeEstimatesEnabled() bool {
	if f == nil || f.TimeEstimates == nil || f.TimeEstimates.Enabled == nil {
		return false
	}
	return *f.TimeEstimates.Enabled
}

func (f *Features) ChecklistsEnabled() bool {
	if f == nil || f.Checklists == nil || f.Checklists.Enabled == nil {
		return false
	}
	return *f.Checklists.Enabled
}

func (f *Features) CustomFieldsEnabled() bool {
	if f == nil || f.CustomFields == nil || f.CustomFields.Enabled == nil {
		return false
	}
	return *f.CustomFields.Enabled
}

func (f *Features) RemapDependenciesEnabled() bool {
	if f == nil || f.RemapDependencies == nil || f.RemapDependencies.Enabled == nil {
		return false
	}
	return *f.RemapDependencies.Enabled
}

func (f *Features) DependencyWarningEnabled() bool {
	if f == nil || f.DependencyWarning == nil || f.DependencyWarning.Enabled == nil {
		return false
	}
	return *f.DependencyWarning.Enabled
}

func (f *Features) PortfoliosEnabled() bool {
	if f == nil || f.Portfolios == nil || f.Portfolios.Enabled == nil {
		return false
	}
	return *f.Portfolios.Enabled
}
