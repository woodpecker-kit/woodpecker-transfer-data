package wd_share_file_browser_upload

const WdShareKeyFileBrowserUpload = "file-browser-upload"

type WdShareFileBrowserUpload struct {
	IsSendSuccess bool `json:"is_send_success"`

	HostUrl             string `json:"host_url"`
	FileBrowserUserName string `json:"file_browser_user_name"`

	ResourceUrl    string `json:"resource_url"`
	DownloadUrl    string `json:"download_url"`
	DownloadPasswd string `json:"download_passwd"`
}
