/*
Copyright 2022 The KubeVela Authors.

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

package v1alpha1

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/kubevela/prism/pkg/util/singleton"
)

var testEnv *envtest.Environment

func TestApplicationResourceTracker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ApplicationResourceTracker Extension API Test")
}

var _ = BeforeSuite(func() {
	logf.SetLogger(zap.New(zap.UseDevMode(true), zap.WriteTo(GinkgoWriter)))
	By("Bootstrapping Test Environment")
	testEnv = &envtest.Environment{
		ControlPlaneStartTimeout: time.Minute,
		ControlPlaneStopTimeout:  time.Minute,
		Scheme:                   scheme.Scheme,
		CRDDirectoryPaths:        []string{"../../../../test/testdata/crds"},
		UseExistingCluster:       pointer.Bool(false),
	}

	cfg, err := testEnv.Start()
	Ω(err).To(Succeed())
	singleton.SetKubeConfig(cfg)
	k8sClient, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	Ω(err).To(Succeed())
	singleton.SetKubeClient(k8sClient)
})

var _ = AfterSuite(func() {
	By("Tearing Down the Test Environment")
	Ω(testEnv.Stop()).To(Succeed())
})
