// MIT License
//
// Copyright (c) 2023 Haoyuan Ma and vHive team
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cluster

import (
	"path"

	configs "github.com/vhive-serverless/vHive/scripts/configs"
	utils "github.com/vhive-serverless/vHive/scripts/utils"
)

func SetupMasterNode(stockContainerd string) error {
	// Original Bash Scripts: scripts/cluster/setup_master_node.sh

	err := InstallCalico()
	if err != nil {
		return err
	}

	err = InstallMetalLB()
	if err != nil {
		return err
	}

	err = InstallIstio()
	if err != nil {
		return err
	}

	err = InstallKnativeServingComponent(stockContainerd)
	if err != nil {
		return err
	}

	err = InstallLocalClusterRegistry()
	if err != nil {
		return err
	}

	err = ConfigureMagicDNS()
	if err != nil {
		return err
	}

	err = DeployIstioPods()
	if err != nil {
		return err
	}

	// Logs for verification
	_, err = utils.ExecShellCmd("kubectl get pods -n knative-serving")
	if !utils.CheckErrorWithMsg(err, "Verification Failed!\n") {
		return err
	}

	err = InstallKnativeEventingComponent()
	if err != nil {
		return err
	}

	// Logs for verification
	_, err = utils.ExecShellCmd("kubectl get pods -n knative-eventing")
	if !utils.CheckErrorWithMsg(err, "Verification Failed!") {
		return err
	}

	err = InstallChannelLayer()
	if err != nil {
		return err
	}

	err = InstallBrokerLayer()
	if err != nil {
		return err
	}

	// Logs for verification
	_, err = utils.ExecShellCmd("kubectl --namespace istio-system get service istio-ingressgateway")
	if !utils.CheckErrorWithMsg(err, "Verification Failed!") {
		return err
	}

	return nil
}

// Install Calico network add-on
func InstallCalico() error {

	utils.WaitPrintf("Installing pod network")
	_, err := utils.ExecShellCmd("kubectl apply -f %s", configs.Kube.PodNetworkAddonConfigURL)
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to install pod network!\n") {
		return err
	}
	return nil
}

// Install and configure MetalLB
func InstallMetalLB() error {
	utils.WaitPrintf("Installing and configuring MetalLB")
	_, err := utils.ExecShellCmd(`kubectl get configmap kube-proxy -n kube-system -o yaml | sed -e "s/strictARP: false/strictARP: true/" | kubectl apply -f - -n kube-system`)
	if !utils.CheckErrorWithMsg(err, "Failed to install and configure MetalLB!\n") {
		return err
	}
	_, err = utils.ExecShellCmd("kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v%s/config/manifests/metallb-native.yaml", configs.Knative.MetalLBVersion)
	if !utils.CheckErrorWithMsg(err, "Failed to install and configure MetalLB!\n") {
		return err
	}
	_, err = utils.ExecShellCmd("kubectl -n metallb-system wait deploy controller --timeout=90s --for=condition=Available")
	if !utils.CheckErrorWithMsg(err, "Failed to install and configure MetalLB!\n") {
		return err
	}

	metalibConfigsDir := "configs/metallb"
	metalibConfigsList := []string{
		"metallb-ipaddresspool.yaml",
		"metallb-l2advertisement.yaml",
	}

	for _, configFile := range metalibConfigsList {
		metalibConfigPath, err := utils.GetVHiveFilePath(path.Join(metalibConfigsDir, configFile))
		if err != nil {
			return err
		}
		_, err = utils.ExecShellCmd("kubectl apply -f %s", metalibConfigPath)
		if !utils.CheckErrorWithMsg(err, "Failed to install and configure MetalLB!\n") {
			return err
		}
	}
	utils.SuccessPrintf("\n")
	return nil
}

// Install istio
func InstallIstio() error {
	// Install istio
	// Download istio
	utils.WaitPrintf("Downloading istio")
	istioFilePath, err := utils.DownloadToTmpDir(configs.Knative.GetIstioDownloadUrl())
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to download istio!\n") {
		return err
	}
	// Extract istio
	utils.WaitPrintf("Extracting istio")
	err = utils.ExtractToDir(istioFilePath, "/usr/local", true)
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to extract istio!\n") {
		return err
	}

	// Grant permissions for other users to use
	_, err = utils.ExecShellCmd("sudo chmod -R o+x /usr/local/istio-%s/bin/istioctl", configs.Knative.IstioVersion)
	if !utils.CheckErrorWithMsg(err, "Failed to grant permissions to istioctl!\n") {
		return err
	}

	// Update PATH
	err = utils.AppendDirToPath("/usr/local/istio-%s/bin", configs.Knative.IstioVersion)
	if !utils.CheckErrorWithMsg(err, "Failed to update PATH!\n") {
		return err
	}
	// Deploy istio operator
	utils.WaitPrintf("Deploying istio operator")
	operatorConfigPath, err := utils.GetVHiveFilePath(configs.Knative.IstioOperatorConfigPath)
	if !utils.CheckErrorWithMsg(err, "Failed to find istio operator config!\n") {
		return err
	}
	_, err = utils.ExecShellCmd("/usr/local/istio-%s/bin/istioctl install -y -f %s", configs.Knative.IstioVersion, operatorConfigPath)
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to deploy istio operator!\n") {
		return err
	}

	return nil
}

// Install Knative Serving component
func InstallKnativeServingComponent(stockContainerd string) error {
	utils.WaitPrintf("Installing Knative Serving component (%s mode)", stockContainerd)
	if stockContainerd == "stock-only" {
		_, err := utils.ExecShellCmd("kubectl apply -f https://github.com/knative/serving/releases/download/knative-v%s/serving-crds.yaml", configs.Knative.KnativeVersion)
		if !utils.CheckErrorWithMsg(err, "Failed to install Knative Serving component!\n") {
			return err
		}
		_, err = utils.ExecShellCmd("kubectl apply -f https://github.com/knative/serving/releases/download/knative-v%s/serving-core.yaml", configs.Knative.KnativeVersion)
		if !utils.CheckErrorWithTagAndMsg(err, "Failed to install Knative Serving component!\n") {
			return err
		}
	} else {
		_, err := utils.ExecShellCmd("kubectl apply -f %s/serving-crds.yaml", configs.Knative.NotStockOnlyKnativeServingYamlUrlPrefix)
		if !utils.CheckErrorWithMsg(err, "Failed to install Knative Serving component!\n") {
			return err
		}
		_, err = utils.ExecShellCmd("kubectl apply -f %s/serving-core.yaml", configs.Knative.NotStockOnlyKnativeServingYamlUrlPrefix)
		if !utils.CheckErrorWithTagAndMsg(err, "Failed to install Knative Serving component!\n") {
			return err
		}
	}
	return nil
}

// Install local cluster registry
func InstallLocalClusterRegistry() error {
	utils.WaitPrintf("Installing local cluster registry")
	_, err := utils.ExecShellCmd("kubectl create namespace registry")
	if !utils.CheckErrorWithMsg(err, "Failed to install local cluster registry!\n") {
		return err
	}
	configFilePath, err := utils.GetVHiveFilePath(configs.Knative.LocalRegistryVolumeConfigPath)
	if !utils.CheckErrorWithMsg(err, "Failed to find local cluster registry config!\n") {
		return err
	}
	_, err = utils.ExecShellCmd("REPO_VOL_SIZE=%s envsubst < %s | kubectl create --filename -", configs.Knative.LocalRegistryRepoVolumeSize, configFilePath)
	if !utils.CheckErrorWithMsg(err, "Failed to install local cluster registry!\n") {
		return err
	}
	dockerRegistryConfigPath, err := utils.GetVHiveFilePath(configs.Knative.LocalRegistryDockerRegistryConfigPath)
	if !utils.CheckErrorWithMsg(err, "Failed to find local cluster registry config!\n") {
		return err
	}
	hostUpdateConfigPath, err := utils.GetVHiveFilePath(configs.Knative.LocalRegistryHostUpdateConfigPath)
	if !utils.CheckErrorWithMsg(err, "Failed to find local cluster registry config!\n") {
		return err
	}
	_, err = utils.ExecShellCmd("kubectl create -f %s && kubectl apply -f %s", dockerRegistryConfigPath, hostUpdateConfigPath)
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to install local cluster registry!\n") {
		return err
	}
	return nil
}

// Configure Magic DNS
func ConfigureMagicDNS() error {
	utils.WaitPrintf("Configuring Magic DNS")
	magicDNSConfigPath, err := utils.GetVHiveFilePath(configs.Knative.MagicDNSConfigPath)
	if !utils.CheckErrorWithMsg(err, "Failed to find Magic DNS config!\n") {
		return err
	}
	_, err = utils.ExecShellCmd("kubectl apply -f %s", magicDNSConfigPath)
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to configure Magic DNS!\n") {
		return err
	}
	return nil
}

// Deploy Istio pods
func DeployIstioPods() error {
	utils.WaitPrintf("Deploying istio pods")
	_, err := utils.ExecShellCmd("kubectl apply -f https://github.com/knative/net-istio/releases/download/knative-v%s/net-istio.yaml", configs.Knative.KnativeVersion)
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to deploy istio pods!\n") {
		return err
	}
	return nil
}

// Install Knative Eventing component
func InstallKnativeEventingComponent() error {
	utils.WaitPrintf("Installing Knative Eventing component")
	_, err := utils.ExecShellCmd("kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v%s/eventing-crds.yaml", configs.Knative.KnativeVersion)
	if !utils.CheckErrorWithMsg(err, "Failed to install Knative Eventing component!\n") {
		return err
	}
	_, err = utils.ExecShellCmd("kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v%s/eventing-core.yaml", configs.Knative.KnativeVersion)
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to install Knative Eventing component!\n") {
		return err
	}
	return nil
}

// Install a default Channel (messaging) layer
func InstallChannelLayer() error {
	utils.WaitPrintf("Installing a default Channel (messaging) layer")
	_, err := utils.ExecShellCmd("kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v%s/in-memory-channel.yaml", configs.Knative.KnativeVersion)
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to install a default Channel (messaging) layer!\n") {
		return err
	}
	return nil
}

// Install a Broker layer
func InstallBrokerLayer() error {
	utils.WaitPrintf("Installing a Broker layer")
	_, err := utils.ExecShellCmd("kubectl apply -f https://github.com/knative/eventing/releases/download/knative-v%s/mt-channel-broker.yaml", configs.Knative.KnativeVersion)
	if !utils.CheckErrorWithTagAndMsg(err, "Failed to install a Broker layer!\n") {
		return err
	}
	return nil
}
