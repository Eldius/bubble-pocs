package purpur

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseGetMineVersionsResponse(t *testing.T) {
	res, err := parseGetMineVersionsResponse([]byte(`{
    "project": "purpur",
    "metadata": {
        "current": "1.21.1"
    },
    "versions": [
        "1.14.1",
        "1.14.2",
        "1.14.3",
        "1.14.4",
        "1.15",
        "1.15.1",
        "1.15.2",
        "1.16.1",
        "1.16.2",
        "1.16.3",
        "1.16.4",
        "1.16.5",
        "1.17",
        "1.17.1",
        "1.18",
        "1.18.1",
        "1.18.2",
        "1.19",
        "1.19.1",
        "1.19.2",
        "1.19.3",
        "1.19.4",
        "1.20",
        "1.20.1",
        "1.20.2",
        "1.20.4",
        "1.20.6",
        "1.21",
        "1.21.1",
        "1.21.3"
    ]
}`), 200)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "purpur", res.Project)
	assert.Len(t, res.Versions, 30)
}
