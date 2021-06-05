package go_version

import "testing"

func TestDownLoader_GetAllDownload(t *testing.T) {
	downLoader := NewDownLoader("https://golang.google.cn")
	info, err := downLoader.GetAllDownload()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", info)
}

func TestDownLoader_GetFeaturedDownload(t *testing.T) {
	downLoader := NewDownLoader("https://golang.google.cn")
	info, err := downLoader.GetFeaturedDownload()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", info)
}
