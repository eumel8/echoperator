package v1alpha1

import (
	internalinterfaces "github.com/eumel8/echoperator/pkg/rds/v1alpha1/apis/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Rdss returns a RdsInformer.
	Rdss() RdsInformer
	// ScheduledRdss returns a ScheduledRdsInformer.
	ScheduledRdss() ScheduledRdsInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Rdss returns a RdsInformer.
func (v *version) Rdss() RdsInformer {
	return &echoInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ScheduledRdss returns a ScheduledRdsInformer.
func (v *version) ScheduledRdss() ScheduledRdsInformer {
	return &scheduledRdsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
