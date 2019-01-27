/*
Copyright YEAR The Kubernetes Authors.

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

package clustercredential

import (
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	"k8s.io/client-go/tools/record"

	clusterregistryv1alpha1 "k8s.io/cluster-registry/pkg/apis/clusterregistry/v1alpha1"
	clusterregistryv1alpha1client "k8s.io/cluster-registry/pkg/client/clientset/versioned/typed/clusterregistry/v1alpha1"
	clusterregistryv1alpha1informer "k8s.io/cluster-registry/pkg/client/informers/externalversions/clusterregistry/v1alpha1"
	clusterregistryv1alpha1lister "k8s.io/cluster-registry/pkg/client/listers/clusterregistry/v1alpha1"

	"k8s.io/cluster-registry/pkg/inject/args"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for ClusterCredential resources goes here.

func (bc *ClusterCredentialController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	log.Printf("Implement the Reconcile function on clustercredential.ClusterCredentialController to reconcile %s\n", k.Name)
	return nil
}

// +kubebuilder:controller:group=clusterregistry,version=v1alpha1,kind=ClusterCredential,resource=clustercredentials
type ClusterCredentialController struct {
	// INSERT ADDITIONAL FIELDS HERE
	clustercredentialLister clusterregistryv1alpha1lister.ClusterCredentialLister
	clustercredentialclient clusterregistryv1alpha1client.ClusterregistryV1alpha1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	clustercredentialrecorder record.EventRecorder
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &ClusterCredentialController{
		clustercredentialLister: arguments.ControllerManager.GetInformerProvider(&clusterregistryv1alpha1.ClusterCredential{}).(clusterregistryv1alpha1informer.ClusterCredentialInformer).Lister(),

		clustercredentialclient:   arguments.Clientset.ClusterregistryV1alpha1(),
		clustercredentialrecorder: arguments.CreateRecorder("ClusterCredentialController"),
	}

	// Create a new controller that will call ClusterCredentialController.Reconcile on changes to ClusterCredentials
	gc := &controller.GenericController{
		Name:             "ClusterCredentialController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&clusterregistryv1alpha1.ClusterCredential{}); err != nil {
		return gc, err
	}

	// IMPORTANT:
	// To watch additional resource types - such as those created by your controller - add gc.Watch* function calls here
	// Watch function calls will transform each object event into a ClusterCredential Key to be reconciled by the controller.
	//
	// **********
	// For any new Watched types, you MUST add the appropriate // +kubebuilder:informer and // +kubebuilder:rbac
	// annotations to the ClusterCredentialController and run "kubebuilder generate.
	// This will generate the code to start the informers and create the RBAC rules needed for running in a cluster.
	// See:
	// https://godoc.org/github.com/kubernetes-sigs/kubebuilder/pkg/gen/controller#example-package
	// **********

	return gc, nil
}
