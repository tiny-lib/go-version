package go_version

import (
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/PuerkitoBio/goquery"
)

type GoVersionInfoLite struct {
	OS          string `json:"os,omitempty"`
	FileName    string `json:"file_name,omitempty"`
	DownloadUrl string `json:"download_url,omitempty"`
	Size        string `json:"size"`
}

type GoVersionInfoFull struct {
	GoVersionInfoLite
	Arch        string `json:"arch,omitempty"`
	PackageType string `json:"package_type,omitempty"`
	SHA256      string `json:"sha256,omitempty"`
}

type GoVersionInfoList struct {
	Category string
	InfoList []GoVersionInfoFull
}

type DownLoader struct {
	Host string
}

func NewDownLoader(host string) *DownLoader {
	return &DownLoader{Host: host}
}

// GetFeaturedDownload returns the latest version golang download info but with few extra info
func (d DownLoader) GetFeaturedDownload() ([]GoVersionInfoLite, error) {
	result := make([]GoVersionInfoLite, 0)
	hostUrl, _ := urlJoin(d.Host, "dl")
	res, err := http.Get(hostUrl.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	doc.Find("a.downloadBox").Each(func(i int, s *goquery.Selection) {
		versionInfo := GoVersionInfoLite{}
		downloadPath, exist := s.Attr("href")
		if exist {
			downloadUrl, _ := urlJoin(d.Host, downloadPath)
			versionInfo.DownloadUrl = downloadUrl.String()
		}
		versionInfo.OS = s.Find("div.platform").Text()

		versionInfo.FileName = s.Find("span.filename").Text()
		versionInfo.Size = s.Find("span.size").Text()
		result = append(result, versionInfo)
	})
	return result, nil
}

// GetAllDownload returns the All version golang download info
func (d DownLoader) GetAllDownload() ([]GoVersionInfoList, error) {
	var (
		result = make([]GoVersionInfoList, 0)
	)

	hostUrl, _ := urlJoin(d.Host, "dl")
	res, err := http.Get(hostUrl.String())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}
	doc.Find("div.toggle").Each(func(i int, s *goquery.Selection) {
		tag, exist := s.Attr("id")
		if exist {
			versionInfo := GoVersionInfoList{Category: tag}
			versionInfo.InfoList = make([]GoVersionInfoFull, 0)
			s.Find("div.expanded>table.codetable>tbody tr").Each(func(i int, trSelection *goquery.Selection) {
				contents := trSelection.Find("td")
				versionFull := GoVersionInfoFull{}
				contents.Each(func(i int, tdSelection *goquery.Selection) {
					if tdSelection.HasClass("filename") {
						downloadSelection := tdSelection.Find("a.download")
						downloadUrlAttr, exists := downloadSelection.Attr("href")
						if exists {
							downloadUrl, _ := urlJoin(d.Host, downloadUrlAttr)
							versionFull.DownloadUrl = downloadUrl.String()
						}
					} else {
						switch i {
						case 0:
							versionFull.FileName = tdSelection.Text()
						case 1:
							versionFull.PackageType = tdSelection.Text()
						case 2:
							versionFull.OS = tdSelection.Text()
						case 3:
							versionFull.Arch = tdSelection.Text()
						case 4:
							versionFull.Size = tdSelection.Text()
						case 5:
							versionFull.SHA256 = tdSelection.Text()
						}
					}

				})
				versionInfo.InfoList = append(versionInfo.InfoList, versionFull)

			})
			result = append(result, versionInfo)
		}

	})
	return result, nil
}

func urlJoin(Host, Path string) (*url.URL, error) {
	hostUrl, err := url.Parse(Host)
	if err != nil {
		return nil, err
	}
	hostUrl.Path = path.Join(hostUrl.Path, Path)
	return hostUrl, nil
}
