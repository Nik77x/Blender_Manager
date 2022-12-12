package Data

import (
	"log"
)

type BlendInfo struct {

	// in build-info -> build-details
	VersionTitle       string `json:"version_title"`
	FileExtension      string `json:"file_extension"`
	FileSize           string `json:"file_size"`
	CommitLink         string `json:"commit_link"`
	CommitHash         string `json:"commit_hash"`
	CommitsHistoryLink string `json:"commits_history_link"`
	UploadDate         string `json:"upload_date"`

	BuildType string `json:"build_type"`

	// in build meta
	DownloadLink       string `json:"download_link"`
	Sha256DownloadLink string `json:"sha_256_download_link"`

	IsDownloading    bool    `json:"is_downloading"`
	DownloadProgress float32 `json:"download_progress"`
}

func (bi *BlendInfo) Print() {

	log.Printf("\nTitle: %q,\nDownload link: %q", bi.VersionTitle, bi.DownloadLink)

}
