package mt5tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testGetTotalGroups(t *testing.T) {
	total, err := mt.GetTotalGroups()
	require.NoError(t, err)
	require.NotEmpty(t, total)
}

func testGetGroupByName(t *testing.T) {
	group, err := mt.GetGroupByName("demo\\group")
	require.NoError(t, err)
	require.NotEmpty(t, group)
}

func testGetAllGroups(t *testing.T) {
	groups, err := mt.GetAllGroups()
	require.NoError(t, err)
	require.NotEmpty(t, groups)
}
