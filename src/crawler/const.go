package crawler

// Sorting order of the CVEs
type order string

const DOMAIN = "nvd.nist.gov"

const (
	CVE_NUMBER_DES order = "1"
	CVE_NUMBER_AS        = "2"
	CVE_SCORE_DES        = "3"
	N_EXPLOITS_DES       = "4"
)

const (
	CVE_ID    = 0
	VULN_TYPE = 4
	SCOREv2   = 5
	SCOREv3   = 6
	LINK      = 1
)

const FIELDS_NUMBER = 3

// decide score of the CVE
const CVSSSCOREMIN = "3"

// substitute with current datetime
const YEAR = "2022"
const MONTH = "7"

const CVE_LOGO_URL = "https://www.ibas.tv/2021/wp-content/uploads/2021/11/Logo_CVE_HD.png"
const CVE_THUMBNAIL_URL = "https://assets.website-files.com/5ff66329429d880392f6cba2/61017d1f8c4c0160eb630d51_logo%20CVE.png"
