package kubeconfig

import (
	"testing"

	"github.com/anodyne4j/eksctl/pkg/testutils"
)

func TestKubeConfig(t *testing.T) {
	testutils.RegisterAndRun(t, "KubeConfig Suite")
}
