package testutils

import (
	"context"
	"fmt"
	testutils "github.com/RogerioBirne/go-testutils/testutils/internal"
	"github.com/RogerioBirne/go-testutils/testutils/internal/mocks"
	"github.com/go-faker/faker/v4"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewArgumentInjector(t *testing.T) {
	t.Run("Given Valid EntityID When ArgumentInjector Then Return Entity", func(t *testing.T) {
		entityID := faker.Word()

		// Value to be injected
		e := testutils.Entity{
			ID:   entityID,
			Name: fmt.Sprintf("name_%s", entityID),
		}

		rep := mocks.NewRepository(t)
		rep.EXPECT().
			GetByID(context.TODO(), entityID, NewArgumentInjector(e)). // Inject the value on stubby
			Return(nil)

		result, errResult := func(id string, repository testutils.Repository) (e *testutils.Entity, err error) {
			err = repository.GetByID(context.TODO(), id, &e) // Call the target function
			if err != nil {
				return nil, err
			}
			return e, nil
		}(entityID, rep)

		assert.NotNil(t, result)
		assert.Equal(t, *result, e) // Check if the result is the same as the injected value
		assert.Nil(t, errResult)
	})
}
