package k8s

import (
	"fmt"

	apps_v1 "k8s.io/api/apps/v1"
	core_v1 "k8s.io/api/core/v1"
)

// GenericResource - generic resource,
// used to work with multiple kinds of k8s resources
type GenericResource struct {
	// original resource
	obj interface{}

	Identifier string
	Namespace  string
	Name       string
}

type genericResource []*GenericResource

func (c genericResource) Len() int {
	return len(c)
}

func (c genericResource) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c genericResource) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}

// NewGenericResource - create new generic k8s resource
func NewGenericResource(obj interface{}) (*GenericResource, error) {

	switch obj.(type) {
	case *apps_v1.Deployment, *apps_v1.StatefulSet, *apps_v1.DaemonSet:
		// ok
	default:
		return nil, fmt.Errorf("unsupported resource")
	}

	gr := &GenericResource{
		obj: obj,
	}

	gr.Identifier = gr.GetIdentifier()
	gr.Namespace = gr.GetNamespace()
	gr.Name = gr.GetName()

	return gr, nil
}

// GetIdentifier returns resource identifier
func (r *GenericResource) GetIdentifier() string {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getDeploymentIdentifier(obj)
	case *apps_v1.StatefulSet:
		return getStatefulSetIdentifier(obj)
	case *apps_v1.DaemonSet:
		return getDaemonsetSetIdentifier(obj)
	}
	return ""
}

// GetName returns resource name
func (r *GenericResource) GetName() string {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return obj.GetName()
	case *apps_v1.StatefulSet:
		return obj.GetName()
	case *apps_v1.DaemonSet:
		return obj.GetName()
	}
	return ""
}

// GetNamespace returns resource namespace
func (r *GenericResource) GetNamespace() string {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return obj.GetNamespace()
	case *apps_v1.StatefulSet:
		return obj.GetNamespace()
	case *apps_v1.DaemonSet:
		return obj.GetNamespace()
	}
	return ""
}

// Kind returns a type of resource that this structure represents
func (r *GenericResource) Kind() string {
	switch r.obj.(type) {
	case *apps_v1.Deployment:
		return "deployment"
	case *apps_v1.StatefulSet:
		return "statefulset"
	case *apps_v1.DaemonSet:
		return "daemonset"
	}
	return ""
}

// GetResource - get resource
func (r *GenericResource) GetResource() interface{} {
	return r.obj
}

// GetLabels - get resource labels
func (r *GenericResource) GetLabels() (labels map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return obj.GetLabels()
	case *apps_v1.StatefulSet:
		return obj.GetLabels()
	case *apps_v1.DaemonSet:
		return obj.GetLabels()
	}
	return
}

// SetLabels - set resource labels
func (r *GenericResource) SetLabels(labels map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		obj.SetLabels(labels)
	case *apps_v1.StatefulSet:
		obj.SetLabels(labels)
	case *apps_v1.DaemonSet:
		obj.SetLabels(labels)
	}
	return
}

// GetAnnotations - get resource annotations
func (r *GenericResource) GetAnnotations() (annotations map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return obj.GetAnnotations()
	case *apps_v1.StatefulSet:
		return obj.GetAnnotations()
	case *apps_v1.DaemonSet:
		return obj.GetAnnotations()
	}
	return
}

// SetAnnotations - set resource annotations
func (r *GenericResource) SetAnnotations(annotations map[string]string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		obj.SetAnnotations(annotations)
	case *apps_v1.StatefulSet:
		obj.SetAnnotations(annotations)
	case *apps_v1.DaemonSet:
		obj.SetAnnotations(annotations)
	}
	return
}

// GetImagePullSecrets - returns secrets from pod spec
func (r *GenericResource) GetImagePullSecrets() (secrets []string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getImagePullSecrets(obj.Spec.Template.Spec.ImagePullSecrets)
	case *apps_v1.StatefulSet:
		return getImagePullSecrets(obj.Spec.Template.Spec.ImagePullSecrets)
	case *apps_v1.DaemonSet:
		return getImagePullSecrets(obj.Spec.Template.Spec.ImagePullSecrets)
	}
	return
}

// GetImages - returns images used by this resource
func (r *GenericResource) GetImages() (images []string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return getContainerImages(obj.Spec.Template.Spec.Containers)
	case *apps_v1.StatefulSet:
		return getContainerImages(obj.Spec.Template.Spec.Containers)
	case *apps_v1.DaemonSet:
		return getContainerImages(obj.Spec.Template.Spec.Containers)
	}
	return
}

// Containers - returns containers managed by this resource
func (r *GenericResource) Containers() (containers []core_v1.Container) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		return obj.Spec.Template.Spec.Containers
	case *apps_v1.StatefulSet:
		return obj.Spec.Template.Spec.Containers
	case *apps_v1.DaemonSet:
		return obj.Spec.Template.Spec.Containers
	}
	return
}

// UpdateContainer - updates container image
func (r *GenericResource) UpdateContainer(index int, image string) {
	switch obj := r.obj.(type) {
	case *apps_v1.Deployment:
		updateDeploymentContainer(obj, index, image)
	case *apps_v1.StatefulSet:
		updateStatefulSetContainer(obj, index, image)
	case *apps_v1.DaemonSet:
		updateDaemonsetSetContainer(obj, index, image)
	}
	return
}