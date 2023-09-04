package customize

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func Test_loadConfigJSON(t *testing.T) {
	t.Parallel()

	conf, err := loadConfigJSON("testdata/config.json")

	assert.NoError(t, err)

	assert.NotNil(t, gjson.GetBytes(conf, "tabs.custom"))

	_, err = loadConfigJSON("testdata/config-bad.json")

	assert.Error(t, err)

	_, err = loadConfigJSON("testdata/config-not-exists.json")

	assert.Error(t, err)
}

func TestCustomize(t *testing.T) {
	t.Parallel()

	conf, err := Customize(testconfig)

	assert.NoError(t, err)

	assert.False(t, gjson.GetBytes(conf, `tabs.#(id="custom")`).Exists())
}

func TestCustomize_env_found(t *testing.T) { //nolint:paralleltest
	t.Setenv("XK6_DASHBOARD_CONFIG", "testdata/config-custom.js")

	conf, err := Customize(testconfig)

	assert.NoError(t, err)

	assert.True(t, gjson.GetBytes(conf, `tabs.#(id="custom")`).Exists())

	t.Setenv("XK6_DASHBOARD_CONFIG", "testdata/config.json")

	assert.NoError(t, err)
}
