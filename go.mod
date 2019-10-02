module github.ablevets.com/Digital-Transformation/av

exclude github.com/jenkins-x/jx/pkg/prow v0.0.0-20190912224545-e8f82ee218ba

exclude knative.dev/pkg v0.0.0-20191001225505-346b0abf16cd

replace k8s.io/api => k8s.io/api v0.0.0-20181128191700-6db15a15d2d3

replace k8s.io/metrics => k8s.io/metrics v0.0.0-20181128195641-3954d62a524d

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190122181752-bebe27e40fb7

replace k8s.io/client-go => k8s.io/client-go v2.0.0-alpha.0.0.20190115164855-701b91367003+incompatible

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20181128195303-1f84094d7e8e

replace github.com/banzaicloud/bank-vaults => github.com/banzaicloud/bank-vaults v0.0.0-20190508130850-5673d28c46bd

replace github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v21.1.0+incompatible

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v10.15.5+incompatible

go 1.12

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0 // indirect
	github.com/Netflix/go-expect v0.0.0-20190729225929-0e00d9168667
	github.com/Pallinder/go-randomdata v1.2.0
	github.com/alexflint/go-filemutex v0.0.0-20171028004239-d358565f3c3f
	github.com/blang/semver v3.5.1+incompatible
	github.com/briandowns/spinner v1.7.0 // indirect
	github.com/c2h5oh/datasize v0.0.0-20171227191756-4eba002a5eae // indirect
	github.com/docker/spdystream v0.0.0-20181023171402-6480d4af844c // indirect
	github.com/fatih/color v1.7.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/heptio/sonobuoy v0.16.1
	github.com/jenkins-x/draft-repo v0.9.0
	github.com/jenkins-x/golang-jenkins v0.0.0-20180919102630-65b83ad42314
	github.com/jenkins-x/jx v1.3.1119
	github.com/jetstack/cert-manager v0.10.1
	github.com/knative/build v0.7.0
	github.com/knative/serving v0.9.0
	github.com/kylelemons/godebug v1.1.0 // indirect
	github.com/pborman/uuid v1.2.0
	github.com/petergtz/pegomock v2.6.0+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/rickar/props v0.0.0-20170718221555-0b06aeb2f037
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.4.0
	github.com/tektoncd/pipeline v0.7.0
	github.com/viniciuschiele/tarx v0.0.0-20151205142357-6e3da540444d // indirect
	gocloud.dev v0.17.0
	gopkg.in/AlecAivazis/survey.v1 v1.8.7
	gopkg.in/src-d/go-git.v4 v4.13.1
	k8s.io/api v0.0.0-20190718183219-b59d8169aab5
	k8s.io/apiextensions-apiserver v0.0.0-20190718185103-d1ef975d28ce
	k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	knative.dev/pkg v0.0.0-20191002055904-849fcc967b59 // indirect
	knative.dev/serving v0.9.0 // indirect
	sigs.k8s.io/yaml v1.1.0
)
