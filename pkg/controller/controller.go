package controller

import (
	"context"
	"errors"
	"time"

	"github.com/gotway/gotway/pkg/log"

	rdsv1alpha1 "github.com/eumel8/echoperator/pkg/rds/v1alpha1"
	rdsv1alpha1clientset "github.com/eumel8/echoperator/pkg/rds/v1alpha1/apis/clientset/versioned"
	rdsinformers "github.com/eumel8/echoperator/pkg/rds/v1alpha1/apis/informers/externalversions"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	kubeClientSet kubernetes.Interface

	rdsInformer           cache.SharedIndexInformer
	jobInformer           cache.SharedIndexInformer
	scheduledEchoInformer cache.SharedIndexInformer
	cronjobInformer       cache.SharedIndexInformer

	queue workqueue.RateLimitingInterface

	namespace string

	logger log.Logger
}

func (c *Controller) Run(ctx context.Context, numWorkers int) error {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	c.logger.Info("starting controller")

	c.logger.Info("starting informers")
	for _, i := range []cache.SharedIndexInformer{
		c.rdsInformer,
		c.scheduledEchoInformer,
		c.jobInformer,
		c.cronjobInformer,
	} {
		go i.Run(ctx.Done())
	}

	c.logger.Info("waiting for informer caches to sync")
	if !cache.WaitForCacheSync(ctx.Done(), []cache.InformerSynced{
		c.rdsInformer.HasSynced,
		c.scheduledEchoInformer.HasSynced,
		c.jobInformer.HasSynced,
		c.cronjobInformer.HasSynced,
	}...) {
		err := errors.New("failed to wait for informers caches to sync")
		utilruntime.HandleError(err)
		return err
	}

	c.logger.Infof("starting %d workers", numWorkers)
	for i := 0; i < numWorkers; i++ {
		go wait.Until(func() {
			c.runWorker(ctx)
		}, time.Second, ctx.Done())
	}
	c.logger.Info("controller ready")

	<-ctx.Done()
	c.logger.Info("stopping controller")

	return nil
}

func (c *Controller) addEcho(obj interface{}) {
	c.logger.Debug("adding rds")
	rds, ok := obj.(*rdsv1alpha1.Rds)
	if !ok {
		c.logger.Errorf("unexpected object %v", obj)
		return
	}
	c.queue.Add(event{
		eventType: addEcho,
		newObj:    rds.DeepCopy(),
	})
}

func New(
	kubeClientSet kubernetes.Interface,
	rdsClientSet rdsv1alpha1clientset.Interface,
	namespace string,
	logger log.Logger,
) *Controller {

	rdsInformerFactory := rdsinformers.NewSharedInformerFactory(
		rdsClientSet,
		10*time.Second,
	)
	rdsInformer := rdsInformerFactory.Mcsps().V1alpha1().Rdss().Informer()

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClientSet, 10*time.Second)
	jobInformer := kubeInformerFactory.Batch().V1().Jobs().Informer()
	cronjobInformer := kubeInformerFactory.Batch().V1().CronJobs().Informer()

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	ctrl := &Controller{
		kubeClientSet: kubeClientSet,

		rdsInformer:     rdsInformer,
		jobInformer:     jobInformer,
		cronjobInformer: cronjobInformer,

		queue: queue,

		namespace: namespace,

		logger: logger,
	}

	rdsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: ctrl.addEcho,
	})

	return ctrl
}
