// This is a submodule to isolate github.com/ducesoft/code-generator from k8s.io/{api,apimachinery,client-go} dependencies in generated code

module github.com/ducesoft/code-generator/examples

go 1.16

require (
	github.com/go-openapi/spec v0.19.3
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/kube-openapi v0.0.0-20210305001622-591a79e4bda7
)

replace (
	k8s.io/api => ../../api
	k8s.io/apimachinery => ../../apimachinery
	k8s.io/client-go => ../../client-go
)
