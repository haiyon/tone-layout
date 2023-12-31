package utils

import "regexp"

var (
	// match all the images tags
	imageRegex = `(?i)<img.*?src="(.*?)".*?>`
	// match all the video tags
	videoRegex = `(?i)<video.*?src="(.*?)".*?>`
	// match all the audio tags
	audioRegex = `(?i)<audio.*?src="(.*?)".*?>`
	// match all the links tags
	linkRegex = `(?i)<a.*?href="(.*?)".*?>`
)

// GetImagesFrom  - get all the images from content
func GetImagesFrom(content string) []string {
	matches := getMatch(content, imageRegex)
	var images []string
	for _, m := range matches {
		images = append(images, m[1])
	}
	return images
}

// GetVideosFrom  - get all the videos from content
func GetVideosFrom(content string) []string {
	matches := getMatch(content, videoRegex)
	var videos []string
	for _, m := range matches {
		videos = append(videos, m[1])
	}
	return videos
}

// GetAudiosFrom  - get all the audios from content
func GetAudiosFrom(content string) []string {
	matches := getMatch(content, audioRegex)
	var audios []string
	for _, m := range matches {
		audios = append(audios, m[1])
	}
	return audios
}

// GetLinksFrom  - get all the links from content
func GetLinksFrom(content string) []string {
	matches := getMatch(content, linkRegex)
	var links []string
	for _, m := range matches {
		links = append(links, m[1])
	}
	return links
}

// getMatch - get match
func getMatch(content, regex string) [][]string {
	re := regexp.MustCompile(regex)
	return re.FindAllStringSubmatch(content, -1)
}
