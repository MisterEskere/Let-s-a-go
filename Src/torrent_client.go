package main

import (
	"fmt"
	"log"
	"time"

	"github.com/anacrolix/torrent"
)

func download_torrent() {
	// Create client with default config
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = "./Torrents" // Where to download files

	client, err := torrent.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating client: %s", err)
	}
	defer client.Close()

	// Add a magnet link
	t, err := client.AddMagnet("magnet:?xt=urn:btih:PVJBBJYRFEOXDAOW4B2M4XV5K3Z75XLA&dn=debian-12.10.0-amd64-netinst.iso&xl=663748608&tr=http%3A%2F%2Fbttracker.debian.org%3A6969%2Fannounce")
	if err != nil {
		log.Fatalf("Error adding magnet: %s", err)
	}

	// Wait for torrent info to resolve
	<-t.GotInfo()
	fmt.Printf("Got torrent info: %s\n", t.Name())

	// Start downloading all files
	t.DownloadAll()

	// Wait for download to finish
	for {

		time.Sleep(1 * time.Second)

	}
}
