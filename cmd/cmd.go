package main

import (
	"flag"
	"fmt"
	log "github.com/KaguraRinko/sp-dl-go/logger"
	"github.com/KaguraRinko/sp-dl-go/spotify"
	"os"
)

func init() {
	log.Init(log.LevelInfo)
}

func main() {
	println(spotify.SpIDToHex("7r7vaLOWzLmhecQ9Ed2MCw"))
	id := flag.String("id", "", "Spotify URL/URI/ID (required). Example usage: -id https://open.spotify.com/track/4jTrKMoc44RYZsoFsIlQev")
	quality := flag.String("quality", spotify.Quality128MP4Dual, "Quality level. Options: MP4_128, MP4_128_DUAL, MP4_256, MP4_256_DUAL, OGG_VORBIS_320, OGG_VORBIS_160, OGG_VORBIS_96")
	output := flag.String("output", "./output", "Output path.")
	debug := flag.Bool("debug", false, "Print debug information. Use this to enable more detailed logging for troubleshooting.")
	isConvertToMP3 := flag.Bool("mp3", false, "Convert downloaded music to mp3 format")
	isSkipAddingMetadata := flag.Bool("no-metadata", false, "Skip adding metadata to downloaded files.")
	showHelp := flag.Bool("help", false, "Show this help message.")

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	if *id == "" {
		fmt.Println("Error: -id is required")
		flag.Usage()
		os.Exit(1)
	}

	if *debug {
		log.Init(log.LevelDebug)
	}

	sp := spotify.NewDownloader()

	sp.OutputFolder = *output
	log.Infof("Set Output path: %s", *output)

	if err := sp.SetQuality(*quality); err != nil {
		log.Fatalf("Error: %v", err)
	}
	log.Infof("Set quality level: %s", *quality)

	if *isConvertToMP3 {
		sp.ConvertToMP3(*isConvertToMP3)
		log.Infoln("Downloaded music will be converted to mp3")
	}

	if *isSkipAddingMetadata {
		sp.SkipAddingMetadata(*isSkipAddingMetadata)
		log.Infoln("Skip adding metadata to downloaded files")
	}

	log.Infof("Initializing Downloader")
	sp.Initialize()

	err := sp.Download(*id)
	if err != nil {
		log.Fatalln(err)
	}
}
