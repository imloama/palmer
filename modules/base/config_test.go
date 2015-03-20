package base

import (
	"testing"
)

func Test_GetConfig(t *testing.T) {
	conf := GetConfig()
	t.Logf("conf:%v\n", conf)
}
