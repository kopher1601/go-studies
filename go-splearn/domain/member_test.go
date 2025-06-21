package domain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPasswordEncoder struct{}

func (m *mockPasswordEncoder) Encode(password string) (string, error) {
	return strings.ToUpper(password), nil
}

func (m *mockPasswordEncoder) Matches(password, encodedPassword string) bool {
	return strings.ToUpper(password) == encodedPassword
}

func createTestMember(t *testing.T) *Member {
	t.Helper()
	member, err := CreateMember("kopher@goplearn.app", "Kopher", "secret", &mockPasswordEncoder{})
	if err != nil {
		t.Fatal(err)
	}
	return member
}

func TestMember_Status(t *testing.T) {
	t.Run("activate", func(t *testing.T) {
		member := createTestMember(t)
		member.Activate()

		assert.Equal(t, member.status, MemberStatusActive)
	})

	t.Run("activate_fail", func(t *testing.T) {
		member := createTestMember(t)
		member.Activate()
		err := member.Activate()

		if ok := assert.Error(t, err); ok {
			assert.Equal(t, err, ErrIllegalState)
		}
	})

	t.Run("deactivate", func(t *testing.T) {
		member := createTestMember(t)
		member.Activate()
		member.Deactivate()

		assert.Equal(t, member.status, MemberStatusDeactivated)
	})

	t.Run("deactivate_fail", func(t *testing.T) {
		member := createTestMember(t)
		err := member.Deactivate()

		if ok := assert.Error(t, err); ok {
			assert.Equal(t, err, ErrIllegalState)
		}
	})

	t.Run("deactivate_fail_twice", func(t *testing.T) {
		member := createTestMember(t)
		member.Deactivate()
		err := member.Deactivate()

		if ok := assert.Error(t, err); ok {
			assert.Equal(t, err, ErrIllegalState)
		}
	})

}
