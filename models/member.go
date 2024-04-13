package models

type Member struct {
	ID        int    `json:"id"`
	ProjectID int    `json:"project_id"`
	Active    bool   `json:"active"`
	Role      string `json:"role"`
}

type Members []Member

func (m *Members) GetMembers() Members {
	return *m
}

func (m *Members) GetMemberByID(id int) *Member {
	for _, member := range *m {
		if member.ID == id {
			return &member
		}
	}
	return nil
}

func (m *Members) AddMember(member Member) Members {
	*m = append(*m, member)
	return *m
}
