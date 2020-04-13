package clickup

type Member struct {
	User *User `json:"user,omitempty"`
}

func (m *Member) GetUser() *User {
	if m == nil {
		return nil
	}
	return m.User
}

type User struct {
	ID             *int    `json:"id,omitempty"`
	Username       *string `json:"username,omitempty"`
	Email          *string `json:"email,omitempty"`
	Color          *string `json:"color,omitempty"`
	ProfilePicture *string `json:"profilePicture,omitempty"`
	Initials       *string `json:"initials,omitempty"`
	Role           *int    `json:"role,omitempty"`
}

func (u *User) GetID() int {
	if u == nil || u.ID == nil {
		return -1
	}
	return *u.ID
}

func (u *User) GetUsername() string {
	if u == nil || u.Username == nil {
		return ""
	}
	return *u.Username
}

func (u *User) GetEmail() string {
	if u == nil || u.Email == nil {
		return ""
	}
	return *u.Email
}

func (u *User) GetColor() string {
	if u == nil || u.Color == nil {
		return ""
	}
	return *u.Color
}

func (u *User) GetProfilePicture() string {
	if u == nil || u.ProfilePicture == nil {
		return ""
	}
	return *u.Color
}

func (u *User) GetInitials() string {
	if u == nil || u.Initials == nil {
		return ""
	}
	return *u.Initials
}

func (u *User) GetRole() int {
	if u == nil || u.Role == nil {
		return -1
	}
	return *u.Role
}
