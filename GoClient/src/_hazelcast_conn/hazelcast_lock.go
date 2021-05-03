package main

import (
	"fmt"

	"github.com/hazelcast/hazelcast-go-client"
)

// hazelcast-4.2 服务器 端报   Unknown protocol: CB2 ???
//0.6.0是 2020-03-09提交
//Go Client 4.0 (2021 Q2) 对 IMDG 4.0 的改变
func main() {

	config := hazelcast.NewConfig()
	//config.GroupConfig().SetName("myCluster")

	config.GroupConfig().SetName("dev")
	config.NetworkConfig().AddAddress("127.0.0.1:5701")
	client, err := hazelcast.NewClientWithConfig(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client.Name()) // Connects and prints the name of the client
}
