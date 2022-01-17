# deploymentgroup

This repository implements a simple controller for watching Foo resources as
defined with a CustomResourceDefinition (CRD).

This particular example demonstrates how to perform basic operations such as:

* How to register a new custom resource (custom resource type) of type `Foo` using a CustomResourceDefinition.
* How to create/get/list instances of your new resource type `Foo`.
* How to setup a controller on resource handling create/update/delete events.

It makes use of the generators in [k8s.io/code-generator](https://github.com/kubernetes/code-generator)
to generate a typed client, informers, listers and deep-copy functions. You can
do this yourself using the `./hack/update-codegen.sh` script.

The `update-codegen` script will automatically generate the following files &
directories:

* `pkg/apis/deploymentgroup/v1alpha1/zz_generated.deepcopy.go`
* `pkg/generated/`

Changes should not be made to these files manually, and when creating your own
controller based off of this implementation you should not copy these files and
instead run the `update-codegen` script to generate your own.

## Details

The sample controller uses [client-go library](https://github.com/kubernetes/client-go/tree/master/tools/cache) extensively.

Note that if you intend to
generate code then you will also need the
code-generator repo to exist in an old-style location.  One easy way
to do this is to use the command `go mod vendor` to create and
populate the `vendor` directory.

### A Note on kubernetes/kubernetes

## Running

**Prerequisite**: Since the sample-controller uses `apps/v1` deployments, the Kubernetes cluster version should be greater than 1.9.

## Subresources

Custom Resources support `/status` and `/scale` [subresources](https://kubernetes.io/docs/tasks/access-kubernetes-api/custom-resources/custom-resource-definitions/#subresources). The `CustomResourceSubresources` feature is in GA from v1.16.

## A Note on the API version
The [group](https://kubernetes.io/docs/reference/using-api/#api-groups) version of the custom resource in `crd.yaml` is `v1alpha`, this can be evolved to a stable API version, `v1`, using [CRD Versioning](https://kubernetes.io/docs/tasks/extend-kubernetes/custom-resources/custom-resource-definition-versioning/).

## Compatibility

HEAD of this repository will match HEAD of k8s.io/apimachinery and
k8s.io/client-go.
