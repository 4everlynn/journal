package support

// GitLog struct parsed
type GitLog struct {
	author    string
	timestamp int64
	Message   string
}

type GitReleaseAsset struct {
	Url  string `json:"browser_download_url"`
	Name string `json:"name"`
	Size int64  `json:"size"`
}

type GitRelease struct {
	Name    string            `json:"name"`
	TagName string            `json:"tag_name"`
	Assets  []GitReleaseAsset `json:"assets"`
}
