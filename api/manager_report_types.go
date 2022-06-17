package api

import (
	"fmt"
	"io"
)

// ClientType is required by API.
const (
	EmailClientType = iota + 1
	VPNClientType
)

type reportAtt struct {
	name, mime string
	body       io.Reader
}

// ReportBugReq stores data for report.
type ReportBugReq struct {
	OS                string      `json:",omitempty"`
	OSVersion         string      `json:",omitempty"`
	Browser           string      `json:",omitempty"`
	BrowserVersion    string      `json:",omitempty"`
	BrowserExtensions string      `json:",omitempty"`
	Resolution        string      `json:",omitempty"`
	DisplayMode       string      `json:",omitempty"`
	Client            string      `json:",omitempty"`
	ClientVersion     string      `json:",omitempty"`
	ClientType        int         `json:",omitempty"`
	Title             string      `json:",omitempty"`
	Description       string      `json:",omitempty"`
	Username          string      `json:",omitempty"`
	Email             string      `json:",omitempty"`
	Country           string      `json:",omitempty"`
	ISP               string      `json:",omitempty"`
	Debug             string      `json:",omitempty"`
	Attachments       []reportAtt `json:",omitempty"`
}

// AddAttachment to report.
func (rep *ReportBugReq) AddAttachment(name, mime string, r io.Reader) {
	rep.Attachments = append(rep.Attachments, reportAtt{name: name, mime: mime, body: r})
}

func (rep *ReportBugReq) GetMultipartFormData() map[string]string {
	return map[string]string{
		"OS":                rep.OS,
		"OSVersion":         rep.OSVersion,
		"Browser":           rep.Browser,
		"BrowserVersion":    rep.BrowserVersion,
		"BrowserExtensions": rep.BrowserExtensions,
		"Resolution":        rep.Resolution,
		"DisplayMode":       rep.DisplayMode,
		"Client":            rep.Client,
		"ClientVersion":     rep.ClientVersion,
		"ClientType":        fmt.Sprintf("%d", rep.ClientType),
		"Title":             rep.Title,
		"Description":       rep.Description,
		"Username":          rep.Username,
		"Email":             rep.Email,
		"Country":           rep.Country,
		"ISP":               rep.ISP,
		"Debug":             rep.Debug,
	}
}
