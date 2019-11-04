package Utils

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func GetNosDeputesHTTPClient() *resty.Client {
	// Create a Resty Client
	client := resty.New()

	// Retries are configured per client
	client.
		// Enable debug mode
		SetDebug(true).
		// Host URL for all request. So you can use relative URL in the request
		SetHostURL("https://www.nosdeputes.fr").
		// Method available at Client and Request level
		SetDoNotParseResponse(true).
		// Set retry count to non zero to enable retries
		SetRetryCount(3).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(5 * time.Second).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second)

	return client
}
