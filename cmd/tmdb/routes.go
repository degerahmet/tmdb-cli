package main

import "fmt"

func getEndpoint(movieType string) (string, error) {
	switch movieType {
	case "playing":
		return "/movie/now_playing", nil
	case "popular":
		return "/movie/popular", nil
	case "top":
		return "/movie/top_rated", nil
	case "upcoming":
		return "/movie/upcoming", nil
	default:
		return "", fmt.Errorf("invalid --type %q (allowed: playing, popular, top, upcoming)", movieType)
	}
}
