#!/bin/bash

# MIT License
#
# Copyright (c) 2020 Dmitrii Ustiugov, Shyam Jesalpura and EASE lab
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

set -Eeuo pipefail

cd "$( dirname "${BASH_SOURCE[0]}" )"

if (( $# != 1)); then
    echo "Invalid number of parameters"
    echo "USAGE: setup_cri_test_env.sh <sandbox>"
    exit 1
fi

SANDBOX=$1
VHIVE_ROOT="$(git rev-parse --show-toplevel)"

$VHIVE_ROOT/scripts/vhive_setup -vhive-repo-dir $VHIVE_ROOT start_onenode_vhive_cluster $SANDBOX
# $VHIVE_ROOT/scripts/cloudlab/start_onenode_vhive_cluster.sh "$SANDBOX"
sleep 30s

# KUBECONFIG=/etc/kubernetes/admin.conf sudo $VHIVE_ROOT/scripts/setup_zipkin.sh
$VHIVE_ROOT/scripts/vhive_setup -vhive-repo-dir $VHIVE_ROOT setup_zipkin

# FIXME (gh-709)
#source etc/profile && go run $VHIVE_ROOT/examples/registry/populate_registry.go -imageFile $VHIVE_ROOT/examples/registry/images.txt

KUBECONFIG=/etc/kubernetes/admin.conf bash -c "sudo kn service apply helloworld -f $VHIVE_ROOT/configs/knative_workloads/helloworld.yaml"
# FIXME (gh-709)
#KUBECONFIG=/etc/kubernetes/admin.conf sudo kn service apply helloworldlocal -f $VHIVE_ROOT/configs/knative_workloads/helloworld_local.yaml
#                                      ^^^^^^^ This WILL NOT work because ${KUBECONFIG} would not be set in the context when executing `kn`
#                                      VVVVVVV Use `bash -c` to fix the issue
KUBECONFIG=/etc/kubernetes/admin.conf bash -c "sudo kn service apply helloworldserial -f $VHIVE_ROOT/configs/knative_workloads/helloworldSerial.yaml"
KUBECONFIG=/etc/kubernetes/admin.conf bash -c "sudo kn service apply pyaes -f $VHIVE_ROOT/configs/knative_workloads/pyaes.yaml"
sleep 30s
