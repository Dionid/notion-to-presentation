package ntp

import (
	"embed"

	"github.com/Dionid/notion-to-presentation/libs/file"
)

//go:embed sources/*
var embeddedSources embed.FS

func CopySources(dest string) error {
	err := file.CopyFromEmbed(embeddedSources, "sources", dest)
	if err != nil {
		return err
	}

	return nil
}
