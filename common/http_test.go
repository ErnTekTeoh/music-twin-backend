package common

import (
	"fmt"
	"testing"
)

func TestCallHTTPEndpoint(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	resp, err := CallHTTPEndpoint(url, "GET", nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", resp)
}
func TestCallHTTPEndpointWithAuth(t *testing.T) {
	url := "https://api.spotify.com/v1/me"
	auth := "Bearer BQA7tWcd4jNxX0_YeHfBSYtNwFUfep5WadAtCNnh4c1LahI1EDgJ2I0UQPiij8kS0KMLZ7SaDXXDvdAhOFWYhAx2zl1hSXGVBa0-KGCCyW0D-QAWabKMqVDcGLqpUGMuLNNFPz5V5B5PQjMrrpHUrSl18-X7ceFkc6gb7wpXWiNzjsAM-pKTv8F1cHXZcPKXmvWYfi7lWKNJyGpLS2a4e7heAutlekdoWVzuulM"
	resp, err := CallHTTPEndpointWithAuth(url, "GET", auth, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", resp)
}

func TestCall(t *testing.T) {
	url := "https://api.discogs.com/database/search?artist=babymonster&title=drip"
	headers := map[string]string{
		"Authorization": "Discogs key=siwseMRvWtdyUGJKuDEv, secret=ODUDvwoAzbbCiyTswyGheMiKqMObxAQU",
		"Cookie":        "__cf_bm=s69EMZXPidl15QC99p0cFDBiH_.YYkMqbPUjF.Q5UHs-1765091907-1.0.1.1-qe6mv36rfxLU2RVyt_MDY6xpj1Wo6Ob.hL3FKJ84C2DVB3b08aUfZcFQSFBntpI1_S8CAEMoK9bgHBK.TkFRSSS4luepMaLW2mtWGT9KuQA",
	}
	resp, err := CallHTTPEndpointWithHeaders(url, "GET", headers, nil)
	fmt.Println(resp)
	fmt.Println(err)
}
