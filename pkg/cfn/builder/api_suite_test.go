package builder_test

import (
	"testing"

	"github.com/anodyne4j/eksctl/pkg/testutils"
)

func TestCFNBuilder(t *testing.T) {
	testutils.RegisterAndRun(t, "API Suite")
}
