package clickup

import (
	"fmt"
	"net/http"
)

type FoldersService service

// ClickUp API docs: https://jsapi.apiary.io/apis/clickup20/reference/0/folders/create-folder.html
func (s *FoldersService) Create(spaceID string, folder *Folder) (*Folder, *http.Response, error) {
	if folder == nil {
		return nil, nil, fmt.Errorf("folder must be provided")
	}

	u := fmt.Sprintf("space/%v/folder", spaceID)
	req, err := s.client.NewRequest(http.MethodPost, u, folder)
	if err != nil {
		return nil, nil, err
	}

	created := new(Folder)
	resp, err := s.client.Do(req, created)
	if err != nil {
		return nil, resp, err
	}
	return created, resp, nil
}

// ClickUp API docs: https://jsapi.apiary.io/apis/clickup20/reference/0/folders/update-folder.html
func (s *FoldersService) Update(folderID string, folder *Folder) (*Folder, *http.Response, error) {
	if folder == nil {
		return nil, nil, fmt.Errorf("folder must be provided")
	}

	u := fmt.Sprintf("folder/%v", folderID)
	req, err := s.client.NewRequest(http.MethodPut, u, folder)
	if err != nil {
		return nil, nil, err
	}

	updated := new(Folder)
	resp, err := s.client.Do(req, updated)
	if err != nil {
		return nil, resp, err
	}
	return updated, resp, nil
}

// ClickUp API docs: https://jsapi.apiary.io/apis/clickup20/reference/0/folders/delete-folder.html
func (s *FoldersService) Delete(folderID string) (*http.Response, error) {
	u := fmt.Sprintf("folder/%v", folderID)
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

type Folders struct {
	Folders []Folder `json:"folders,omitempty"`
}

// ClickUp API docs: https://jsapi.apiary.io/apis/clickup20/reference/0/folders/get-folders.html
func (s *FoldersService) GetAll(spaceID string, archived bool) (*Folders, *http.Response, error) {
	u := fmt.Sprintf("space/%v/folder?archived=%t", spaceID, archived)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	folders := new(Folders)
	resp, err := s.client.Do(req, folders)
	if err != nil {
		return nil, resp, err
	}
	return folders, resp, nil
}

// ClickUp API docs: https://jsapi.apiary.io/apis/clickup20/reference/0/folders/get-folder.html
func (s *FoldersService) Get(folderID string) (*Folder, *http.Response, error) {
	u := fmt.Sprintf("folder/%v", folderID)
	req, err := s.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	folder := new(Folder)
	resp, err := s.client.Do(req, folder)
	if err != nil {
		return nil, resp, err
	}
	return folder, resp, nil
}

type Folder struct {
	ID               *string      `json:"id,omitempty"`
	Name             *string      `json:"name,omitempty"`
	OrderIndex       *int         `json:"orderindex,omitempty"`
	OverrideStatuses *bool        `json:"override_statuses,omitempty"`
	Hidden           *bool        `json:"hidden,omitempty"`
	Space            *FolderSpace `json:"space,omitempty"`
	TaskCount        *string      `json:"task_count,omitempty"`
}

func (f *Folder) GetID() string {
	if f == nil || f.ID == nil {
		return ""
	}
	return *f.ID
}

func (f *Folder) GetName() string {
	if f == nil || f.Name == nil {
		return ""
	}
	return *f.Name
}

func (f *Folder) GetOrderIndex() int {
	if f == nil || f.OrderIndex == nil {
		return -1
	}
	return *f.OrderIndex
}

func (f *Folder) CanOverrideStatuses() bool {
	if f == nil || f.OverrideStatuses == nil {
		return false
	}
	return *f.OverrideStatuses
}

func (f *Folder) IsHidden() bool {
	if f == nil || f.Hidden == nil {
		return false
	}
	return *f.Hidden
}

func (f *Folder) GetSpace() *FolderSpace {
	if f == nil {
		return nil
	}
	return f.Space
}

func (f *Folder) GetTaskCount() string {
	if f == nil || f.TaskCount == nil {
		return ""
	}
	return *f.TaskCount
}

type FolderSpace struct {
	ID     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Access *bool   `json:"access,omitempty"`
}

func (s *FolderSpace) GetID() string {
	if s == nil || s.ID == nil {
		return ""
	}
	return *s.ID
}

func (s *FolderSpace) GetName() string {
	if s == nil || s.Name == nil {
		return ""
	}
	return *s.Name
}

func (s *FolderSpace) CanAccess() bool {
	if s == nil || s.Access == nil {
		return false
	}
	return *s.Access
}
