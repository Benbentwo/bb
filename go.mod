module github.ablevets.com/Digital-Transformation/av

require (
	cloud.google.com/go v0.38.0 // indirect
	github.com/gophercloud/gophercloud v0.1.0 // indirect
	github.com/jenkins-x/draft-repo v0.0.0-20180417100212-2f66cc518135
	github.com/jenkins-x/golang-jenkins v0.0.0-20180919102630-65b83ad42314
	github.com/jenkins-x/jx v0.0.0-20191001143610-9a26e9871d83
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.3.2
	google.golang.org/appengine v1.5.0 // indirect
	k8s.io/metrics v0.0.0-20190926001138-4e1cdcf4c305 // indirect
	k8s.io/utils v0.0.0-20190920012459-5008bf6f8cd6 // indirect
)

exclude github.com/jenkins-x/jx/pkg/prow v0.0.0-20190912224545-e8f82ee218ba

go 1.12
