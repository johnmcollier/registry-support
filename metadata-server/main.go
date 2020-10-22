package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/deislabs/oras/pkg/content"
	"github.com/deislabs/oras/pkg/oras"
	indexGenerator "github.com/devfile/registry-support/index/generator/library"
	indexSchema "github.com/devfile/registry-support/index/generator/schema"

	"github.com/containerd/containerd/remotes/docker"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
)

const (
	devfileFileName     = "devfile.yaml"
	devfileMediaType    = "vnd.devfileio.devfile.layer.v1"
	registryDevfilePath = "/registry/stacks"
	registryService     = "localhost:5000"
)

var index []indexSchema.Schema

func serveRegistryIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(index)
}

func main() {
	// Generate the index
	var err error
	index, err = indexGenerator.GenerateIndexStruct(registryDevfilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Before starting the server, push the devfile artifacts to the registry
	for _, devfileIndex := range index {
		err := pushStackToRegistry(devfileIndex)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	// Start the HTTP server and serve the index.json
	http.HandleFunc("/", serveRegistryIndex)
	http.ListenAndServe(":8080", nil)
}

// pushStackToRegistry pushes the given devfile stack to the OCI registry
// If the push fails for whatever reason, an error is returned
func pushStackToRegistry(devfileIndex indexSchema.Schema) error {
	// Load the devfile into memory
	devfileData, err := ioutil.ReadFile(filepath.Join(registryDevfilePath, devfileIndex.Name, devfileFileName))
	if err != nil {
		return err
	}
	ref := registryService + "/" + devfileIndex.Links["self"]

	ctx := context.Background()
	resolver := docker.NewResolver(docker.ResolverOptions{PlainHTTP: true})

	// Add the devfile (and its custom media type) to the memory store
	memoryStore := content.NewMemoryStore()
	desc := memoryStore.Add(devfileFileName, devfileMediaType, devfileData)
	pushContents := []ocispec.Descriptor{desc}

	fmt.Printf("Pushing %s to %s...\n", devfileFileName, ref)
	desc, err = oras.Push(ctx, resolver, ref, memoryStore, pushContents)
	if err != nil {
		return err
	}
	fmt.Printf("Pushed to %s with digest %s\n", ref, desc.Digest)
	return nil
}
