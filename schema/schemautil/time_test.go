package schemautil

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeMixin(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		t.Parallel()
		fields := CreateTimeMixin{}.Fields()
		require.Len(t, fields, 1)
		desc := fields[0].Descriptor()
		assert.Equal(t, "created_at", desc.Name)
		assert.True(t, desc.Immutable)
		assert.NotNil(t, desc.Default)
		assert.Nil(t, desc.UpdateDefault)
	})
	t.Run("Update", func(t *testing.T) {
		t.Parallel()
		fields := UpdateTimeMixin{}.Fields()
		require.Len(t, fields, 1)
		desc := fields[0].Descriptor()
		assert.Equal(t, "updated_at", desc.Name)
		assert.True(t, desc.Immutable)
		assert.NotNil(t, desc.Default)
		assert.NotNil(t, desc.UpdateDefault)
	})
	t.Run("Compose", func(t *testing.T) {
		t.Parallel()
		fields := TimeMixin{}.Fields()
		require.Len(t, fields, 2)
		assert.Equal(t, "created_at", fields[0].Descriptor().Name)
		assert.Equal(t, "updated_at", fields[1].Descriptor().Name)
	})
}
