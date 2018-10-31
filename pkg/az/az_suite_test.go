package az_test

import (
	"testing"

	"github.com/anodyne4j/eksctl/pkg/testutils"
)

func TestAZ(t *testing.T) {
	testutils.RegisterAndRun(t, "AZ Suite")
}
