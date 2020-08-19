package operator

import (
	"os/exec"

	"github.com/Shanghai-Lunara/abstract-runner/proto"
	"k8s.io/klog"
)

type GitInterface interface {
}

type git struct {
	c proto.Git
}

func (g *git) Clone() {
	res, err := g.ExecuteWithArgs("pwd")
	if err != nil {
		klog.V(2).Info(err)
	}
	klog.Info("res:", string(res))
}

func (g *git) ExecuteWithArgs(cmd string, args ...string) (res []byte, err error) {
	out, err := exec.Command(cmd, args...).Output()
	//if err != nil {
	//	klog.V(2).Info(err)
	//}
	return out, nil
}
