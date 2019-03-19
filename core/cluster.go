package core

//ClusterProvider cluster provider .
type ClusterProvider string

const (
	//Kubernetes k8s cluster
	Kubernetes ClusterProvider = "kubernetes"
	//Swarm docker swarm cluster
	Swarm ClusterProvider = "swarm"
)

type (
	//Cluster cluster definition, include kubernetes and docker swarm.
	Cluster struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Host        string `json:"host"`
		Token       string `json:"token"`
		Certificate string `json:"certificate"`
		Created     int64  `json:"created"`
		Updated     int64  `json:"updated"`
	}
	//ClusterStore cluster storage.
	ClusterStore struct {
	}
)
