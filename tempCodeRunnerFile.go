func homeUrl(url string) {
	var homepage string
	// if homepage == "" {
	if strings.HasPrefix(url, "http") {
		re := regexp.MustCompile(`https?:\/\/([^\?\/]+)`)
		matches := re.FindStringSubmatch(url)
		if len(matches) > 1 {
			homepage = matches[1]
		}
	}
// }
if homepage == "" {
	fmt.Println("No home page found")
}
fmt.Printf("Homepage: %s\n", homepage)
}