package main

import "flag"

func getFlags() {
	defaultDb := "/usr/local/var/GeoIP/GeoLite2-Country.mmdb"
	defaultDestDir := "./country-subnets"

	flag.StringVar(&db, "db", defaultDb, "database to open")
	flag.StringVar(&destDir, "destDir", defaultDestDir, "target dir to write subnet files")
	flag.Parse()
}
