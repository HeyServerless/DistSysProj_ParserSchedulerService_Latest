package services

import (
	// "encoding/json"

	// appsv1 "k8s.io/api/core/v1"
	// metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	// "k8s.io/client-go/util/retry"
	// "k8s.io/apimachinery/pkg/runtime/schema"
	// "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func int32Ptr(i int32) *int32 { return &i }

type CreateDeploymentRequest struct {
	Name       string
	Replicas   int32
	Selector   map[string]string
	Labels     map[string]string
	Containers []corev1.Container
}

type CreateDeploymentResponse struct {
	Deployment *appsv1.Deployment
}

type DeploymentList struct {
	Deployments *appsv1.DeploymentList
}

func UpdatePodsInDeployment(deploymentName string, replicas int32) {

	fmt.Println("Updating deployment...name = ", deploymentName, "replicas = ", replicas)
	url := "http://52.7.218.132:8081/UpdateDeployment"
	method := "POST"

	payload := strings.NewReader(`{"name":"` + deploymentName + `","namespace":"default","replicas":` + fmt.Sprint(replicas) + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	log.Println("Deployment updated successfully added replicas = ", replicas)
}

// Get the deployment with the name specified in req
