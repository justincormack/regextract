package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	"github.com/heroku/docker-registry-client/registry"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatalf("usage: register image[:tag] files...")
	}
	tag := "latest"
	it := strings.SplitN(flag.Arg(0), ":", 2)
	if len(it) == 2 {
		tag = it[1]
	}
	image := it[0]

	allfiles := flag.NArg() == 1
	files := flag.Args()[1:]
	fileset := make(map[string]bool)
	for _, v := range files {
		fileset[v] = true
	}

	url := "https://registry-1.docker.io/"
	username := ""
	password := ""
	hub, err := registry.New(url, username, password)
	if err != nil {
		log.Fatalf("Cannot connect to registry")
	}

	manifest, err := hub.Manifest(image, tag)
	if err != nil {
		log.Fatalf("Cannot fetch manifest: %s", err)
	}

	log.Printf("Found %d manifest layers, using layer %d", len(manifest.FSLayers), len(manifest.FSLayers)-1)
	layer := manifest.FSLayers[len(manifest.FSLayers)-1]

	reader, err := hub.DownloadLayer(image, layer.BlobSum)
	if err != nil {
		log.Fatalf("cannot read layer")
	}
	defer reader.Close()

	unzipper, err := gzip.NewReader(reader)
	if err != nil {
		log.Fatalf("Cannot uncompress: %s", err)
	}

	tr := tar.NewReader(unzipper)
	tw := tar.NewWriter(os.Stdout)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		if fileset[hdr.Name] || allfiles {
			err = tw.WriteHeader(hdr)
			if err != nil {
				log.Fatalf("cannot write to tar: %s", err)
			}
			_, err = io.Copy(tw, tr)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}
	err = tw.Close()
	if err != nil {
		log.Fatalf("Tar close error: %s", err)
	}
}
