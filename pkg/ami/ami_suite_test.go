package ami_test

import (
	"testing"

	"github.com/anodyne4j/eksctl/pkg/testutils"
)

func TestAmi(t *testing.T) {
	testutils.RegisterAndRun(t, "Ami Suite")
}
