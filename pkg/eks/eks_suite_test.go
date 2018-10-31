package eks_test

import (
	"testing"

	"github.com/anodyne4j/eksctl/pkg/testutils"
)

func TestEks(t *testing.T) {
	testutils.RegisterAndRun(t, "Eks Suite")
}
