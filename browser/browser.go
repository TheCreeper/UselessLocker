package browser

// OpenURL opens a new browser window pointing to the specified url.
func OpenURL(u string) error {
	return openURL(u)
}

// OpenFile opens a new browser pointing to the specified file.
func OpenFile(filename string) error {
	return openURL("file://" + filename)
}
