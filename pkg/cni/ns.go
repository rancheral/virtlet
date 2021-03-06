/*
Copyright 2016 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cni

import (
	"os/exec"
	"path"

	"github.com/golang/glog"
)

// we are using "ip netns" command instead of cni/pkg/ns because usage
// of later will lead to file descriptor leakages
// also we could be affected by https://github.com/containernetworking/cni/issues/262

const (
	netnsBasePath = "/var/run/netns"
)

func CreateNetNS(name string) error {
	return callIpNetns("add", name)
}

func DestroyNetNS(name string) error {
	return callIpNetns("del", name)
}

func callIpNetns(command, name string) error {
	cmd := exec.Command("ip", "netns", command, name)
	output, err := cmd.CombinedOutput()
	if err != nil {
		glog.Errorf("\"ip netns %s %s\" failed with output: %s", command, name, string(output))
	}
	return err
}

func PodNetNSPath(name string) string {
	return path.Join(netnsBasePath, name)
}
