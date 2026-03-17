package services

import "os"

func lastFMKey() string  { return os.Getenv("LASTFM_KEY") }
func lastFMUser() string { return os.Getenv("LASTFM_USER") }
