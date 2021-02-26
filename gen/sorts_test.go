package gen

import (
	"go/format"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sf9v/nero/example"
)

func Test_newSortsFile(t *testing.T) {
	schema := (&example.User{}).Schema()
	require.NotNil(t, schema)

	buf, err := newSortsFile(schema)
	require.NoError(t, err)

	_, err = format.Source(buf.Bytes())
	require.NoError(t, err)

	_, err = newSortsFile(nil)
	assert.Error(t, err)
}
