module github.ablevets.com/Digital-Transformation/av

require (
	cloud.google.com/go v0.38.0 // indirect
	github.com/Pallinder/go-randomdata v0.0.0-20180616180521-15df0648130a
	github.com/alexflint/go-filemutex v0.0.0-20171028004239-d358565f3c3f
	github.com/banzaicloud/bank-vaults v0.0.0-20190508130850-5673d28c46bd
	github.com/blang/semver v3.5.1+incompatible
	github.com/fatih/color v1.7.0
	github.com/gophercloud/gophercloud v0.1.0 // indirect
	github.com/heptio/sonobuoy v0.16.0
	github.com/jenkins-x/draft-repo v0.0.0-20180417100212-2f66cc518135
	github.com/jenkins-x/golang-jenkins v0.0.0-20180919102630-65b83ad42314
	github.com/jenkins-x/jx v0.0.0-20191001143610-9a26e9871d83
	github.com/jetstack/cert-manager v0.5.2
	github.com/knative/build v0.5.0
	github.com/knative/pkg v0.0.0-20191001225505-346b0abf16cd // indirect
	github.com/knative/serving v0.5.0
	github.com/pborman/uuid v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/rickar/props v0.0.0-20170718221555-0b06aeb2f037
	github.com/sirupsen/logrus v1.4.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.2
	github.com/tektoncd/pipeline v0.5.1
	gocloud.dev v0.9.0
	google.golang.org/appengine v1.5.0 // indirect
	gopkg.in/AlecAivazis/survey.v1 v1.8.3
	gopkg.in/src-d/go-git.v4 v4.5.0
	k8s.io/api v0.0.0-20190925180651-d58b53da08f5
	k8s.io/apiextensions-apiserver v0.0.0-20190308081736-3a66ae4d2f93
	k8s.io/apimachinery v0.0.0-20190925235427-62598f38f24e
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	// k8s.io/client-go v11.0.0+incompatible
	k8s.io/metrics v0.0.0-20190926001138-4e1cdcf4c305 // indirect
	k8s.io/utils v0.0.0-20190920012459-5008bf6f8cd6 // indirect
	sigs.k8s.io/yaml v1.1.0
)

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
