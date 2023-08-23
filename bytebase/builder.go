package bytebase


import (
	"log"

	"github.com/go-resty/resty/v2"
)
//
//
func BuildByteBase(url string, debug bool,) (IByteBaseClient, error) {
	byteBaseClient := &byteBase{
		Url:        url,
		restClient: resty.New(),
	}
	// You can override all below settings and options at request level if you want to
	//--------------------------------------------------------------------------------
	// Host URL for all request. So you can use relative URL in the request
	byteBaseClient.restClient.SetBaseURL(url)
	//
	if debug {
		byteBaseClient.restClient.SetDebug(true)
		byteBaseClient.debug = true
		log.Println("Debug mode is enabled for the IbyteBaseClient client ")
	}
	return byteBaseClient, nil
}
//