package spotify

import (
	"encoding/json"
	"fmt"
	log "github.com/KaguraRinko/sp-dl-go/logger"
	"net/http"
)

func (d *Downloader) getAlbumTracksAPI(albumID string, offset int) (albumTracksData, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/albums/%s/tracks?offset=%d&limit=50", albumID, offset)
	data, err := d.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Debugf("Fetch Album tracks Failed: %v", err)
		return albumTracksData{}, err
	}

	var albumTracks albumTracksData
	if err := json.Unmarshal(data, &albumTracks); err != nil {
		return albumTracks, fmt.Errorf("failed to decode album data: %w", err)
	}
	return albumTracks, nil
}

func (d *Downloader) getPlaylistTracksAPI(playlistID string, offset int) (playlistTracksData, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks?offset=%d&limit=100", playlistID, offset)
	data, err := d.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Debugf("Fetch Playlist tracks Failed: %v", err)
		return playlistTracksData{}, err
	}

	var playlistTracks playlistTracksData
	if err := json.Unmarshal(data, &playlistTracks); err != nil {
		return playlistTracks, fmt.Errorf("failed to decode playlist data: %w", err)
	}
	return playlistTracks, nil
}

func (d *Downloader) getShowTracksAPI(showID string, offset int) (showTracksData, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/shows/%s/episodes?offset=%d&limit=50", showID, offset)
	data, err := d.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Debugf("Fetch Show episodes Failed: %v", err)
		return showTracksData{}, err
	}

	var showTracks showTracksData
	if err := json.Unmarshal(data, &showTracks); err != nil {
		return showTracks, fmt.Errorf("failed to decode show data: %w", err)
	}
	return showTracks, nil
}

func (d *Downloader) getAlbumAPI(albumID string) (albumData, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/albums/%s", albumID)
	data, err := d.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Debugf("Fetch Album Failed: %v", err)
		return albumData{}, err
	}

	var album albumData
	if err := json.Unmarshal(data, &album); err != nil {
		return albumData{}, fmt.Errorf("failed to decode album data: %w", err)
	}
	return album, nil
}

func (d *Downloader) getTrackAPI(trackID string) (trackData, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", trackID)
	data, err := d.makeRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Debugf("Fetch Track Failed: %v", err)
		return trackData{}, err
	}

	var track trackData
	if err := json.Unmarshal(data, &track); err != nil {
		return trackData{}, fmt.Errorf("failed to decode track data: %w", err)
	}
	return track, nil
}
