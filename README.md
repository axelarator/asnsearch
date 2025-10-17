# asnsearch
Credit to Will Thomas for the idea: [My Simple Guide to Mapping ASNs per Sector & Region](https://www.linkedin.com/feed/update/urn:li:activity:7384895078663606272/)

This script allows searching for ASNs by either the ASN itself, a country code, or a keyword. All arguments can be used together too if you want to filter on a specific keyword in a country.

The default service does not require an API key.

When matches are found, results are saved to `asn.txt` in an easy to parse format:

`./asnsearch -country CA -keyword city`
```
AS5110  CITY-OF-VANCOUVER, CA
AS14836 CITY-OF-TORONTO, CA
AS15052 CITY-OF-BURNABY, CA
AS18988 CITYWEST-CORP, CA
```

```bash
Usage of ./asnsearch:
  -api string
        API URL to fetch ASN data from (default "https://bgp.potaroo.net/cidr/autnums.html")
  -asn string
        ASN to filter on (e.g., AS14836)
  -country string
        Country code to filter on (e.g., US)
  -keyword string
        Keyword to filter on (e.g., bank, vultr)
Filters can also be combined. Ex. -country US -keyword vultr
```
