#!/usr/bin/env bash

# This is the hardcoded result of running ./hack/ginkgo-e2e.sh with a cluster that
# was created using kube-up.sh

# bazel build //cmd/kubectl //test/e2e:e2e.test //vendor/github.com/onsi/ginkgo/ginkgo

/Users/spiffxp/w/go/src/k8s.io/kubernetes/bazel-bin/vendor/github.com/onsi/ginkgo/ginkgo/darwin_amd64_stripped/ginkgo \
  /Users/spiffxp/w/go/src/k8s.io/kubernetes/bazel-bin/test/e2e/e2e.test -- \
  --kubeconfig=/Users/spiffxp/.kube/config \
  --ginkgo.flakeAttempts=1 \
  --host=https://104.197.247.226 \
  --provider=gce \
  --gce-project=spiffxp-gke-dev \
  --gce-zone=us-central1-b \
  --gce-region=us-central1 \
  --gce-multizone=false \
  --gke-cluster=kubernetes \
  --kube-master=kubernetes-master \
  --cluster-tag= \
  --cloud-config-file= \
  --repo-root=/Users/spiffxp/w/go/src/k8s.io/kubernetes \
  --node-instance-group=kubernetes-minion-group \
  --prefix=e2e \
  --network=e2e \
  --node-tag=kubernetes-minion \
  --master-tag=kubernetes-master \
  --docker-config-file= \
  --dns-domain=cluster.local \
  --ginkgo.slowSpecThreshold=300 \
  --master-os-distro=gci \
  --node-os-distro=gci \
  --num-nodes=3 \
  --report-dir=/Users/spiffxp/w/go/src/k8s.io/kubernetes/artifacts \
  "${@:-}"
