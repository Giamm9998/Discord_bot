package crawler

const CVE_DOMAIN = "nvd.nist.gov"
const HN_DOMAIN = "thehackernews.com"

// number of news to collect from website for every scrape
const NEWS_NUMBER = 3

// indexes of news slice
const (
	NEWS_LINK  = 0
	NEWS_TITLE = 1
	NEWS_DATE  = 2
	NEWS_DESC  = 3
	NEWS_IMG   = 4
)

// indexes of cve slice
const (
	CVE_ID    = 0
	VULN_TYPE = 4
	SCOREv2   = 5
	SCOREv3   = 6
	LINK      = 1
)

const CVE_LOGO_URL = "https://www.ibas.tv/2021/wp-content/uploads/2021/11/Logo_CVE_HD.png"
const CVE_THUMBNAIL_URL = "https://assets.website-files.com/5ff66329429d880392f6cba2/61017d1f8c4c0160eb630d51_logo%20CVE.png"
