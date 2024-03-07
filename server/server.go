package server

import (
	"fmt"

	"github.com/gin-gonic/gin"

	// "honnef.co/go/tools/conf

	"github.com/parserSchedulerService/server/handler"
	services "github.com/parserSchedulerService/server/services"

	"k8s.io/client-go/kubernetes"
	// pb "github.com/example/mypackage"
	// EKS requires the AWS SDK to be imported
)

// Server is srv struct that holds srv Kubernetes client
type Server struct {
	KubeClient *kubernetes.Clientset
}

// func ApiMiddleware(cli *kubernetes.Clientset) gin.HandlerFunc {
// 	// do something with the request
// 	return func(c *gin.Context) {
// 		// do something with the request

// 		c.Set("kubeClient", cli)
// 		c.Next()
// 	}
// }

func (srv *Server) Initialize() {
	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err)
	// }
	// client, err := kubernetes.NewForConfig(config)
	// if err != nil {
	// 	panic(err)
	// }
	// srv = &Server{KubeClient: client}
	// ctx := context.Background()

	//listing the namespaces
	// namespaces, _ := client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	// for _, namespace := range namespaces.Items {
	// 	fmt.Println(namespace.Name)
	// }

	CreateInboundQueue := services.CreateQueue("InboundQueue", "https://5ycsge77e1.execute-api.us-east-1.amazonaws.com/default/sqsCallBackFunction")
	CreateOutbounQueue := services.CreateQueue("OutboundQueue", "https://5ycsge77e1.execute-api.us-east-1.amazonaws.com/default/sqsCallBackFunction")
	CreateSubtractAutoscalingQueue := services.CreateQueue("SubtractAutoscalingQueue", "https://5ycsge77e1.execute-api.us-east-1.amazonaws.com/default/sqsCallBackFunction")
	CreateAddAutoscalingQueue := services.CreateQueue("AddAutoscalingQueue", "https://5ycsge77e1.execute-api.us-east-1.amazonaws.com/default/sqsCallBackFunction")
	CreateMultiplyAutoscalingQueue := services.CreateQueue("MultiplyAutoscalingQueue", "https://5ycsge77e1.execute-api.us-east-1.amazonaws.com/default/sqsCallBackFunction")
	CreateDivisionAutoscalingQueue := services.CreateQueue("DivisionAutoscalingQueue", "https://5ycsge77e1.execute-api.us-east-1.amazonaws.com/default/sqsCallBackFunction")
	fmt.Println(CreateInboundQueue)
	fmt.Println(CreateOutbounQueue)
	fmt.Println(CreateSubtractAutoscalingQueue)
	fmt.Println(CreateAddAutoscalingQueue)
	fmt.Println(CreateMultiplyAutoscalingQueue)
	fmt.Println(CreateDivisionAutoscalingQueue)

	fmt.Println("=================================starting server=================================")
	r := gin.Default()
	// r.Use(ApiMiddleware(client))

	/** temop routes start */

	/** temp routes end*/

	// main service routes
	r.POST("/parseExpression", handler.TRIGER_PARSING)

	r.Run(":8082")
}
