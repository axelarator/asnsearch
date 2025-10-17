package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {

	country := flag.String("country", "", "Country code to filter on (e.g., US)")
	keyword := flag.String("keyword", "", "Keyword to filter on (e.g., bank, vultr)")
	asn := flag.String("asn", "", "ASN to filter on (e.g., AS14836)")
	apiURL := flag.String("api", "https://bgp.potaroo.net/cidr/autnums.html", "API URL to fetch ASN data from")

	flag.Parse()

	if *country == "" && *keyword == "" && *asn == "" {
		flag.Usage()
		fmt.Println("Filters can also be combined. Ex. -country US -keyword vultr")
		return
	}

	req, _ := http.NewRequest("GET", *apiURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	data := string(body)
	result := processData(data, *keyword, *country, *asn)

	out, err := os.Create("asn.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer out.Close()
	io.Copy(out, strings.NewReader(strings.Join(result, "\n")))
	fmt.Println("asn.txt file written to current working directory")

}
func processData(data, keyword, country, asn string) []string {
	lines := strings.Split(data, "\n")
	var filtered []string
	var keywordPattern, countryPattern, asnPattern *regexp.Regexp
	// regex for keyword filter
	keywordPattern = regexp.MustCompile(fmt.Sprintf(`(?i).*%s.*`, regexp.QuoteMeta(keyword)))
	// regex for country filter
	countryPattern = regexp.MustCompile(fmt.Sprintf(`(?i).*, %s.*`, regexp.QuoteMeta(country)))
	// regex for asn filter
	asnNum := strings.TrimPrefix(strings.ToUpper(asn), "AS")
	asnPattern = regexp.MustCompile(fmt.Sprintf(`(?i)AS%s\b`, regexp.QuoteMeta(asnNum)))

	for _, line := range lines {
		match := false
		if keywordPattern != nil && countryPattern != nil && asnPattern != nil {
			if keywordPattern.MatchString(line) && countryPattern.MatchString(line) && asnPattern.MatchString(line) {
				match = true
			}
		} else if keywordPattern != nil {
			if keywordPattern.MatchString(line) {
				match = true
			}
		} else if countryPattern != nil {
			if countryPattern.MatchString(line) {
				match = true
			}
		} else if asnPattern != nil {
			if asnPattern.MatchString(line) {
				match = true
			}
		}
		if match {
			filtered = append(filtered, line)
		}
	}

	var formatted []string
	aTag := regexp.MustCompile(`</a>`)
	openTag := regexp.MustCompile(`^.*>`)

	for _, line := range filtered {
		line = aTag.ReplaceAllString(line, "")
		line = openTag.ReplaceAllString(line, "")
		if strings.TrimSpace(line) != "" {
			formatted = append(formatted, line)
		}
	}
	return formatted
}
