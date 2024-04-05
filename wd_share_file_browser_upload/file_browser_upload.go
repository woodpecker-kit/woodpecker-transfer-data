package wd_share_file_browser_upload

const WdShareKeyFileBrowserUpload = "file-browser-upload"

type WdShareFileBrowserUpload struct {
	IsSendSuccess bool `json:"is_send_success"`

	HostUrl             string `json:"host_url"`
	FileBrowserUserName string `json:"file_browser_user_name"`

	// ResourceUrl
	// file browser resource url this is uri for file browser inner path
	ResourceUrl string `json:"resource_url"`
	// DownloadUrl
	// file browser download url is uri for file browser download, if it has passwd, the url will contain token
	DownloadUrl string `json:"download_url"`

	// DownloadPage
	// file browser download page is uri for file browser share page
	DownloadPage string `json:"download_page"`
	// DownloadPasswd
	// file browser download passwd is passwd for DownloadPage
	DownloadPasswd string `json:"download_passwd"`
}
