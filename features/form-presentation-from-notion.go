package features

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/Dionid/notion-to-presentation/libs/ntp"
)

func FormFullHtmlPageFromNotion(
	logger ntp.Logger,
	targetUrl string,
) error {
	fmt.Println("Starting to generate presentation...")

	// # Form urls
	parsedUrl, err := url.Parse(targetUrl)
	if err != nil {
		return err
	}

	domain := fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
	mainPageId := ntp.ExtractPageIdFromUrl(parsedUrl)

	// # Get page blocks
	responseChunks, err := ntp.GetNotionBlocks(logger, domain, mainPageId)
	if err != nil {
		return err
	}

	pageTitle, err := ntp.ExtractPageTitle(responseChunks, mainPageId)
	if err != nil {
		return err
	}

	chunkedBlocks, err := ntp.FormChunkedBlocks(logger, domain, responseChunks, mainPageId)
	if err != nil {
		return err
	}

	// # Form html
	html := ntp.FormFullHtmlPage(pageTitle, chunkedBlocks)

	// # Save to file

	// ## Create directory from sources if not exists
	path := filepath.Join(".", mainPageId)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return err
		}

		err = ntp.CopySources(path)
		if err != nil {
			return err
		}
	}

	// ## Write index.html
	err = os.WriteFile(filepath.Join(path, "index.html"), []byte(html), 0644)
	if err != nil {
		return err
	}

	// # Success
	fmt.Println("Done!")

	return nil
}
