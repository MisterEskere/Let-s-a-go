package main

import (
	"log"

	"github.com/anacrolix/torrent"
)

var client *torrent.Client

func InitClient() (err error) {
	// Get Clients config
	cfg := torrent.NewDefaultClientConfig()
	cfg.DataDir = "./Torrents" // Where to download files

	// Create the client
	client, err = torrent.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating client: %s", err)
		return err
	}

	return nil
}

func DownloadTorrent(magnet string) (torrent *torrent.Torrent, err error) {

	// Add a magnet link
	torrent, err = client.AddMagnet(magnet)
	if err != nil {
		log.Print("Error adding the torrent")
	}

	// Wait for info before proceeding
	<-torrent.GotInfo()

	// Start downloading all files
	torrent.DownloadAll()

	return torrent, nil
}
