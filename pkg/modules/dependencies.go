package modules

const (
	All string = "all"

	HTTPServer string = "http-server"

	Kine                string = "kine"
	KubernetesCRDs      string = "kubernetes-crds"
	KubernetesAPIServer string = "kubernetes-apiserver"
	KubernetesInformers string = "kubernetes-informers"
	KubernetesClientset string = "kubernetes-clientset"
	Kubernetes          string = "kubernetes"

	PublicDashboardsWebhooks string = "public-dashboards-webhooks"
)

var DependencyMap = map[string][]string{
	HTTPServer: {KubernetesAPIServer},

	Kine:                {},
	KubernetesAPIServer: {Kine},
	KubernetesClientset: {KubernetesAPIServer},
	KubernetesCRDs:      {KubernetesClientset},
	KubernetesInformers: {KubernetesCRDs},
	Kubernetes:          {KubernetesInformers},

	PublicDashboardsWebhooks: {KubernetesClientset},

	All: {Kubernetes, HTTPServer, PublicDashboardsWebhooks},
}