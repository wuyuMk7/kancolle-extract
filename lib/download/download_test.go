package download

import (
	"testing"
)

func TestDownload(t *testing.T) {
	url := "https://github.com"
	fileName := "test_download"

	err := Download(url, fileName)
	if err != nil {
		t.Fatal(err)
	}
}
