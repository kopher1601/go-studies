package domain

import "errors"

var ErrIllegalState = errors.New("illegal state")

type MemberStatus string

const (
	MemberStatusPending     MemberStatus = "PENDING"
	MemberStatusActive      MemberStatus = "ACTIVE"
	MemberStatusDeactivated MemberStatus = "DEACTIVATED"
)

type Member struct {
	email        string
	nickname     string
	passwordHash string
	status       MemberStatus
}

func NewMember(email, nickname, passwordHash string) *Member {
	return &Member{
		email:        email,
		nickname:     nickname,
		passwordHash: passwordHash,
		status:       MemberStatusPending,
	}
}

func (m *Member) Activate() error {
	if m.status != MemberStatusPending {
		return ErrIllegalState
	}

	m.status = MemberStatusActive
	return nil
}

func (m *Member) Deactivate() error {
	if m.status != MemberStatusActive {
		return ErrIllegalState
	}

	m.status = MemberStatusDeactivated
	return nil
}
