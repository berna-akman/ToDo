package persistence

import (
	"github.com/couchbase/gocb/v2"
)

type DataSource struct {
	*gocb.Cluster
}

func ConnectCB() (*DataSource, error) {
	clusterAddress := "127.0.0.1"
	username := "Administrator"
	password := "berna1"

	cluster, err := gocb.Connect(clusterAddress, gocb.ClusterOptions{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return &DataSource{Cluster: cluster}, nil
}
