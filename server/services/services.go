package services

import (
	// "encoding/json"

	// appsv1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"context"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	// "k8s.io/client-go/util/retry"
	// "k8s.io/apimachinery/pkg/runtime/schema"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type CreateServiceRequest struct {
	Name     string
	Selector map[string]string
	Ports    []corev1.ServicePort
	Type     corev1.ServiceType
}

type CreateServiceResponse struct {
	Service *corev1.Service
}

type ServiceList struct {
	Services *corev1.ServiceList
}

func CreateService(c *gin.Context, createServiceRequest *CreateServiceRequest) (*CreateServiceResponse, error) {
	log.Println("================================Creating service================================")
	k8Client := c.MustGet("kubeClient").(*kubernetes.Clientset)

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: createServiceRequest.Name,
		},
		Spec: corev1.ServiceSpec{
			Selector: createServiceRequest.Selector,
			Ports:    createServiceRequest.Ports,
			Type:     createServiceRequest.Type,
		},
	}
	log.Println("service spec: ", service.Spec)
	log.Println("service spec cluster ip: ", service.Spec.ClusterIP)
	log.Println("service spec ports: ", service.Spec.Ports)
	log.Println("service spec selector: ", service.Spec.Selector)
	log.Println("service spec type: ", service.Spec.Type)
	log.Println("service spec status: ", service.Status)

	result, err := k8Client.CoreV1().Services("default").Create(context.Background(), service, metav1.CreateOptions{})
	if err != nil {
		log.Println(err)
		log.Println("Error creating service: %v", err.Error)
		if strings.Contains(err.Error(), "already exists") {
			log.Println("================================Service already exists================================")
			return &CreateServiceResponse{
				Service: result,
			}, nil
		} else {
			log.Println("================================Error creating service================================")
			return nil, err
		}

	}

	// get the status of service

	for {
		serviceStatus, err := k8Client.CoreV1().Services("default").Get(context.Background(), createServiceRequest.Name, metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}
		status := serviceStatus.Status
		if len(service.Spec.Ports) > 0 {
			// service has one or more ports
			log.Println("Service has ports")
			for _, port := range service.Spec.Ports {
				log.Println("port: ", port)

				log.Println("port status: ", status)
				return &CreateServiceResponse{
					Service: result,
				}, nil

			}

		}
		log.Println("Service is not ready")
		sleepTime := 5
		log.Printf("Sleeping for %d seconds", sleepTime)
		time.Sleep(time.Duration(sleepTime) * time.Second)
	}

	// list end points of service
	endpoints, err := k8Client.CoreV1().Endpoints("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("endpoints: ", endpoints)
	for _, endpoint := range endpoints.Items {
		log.Println("endpoint: ", endpoint)
	}

	return &CreateServiceResponse{
		Service: result,
	}, nil

}

func getAllServices(c *gin.Context) (*ServiceList, error) {

	k8Client := c.MustGet("kubeClient").(*kubernetes.Clientset)

	services, err := k8Client.CoreV1().Services("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	return &ServiceList{
		Services: services,
	}, nil

}
