package main

import (
	"fmt"
)

func ShowUsage() {
	usage := `Usage:
dk create :bundlename - Create bundle
dk apply :bundlename - Applies bundle
dk list - Lists all available bundles (located in .dklocal directory)
dk list :remote - Lists all bundles in remote data source
dk pull :remote :bundlename - Pulls bundle from remotes source locally. Name of pulled
  bundle will be "<remote source name>|<bundle name>"
dk push :remote :bundlename - Pushes bundle ot remote (normal way to share bundles)
dk rm :bundlename - Removes local bundle
dk rm :remote :bundlename - Removes bundle from remote source`
	fmt.Println(usage)
}
