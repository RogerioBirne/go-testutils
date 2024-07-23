package testutils

import (
	"context"
	"fmt"
	testutils "github.com/RogerioBirne/go-testutils/testutils/internal"
	"github.com/RogerioBirne/go-testutils/testutils/internal/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewArgumentCaptor(t *testing.T) {
	t.Run("Given Valid EntityID When ArgumentCaptor Then Return Entity", func(t *testing.T) {
		entityID := "123456789"
		target := NewArgumentCaptor[testutils.Entity]() // Create a new ArgumentCaptor

		rep := mocks.NewRepository(t)
		rep.EXPECT().
			Save(context.TODO(), entityID, target.Capture()). // Capture the argument on stubby
			Return(nil)

		errResult := func(id string, rep testutils.Repository) error {
			return rep.Save(context.TODO(), id, testutils.Entity{
				ID:   entityID,
				Name: fmt.Sprintf("name_%s", entityID),
				Time: time.Now().Format(time.RFC3339), // Dynamic value
			})
		}(entityID, rep)

		assert.Nil(t, errResult)

		capturedEntity := target.GetValue() // Get the captured argument
		assert.Equal(t, entityID, capturedEntity.ID)
		assert.Equal(t, fmt.Sprintf("name_%s", entityID), capturedEntity.Name)

		// Test dynamic value
		_, err := time.Parse(time.RFC3339, capturedEntity.Time)
		assert.Nil(t, err)
	})

	t.Run("Given a Valid EntityID When func inject a nil valid Then Return error", func(t *testing.T) {
		entityID := "123456789"
		target := NewArgumentCaptor[testutils.Entity]()
		rep := &mocks.Repository{}
		mockCall := rep.EXPECT().
			Save(context.TODO(), entityID, target.Capture()).
			Return(nil)

		assert.Panics(t, func() {
			_ = func(id string, rep testutils.Repository) error {
				return rep.Save(context.TODO(), id, nil)
			}(entityID, rep)
		})

		rep.AssertNotCalled(t, mockCall.Method, mockCall.Arguments...)
	})

	t.Run("Given a Valid EntityID When func inject a not Entity valid Then Return error", func(t *testing.T) {
		entityID := "123456789"
		target := NewArgumentCaptor[testutils.Entity]()

		rep := &mocks.Repository{}
		mockCall := rep.EXPECT().
			Save(context.TODO(), entityID, target.Capture()).
			Return(nil)

		assert.Panics(t, func() {
			_ = func(id string, rep testutils.Repository) error {
				return rep.Save(context.TODO(), id, "not entity")
			}(entityID, rep)
		})

		rep.AssertNotCalled(t, mockCall.Method, mockCall.Arguments...)
	})
}
