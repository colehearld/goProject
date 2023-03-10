//File for creating packets, which host a collection of containers. Once the packets are created, 
//they should be sent to the bundle.
package main

import (
	"context"
	"fmt"
	"os"
	"io/ioutil"
	"gopkg.in/yaml.v2"

	"github.com/docker/docker/client"
)

func packets() {
	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("Failed to create Docker client:", err)
		os.Exit(1)
	}

	// Create a new context
	ctx := context.Background()

	var containerName string
	fmt.Print("Enter the name of a Docker container: ")
	fmt.Scanln(&containerName)

	containerInfo, err := cli.ContainerInspect(ctx, containerName)
	if err != nil {
		fmt.Println("Failed to get information about container:", err)
		os.Exit(1)
	}

	// Store the container ID in a slice
	packet := make([]string, 0) 
	packet = append(packet, containerInfo.ID)

	writePacketToBundle(containerInfo.ID, containerName, "bundlefile.yaml")

}

func getPacket(packet []string) {}

func runPacket() {

}

type Packet struct {
    ID   string `yaml:"id"`
    Name string `yaml:"name"`
}

func writePacketToBundle(containerID string, containerName string, bundlefile string) error {
    p1 := Packet{
        ID:   containerID,
        Name: containerName,
    }

    yamlData, err := yaml.Marshal(&p1)
    if err != nil {
        return fmt.Errorf("error marshaling Packet: %w", err)
    }

    err = ioutil.WriteFile(bundlefile, yamlData, 0644)
    if err != nil {
        return fmt.Errorf("error writing YAML data to file: %w", err)
    }

    return nil
}
