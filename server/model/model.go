package model

type KubeDeployment struct {
	// kubernete deployment request model
	Title         string `json:"title"`
	Replicas      int    `json:"replicas"`
	ContainerName string `json:"container_name"`
}

type KubeService struct {
	// kubernete service request model
	Title         string `json:"title"`
	ContainerName string `json:"container_name"`
}

type KubePods struct {
	// kubernete pods request model
	Title         string `json:"title"`
	ContainerName string `json:"container_name"`
}

type KubeIngress struct {
	// kubernete ingress request model
	Title         string `json:"title"`
	ContainerName string `json:"container_name"`
}

type KubeSecret struct {
	// kubernete secret request model
	Title         string `json:"title"`
	ContainerName string `json:"container_name"`
}

type DeploymentRequest struct {
	// kubernete deployment request model
	Name      string `json:"name"`
	Namespace string `json:"namespace"`

	Replicas      int    `json:"replicas"`
	AppLabel      string `json:"app_label"`
	ContainerName string `json:"container_name"`
	ImageName     string `json:"image_name"`
	Port          int    `json:"port"`
}

type CreatePodResponse struct {
}

type CreatePodRequest struct {
}

type CreateServiceRequest struct {
}
