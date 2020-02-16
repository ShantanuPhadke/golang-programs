package websitechecker

/*WebsiteChecker ... type of function with input of string and return type of bool*/
type WebsiteChecker func(string) bool

/*result ... New type to keep track of our websitechecker outcomes*/
type result struct {
	string
	bool
}

/*CheckWebsites ...
INPUTS:
	wc = WebsiteChecker type object
	urls = a string array of urls
OUTPUTS:
	Returns a map of url strings to the boolean values returned by WebsiteChecker
*/
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// Channel to keep track of our dictionary of results
	resultChannel := make(chan result)

	for _, url := range urls {
		/* Our concurrent goroutine is invoked here */
		go func(currentUrl string) {
			resultChannel <- result{currentUrl, wc(currentUrl)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
