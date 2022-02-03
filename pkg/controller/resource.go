package controller

import (
	echo "github.com/mmontes11/echoperator/pkg/echo"
	echov1alpha1 "github.com/mmontes11/echoperator/pkg/echo/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createJob(newEcho *echov1alpha1.Echo, namespace string) *batchv1.Job {
	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      newEcho.ObjectMeta.Name,
			Namespace: namespace,
			Labels:    make(map[string]string),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(
					newEcho,
					echov1alpha1.SchemeGroupVersion.WithKind(echo.EchoKind),
				),
			},
		},
		Spec: createJobSpec(newEcho.Name, namespace, newEcho.Spec.Message),
	}
}

func createCronJob(
	newScheduledEcho *echov1alpha1.ScheduledEcho,
	namespace string,
) *batchv1.CronJob {
	return &batchv1.CronJob{
		ObjectMeta: metav1.ObjectMeta{
			Name:      newScheduledEcho.ObjectMeta.Name,
			Namespace: namespace,
			Labels:    make(map[string]string),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(
					newScheduledEcho,
					echov1alpha1.SchemeGroupVersion.WithKind(echo.ScheduledEchoKind),
				),
			},
		},
		Spec: batchv1.CronJobSpec{
			Schedule:          newScheduledEcho.Spec.Schedule,
			ConcurrencyPolicy: batchv1.ForbidConcurrent,
			JobTemplate: batchv1.JobTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: newScheduledEcho.Name + "-",
					Namespace:    namespace,
					Labels:       make(map[string]string),
				},
				Spec: createJobSpec(
					newScheduledEcho.Name,
					namespace,
					newScheduledEcho.Spec.Message,
				),
			},
		},
	}
}

func createJobSpec(name, namespace, msg string) batchv1.JobSpec {
	return batchv1.JobSpec{
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: name + "-",
				Namespace:    namespace,
				Labels:       make(map[string]string),
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:            name,
						Image:           "ghcr.io/eumel8/echobusybox:latest",
						Command:         []string{"echo", msg},
						ImagePullPolicy: "IfNotPresent",
					},
				},
				RestartPolicy: corev1.RestartPolicyNever,
			},
		},
	}
}
