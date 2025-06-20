package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMember(t *testing.T) {
	member := NewMember("kopher@goplearn.app", "Kopher", "secret")

	assert.Equal(t, member.status, MemberStatusPending)
}

func TestMember_Status(t *testing.T) {
	member := NewMember("kopher@goplearn.app", "Kopher", "secret")

	t.Run("activate", func(t *testing.T) {
		member.Activate()

		assert.Equal(t, member.status, MemberStatusActive)
	})

	t.Run("activate_fail", func(t *testing.T) {
		member.Activate()
		err := member.Activate()

		if ok := assert.Error(t, err); ok {
			assert.Equal(t, err, ErrIllegalState)
		}
	})

	t.Run("deactivate", func(t *testing.T) {
		member.Activate()
		member.Deactivate()

		assert.Equal(t, member.status, MemberStatusDeactivated)
	})

	t.Run("deactivate_fail", func(t *testing.T) {
		err := member.Deactivate()

		if ok := assert.Error(t, err); ok {
			assert.Equal(t, err, ErrIllegalState)
		}
	})

	t.Run("deactivate_fail_twice", func(t *testing.T) {
		member.Deactivate()
		err := member.Deactivate()

		if ok := assert.Error(t, err); ok {
			assert.Equal(t, err, ErrIllegalState)
		}
	})

}
