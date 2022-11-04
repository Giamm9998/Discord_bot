package crawler

// Sorting order of the CVEs
type order string

const DOMAIN = "cvedetails.com"

const (
	CVE_NUMBER_DES order = "1"
	CVE_NUMBER_AS        = "2"
	CVE_SCORE_DES        = "3"
	N_EXPLOITS_DES       = "4"
)

const (
	CVE_ID    = 1
	VULN_TYPE = 4
	SCORE     = 7
	LINK      = 15
)

const FIELDS_NUMBER = 3

// decide score of the CVE
const CVSSSCOREMIN = "3"

// substitute with current datetime
const YEAR = "2022"
const MONTH = "7"

var URL = `https://www.cvedetails.com/vulnerability-list.php?vendor_id=0&product_id=0&version_id=0&page=1&hasexp=0&opdos=0&opec=0&opov=0&opcsrf=0&opgpriv=0&opsqli=0&opxss=0&opdirt=0&opmemc=0&ophttprs=0&opbyp=0&opfileinc=0&opginf=0&cvssscoremin=` + CVSSSCOREMIN + `&cvssscoremax=0&year=` + YEAR + `&month=` + MONTH + `&cweid=0&order=` + CVE_SCORE_DES + `&trc=819&sha=e0d67584fe23fbc6f5d20e6a4c010e4f62c16c3b`

const CVE_LOGO_URL = "https://www.ibas.tv/2021/wp-content/uploads/2021/11/Logo_CVE_HD.png"
const CVE_THUMBNAIL_URL = "https://assets.website-files.com/5ff66329429d880392f6cba2/61017d1f8c4c0160eb630d51_logo%20CVE.png"
