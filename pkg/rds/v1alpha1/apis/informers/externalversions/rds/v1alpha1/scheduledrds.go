package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/eumel8/echoperator/pkg/rds/v1alpha1/apis/clientset/versioned"
	internalinterfaces "github.com/eumel8/echoperator/pkg/rds/v1alpha1/apis/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/eumel8/echoperator/pkg/rds/v1alpha1/apis/listers/rds/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ScheduledRdsInformer provides access to a shared informer and lister for
// ScheduledRdss.
type ScheduledRdsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ScheduledRdsLister
}

type scheduledRdsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewScheduledRdsInformer constructs a new informer for ScheduledRds type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewScheduledRdsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredScheduledRdsInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredScheduledRdsInformer constructs a new informer for ScheduledRds type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredScheduledRdsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.McspsV1alpha1().ScheduledRdss(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.McspsV1alpha1().ScheduledRdss(namespace).Watch(context.TODO(), options)
			},
		},
		&echov1alpha1.ScheduledRds{},
		resyncPeriod,
		indexers,
	)
}

func (f *scheduledRdsInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredScheduledRdsInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *scheduledRdsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&echov1alpha1.ScheduledRds{}, f.defaultInformer)
}

func (f *scheduledRdsInformer) Lister() v1alpha1.ScheduledRdsLister {
	return v1alpha1.NewScheduledRdsLister(f.Informer().GetIndexer())
}
