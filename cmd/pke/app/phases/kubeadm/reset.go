// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kubeadm

import (
	"fmt"
	"io"

	"github.com/banzaicloud/pke/cmd/pke/app/util/cri"
	"github.com/banzaicloud/pke/cmd/pke/app/util/runner"
)

const (
	cmdKubeadm = "kubeadm"
)

func Reset(out io.Writer, containerRuntime string) error {
	// kubeadm reset --force
	_, _ = fmt.Fprintln(out, "")
	_, _ = fmt.Fprintln(out, "================================================================================")
	_, _ = fmt.Fprintln(out, "Resetting kubeadm changes...")
	_, err := runner.Cmd(out, cmdKubeadm, "reset", "--force", fmt.Sprintf("--cri-socket=%s", cri.GetCRISocket(containerRuntime))).CombinedOutputAsync()
	if err != nil {
		return err
	}
	return nil
}
