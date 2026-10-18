package faker

import (
	"crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"strings"
	"time"
)

// Faker generates fake data. It is safe for concurrent use.
type Faker struct {
	rng *mrand.Rand
}

// New creates a new Faker seeded from crypto/rand.
func New() *Faker {
	seed, _ := rand.Int(rand.Reader, big.NewInt(1<<62))
	return &Faker{
		rng: mrand.New(mrand.NewSource(seed.Int64())),
	}
}

// NewSeeded creates a new Faker with a deterministic seed.
func NewSeeded(seed int64) *Faker {
	return &Faker{
		rng: mrand.New(mrand.NewSource(seed)),
	}
}

// intn returns a random int in [0, n).
func (f *Faker) intn(n int) int {
	if n <= 0 {
		return 0
	}
	return f.rng.Intn(n)
}

// pick returns a random element from a string slice.
func (f *Faker) pick(arr []string) string {
	if len(arr) == 0 {
		return ""
	}
	return arr[f.intn(len(arr))]
}

// pickN returns n random elements from a string slice (with possible repeats).
func (f *Faker) pickN(arr []string, n int) []string {
	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = f.pick(arr)
	}
	return result
}

// digits returns a string of n random digits.
func (f *Faker) digits(n int) string {
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte(byte('0' + f.intn(10)))
	}
	return sb.String()
}

// randomNumber generates a random int between min and max inclusive.
func (f *Faker) randomNumber(min, max int) int {
	return min + f.intn(max-min+1)
}

// --- Person ---

// Person represents a complete fake person.
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
}

// FirstName generates a random first name.
func (f *Faker) FirstName() string {
	if f.intn(2) == 0 {
		return f.pick(firstNamesMale)
	}
	return f.pick(firstNamesFemale)
}

// FirstNameMale generates a random male first name.
func (f *Faker) FirstNameMale() string {
	return f.pick(firstNamesMale)
}

// FirstNameFemale generates a random female first name.
func (f *Faker) FirstNameFemale() string {
	return f.pick(firstNamesFemale)
}

// LastName generates a random last name.
func (f *Faker) LastName() string {
	return f.pick(lastNames)
}

// Name generates a random full name.
func (f *Faker) Name() string {
	return f.FirstName() + " " + f.LastName()
}

// Email generates a random email address.
func (f *Faker) Email() string {
	first := strings.ToLower(f.FirstName())
	last := strings.ToLower(f.LastName())
	domain := f.pick(tlds)
	sep := f.pick([]string{".", "_", ""})
	num := f.intn(100)
	if num > 0 {
		return fmt.Sprintf("%s%s%s%d@example.%s", first, sep, last, num, domain)
	}
	return fmt.Sprintf("%s%s%s@example.%s", first, sep, last, domain)
}

// Phone generates a random US phone number.
func (f *Faker) Phone() string {
	area := f.randomNumber(200, 999)
	exchange := f.randomNumber(200, 999)
	subscriber := f.randomNumber(1000, 9999)
	return fmt.Sprintf("+1-%d-%d-%d", area, exchange, subscriber)
}

// Person generates a complete fake person.
func (f *Faker) Person() Person {
	male := f.intn(2) == 0
	first := ""
	gender := "female"
	if male {
		first = f.FirstNameMale()
		gender = "male"
	} else {
		first = f.FirstNameFemale()
	}
	last := f.LastName()
	return Person{
		FirstName: first,
		LastName:  last,
		Email:     f.Email(),
		Phone:     f.Phone(),
		Gender:    gender,
	}
}

// --- Address ---

// Address represents a complete fake address.
type Address struct {
	Street    string `json:"street"`
	City      string `json:"city"`
	State     string `json:"state"`
	StateCode string `json:"state_code"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
}

// StreetAddress generates a random street address.
func (f *Faker) StreetAddress() string {
	num := f.randomNumber(1, 9999)
	name := f.pick(streetNames)
	suffix := f.pick(streetSuffixes)
	return fmt.Sprintf("%d %s %s", num, name, suffix)
}

// City generates a random city name.
func (f *Faker) City() string {
	return f.pick(cities)
}

// State generates a random US state.
func (f *Faker) State() (string, string) {
	s := usStates[f.intn(len(usStates))]
	return s.Name, s.Abbr
}

// ZipCode generates a random US zip code.
func (f *Faker) ZipCode() string {
	return fmt.Sprintf("%05d", f.randomNumber(1001, 99950))
}

// Country generates a random country name.
func (f *Faker) Country() string {
	return f.pick(countries)
}

// Address generates a complete fake address.
func (f *Faker) Address() Address {
	stateName, stateAbbr := f.State()
	return Address{
		Street:    f.StreetAddress(),
		City:      f.City(),
		State:     stateName,
		StateCode: stateAbbr,
		Zip:       f.ZipCode(),
		Country:   "United States",
	}
}

// --- Company ---

// Company represents a complete fake company.
type Company struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	CatchPhrase string `json:"catch_phrase"`
	Buzzwords   string `json:"buzzwords"`
}

// CompanyName generates a random company name.
func (f *Faker) CompanyName() string {
	return f.pick(companyPrefixes) + " " + f.pick(companySuffixes)
}

// CompanyType generates a random company type/industry.
func (f *Faker) CompanyType() string {
	return f.pick(companyTypes)
}

// CatchPhrase generates a random catch phrase.
func (f *Faker) CatchPhrase() string {
	return strings.Join([]string{
		f.pick(catchPhraseAdjectives),
		f.pick(catchPhraseSubjects),
		f.pick(catchPhraseVerbs),
		f.pick(catchPhraseObjects),
	}, " ")
}

// Buzzwords generates 3 random buzzwords.
func (f *Faker) Buzzwords() string {
	words := f.pickN(buzzwords, 3)
	return strings.Join(words, ", ")
}

// Company generates a complete fake company.
func (f *Faker) Company() Company {
	return Company{
		Name:        f.CompanyName(),
		Type:        f.CompanyType(),
		CatchPhrase: f.CatchPhrase(),
		Buzzwords:   f.Buzzwords(),
	}
}

// --- Internet ---

// Internet represents fake internet-related data.
type Internet struct {
	Username  string `json:"username"`
	Domain    string `json:"domain"`
	URL       string `json:"url"`
	IP        string `json:"ip"`
	UserAgent string `json:"user_agent"`
}

// Username generates a random username.
func (f *Faker) Username() string {
	first := strings.ToLower(f.FirstName())
	last := strings.ToLower(f.LastName())
	sep := f.pick([]string{"_", ".", ""})
	num := f.intn(100)
	if num > 0 {
		return fmt.Sprintf("%s%s%s%d", first, sep, last, num)
	}
	return first + sep + last
}

// Domain generates a random domain name.
func (f *Faker) Domain() string {
	prefix := strings.ToLower(f.pick(companyPrefixes))
	tld := f.pick(tlds)
	return prefix + "." + tld
}

// URL generates a random URL.
func (f *Faker) URL() string {
	scheme := f.pick([]string{"https", "http"})
	return fmt.Sprintf("%s://%s", scheme, f.Domain())
}

// IPv4 generates a random IPv4 address.
func (f *Faker) IPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		f.randomNumber(1, 255),
		f.randomNumber(0, 255),
		f.randomNumber(0, 255),
		f.randomNumber(1, 254),
	)
}

// IPv6 generates a random IPv6 address.
func (f *Faker) IPv6() string {
	parts := make([]string, 8)
	for i := 0; i < 8; i++ {
		parts[i] = fmt.Sprintf("%x", f.randomNumber(0, 65535))
	}
	return strings.Join(parts, ":")
}

// UserAgent generates a random user agent string.
func (f *Faker) UserAgent() string {
	return f.pick(userAgentStrings)
}

// Internet generates complete fake internet data.
func (f *Faker) Internet() Internet {
	return Internet{
		Username:  f.Username(),
		Domain:    f.Domain(),
		URL:       f.URL(),
		IP:        f.IPv4(),
		UserAgent: f.UserAgent(),
	}
}

// --- Commerce ---

// Commerce represents fake commerce data.
type Commerce struct {
	Product  string `json:"product"`
	Price    string `json:"price"`
	Category string `json:"category"`
}

// ProductName generates a random product name.
func (f *Faker) ProductName() string {
	return f.pick(productAdjectives) + " " + f.pick(productNouns)
}

// Price generates a random price string.
func (f *Faker) Price() string {
	whole := f.randomNumber(1, 999)
	cents := f.randomNumber(0, 99)
	return fmt.Sprintf("$%d.%02d", whole, cents)
}

// Category generates a random product category.
func (f *Faker) Category() string {
	return f.pick(productCategories)
}

// Commerce generates complete fake commerce data.
func (f *Faker) Commerce() Commerce {
	return Commerce{
		Product:  f.ProductName(),
		Price:    f.Price(),
		Category: f.Category(),
	}
}

// --- Finance ---

// Finance represents fake financial data.
type Finance struct {
	CreditCard string `json:"credit_card"`
	CardType   string `json:"card_type"`
	IBAN       string `json:"iban"`
	BTC        string `json:"btc_address"`
}

// CreditCard generates a random credit card number with valid Luhn checksum.
func (f *Faker) CreditCard() (string, string) {
	cardTypes := []string{"visa", "mastercard", "amex", "discover"}
	cardType := f.pick(cardTypes)
	prefixes := creditCardPrefixes[cardType]
	prefix := f.pick(prefixes)

	length := 16
	if cardType == "amex" {
		length = 15
	}

	bodyLen := length - len(prefix) - 1
	body := f.digits(bodyLen)
	partial := prefix + body

	checkDigit := luhnCheckDigit(partial)
	return partial + checkDigit, cardType
}

// luhnCheckDigit calculates the Luhn check digit for a given number string.
func luhnCheckDigit(number string) string {
	sum := 0
	for i := 0; i < len(number); i++ {
		digit := int(number[len(number)-1-i] - '0')
		if i%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	check := (10 - (sum % 10)) % 10
	return fmt.Sprintf("%d", check)
}

// IBAN generates a random IBAN.
func (f *Faker) IBAN() string {
	c := ibanCountries[f.intn(len(ibanCountries))]
	bbanLen := c.Length - 4
	bban := f.digits(bbanLen)
	rearranged := bban + c.Code + "00"
	numeric := ibanToNumeric(rearranged)
	num, _ := new(big.Int).SetString(numeric, 10)
	mod := new(big.Int).Mod(num, big.NewInt(97))
	check := new(big.Int).Sub(big.NewInt(98), mod)
	return c.Code + fmt.Sprintf("%02d", check.Int64()) + bban
}

// ibanToNumeric converts an IBAN string to its numeric representation.
func ibanToNumeric(s string) string {
	var sb strings.Builder
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			sb.WriteRune(ch)
		} else if ch >= 'A' && ch <= 'Z' {
			sb.WriteString(fmt.Sprintf("%d", ch-'A'+10))
		}
	}
	return sb.String()
}

// BTCAddress generates a random Bitcoin address (mainnet, P2PKH format).
func (f *Faker) BTCAddress() string {
	const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	length := f.randomNumber(26, 34)
	var sb strings.Builder
	sb.WriteByte('1')
	for i := 1; i < length; i++ {
		sb.WriteByte(alphabet[f.intn(len(alphabet))])
	}
	return sb.String()
}

// Finance generates complete fake financial data.
func (f *Faker) Finance() Finance {
	card, cardType := f.CreditCard()
	return Finance{
		CreditCard: card,
		CardType:   cardType,
		IBAN:       f.IBAN(),
		BTC:        f.BTCAddress(),
	}
}

// --- Date ---

// DateInfo represents fake date data.
type DateInfo struct {
	Past     string `json:"past"`
	Future   string `json:"future"`
	Recent   string `json:"recent"`
	Birthday string `json:"birthday"`
}

// DatePast generates a random date in the past (within the last 5 years).
func (f *Faker) DatePast() string {
	now := time.Now()
	secondsAgo := f.intn(365 * 24 * 3600 * 5)
	t := now.Add(-time.Duration(secondsAgo) * time.Second)
	return t.Format("2006-01-02")
}

// DateFuture generates a random date in the future (within the next 5 years).
func (f *Faker) DateFuture() string {
	now := time.Now()
	secondsAhead := f.intn(365 * 24 * 3600 * 5)
	t := now.Add(time.Duration(secondsAhead) * time.Second)
	return t.Format("2006-01-02")
}

// DateRecent generates a random recent date (within the last 30 days).
func (f *Faker) DateRecent() string {
	now := time.Now()
	secondsAgo := f.intn(30 * 24 * 3600)
	t := now.Add(-time.Duration(secondsAgo) * time.Second)
	return t.Format("2006-01-02")
}

// DateBirthday generates a random birthday (between 18 and 90 years ago).
func (f *Faker) DateBirthday() string {
	now := time.Now()
	ageYears := f.randomNumber(18, 90)
	daysOffset := f.intn(365)
	t := now.AddDate(-ageYears, 0, -daysOffset)
	return t.Format("2006-01-02")
}

// Date generates complete fake date data.
func (f *Faker) Date() DateInfo {
	return DateInfo{
		Past:     f.DatePast(),
		Future:   f.DateFuture(),
		Recent:   f.DateRecent(),
		Birthday: f.DateBirthday(),
	}
}

// --- System ---

// System represents fake system data.
type System struct {
	FilePath  string `json:"file_path"`
	MimeType  string `json:"mime_type"`
	Extension string `json:"extension"`
}

// FilePath generates a random file path.
func (f *Faker) FilePath() string {
	depth := f.randomNumber(1, 4)
	parts := make([]string, depth)
	for i := 0; i < depth; i++ {
		parts[i] = f.pick(filePathDirs)
	}
	name := f.pick(filePathNames)
	ext := f.pick(fileExtensions)
	return "/" + strings.Join(parts, "/") + "/" + name + "." + ext
}

// MimeType generates a random MIME type.
func (f *Faker) MimeType() string {
	return f.pick(mimeTypes)
}

// FileExtension generates a random file extension.
func (f *Faker) FileExtension() string {
	return f.pick(fileExtensions)
}

// System generates complete fake system data.
func (f *Faker) System() System {
	return System{
		FilePath:  f.FilePath(),
		MimeType:  f.MimeType(),
		Extension: f.FileExtension(),
	}
}

// --- Types ---

// AvailableTypes returns all available generation types.
func AvailableTypes() []string {
	return []string{
		"person", "person.name", "person.first", "person.last",
		"address", "address.street", "address.city", "address.state", "address.zip", "address.country",
		"company", "company.name", "company.catchphrase", "company.buzzwords",
		"internet", "internet.username", "internet.domain", "internet.url", "internet.ip", "internet.useragent",
		"commerce", "commerce.product", "commerce.price", "commerce.category",
		"finance", "finance.creditcard", "finance.iban", "finance.btc",
		"date", "date.past", "date.future", "date.recent", "date.birthday",
		"system", "system.filepath", "system.mimetype", "system.extension",
	}
}

// Generate generates a single value for the given type as a plain text string.
func (f *Faker) Generate(typ string) string {
	switch typ {
	case "person":
		p := f.Person()
		return fmt.Sprintf("name=%s %s  email=%s  phone=%s  gender=%s",
			p.FirstName, p.LastName, p.Email, p.Phone, p.Gender)
	case "person.name":
		return f.Name()
	case "person.first":
		return f.FirstName()
	case "person.last":
		return f.LastName()
	case "address":
		a := f.Address()
		return fmt.Sprintf("street=%s  city=%s  state=%s  zip=%s  country=%s",
			a.Street, a.City, a.State, a.Zip, a.Country)
	case "address.street":
		return f.StreetAddress()
	case "address.city":
		return f.City()
	case "address.state":
		name, _ := f.State()
		return name
	case "address.zip":
		return f.ZipCode()
	case "address.country":
		return f.Country()
	case "company":
		c := f.Company()
		return fmt.Sprintf("name=%s  type=%s  catchphrase=%s  buzzwords=%s",
			c.Name, c.Type, c.CatchPhrase, c.Buzzwords)
	case "company.name":
		return f.CompanyName()
	case "company.catchphrase":
		return f.CatchPhrase()
	case "company.buzzwords":
		return f.Buzzwords()
	case "internet":
		i := f.Internet()
		return fmt.Sprintf("username=%s  domain=%s  url=%s  ip=%s  useragent=%s",
			i.Username, i.Domain, i.URL, i.IP, i.UserAgent)
	case "internet.username":
		return f.Username()
	case "internet.domain":
		return f.Domain()
	case "internet.url":
		return f.URL()
	case "internet.ip":
		return f.IPv4()
	case "internet.useragent":
		return f.UserAgent()
	case "commerce":
		c := f.Commerce()
		return fmt.Sprintf("product=%s  price=%s  category=%s",
			c.Product, c.Price, c.Category)
	case "commerce.product":
		return f.ProductName()
	case "commerce.price":
		return f.Price()
	case "commerce.category":
		return f.Category()
	case "finance":
		fin := f.Finance()
		return fmt.Sprintf("creditcard=%s  cardtype=%s  iban=%s  btc=%s",
			fin.CreditCard, fin.CardType, fin.IBAN, fin.BTC)
	case "finance.creditcard":
		card, cardType := f.CreditCard()
		return fmt.Sprintf("%s (%s)", card, cardType)
	case "finance.iban":
		return f.IBAN()
	case "finance.btc":
		return f.BTCAddress()
	case "date":
		d := f.Date()
		return fmt.Sprintf("past=%s  future=%s  recent=%s  birthday=%s",
			d.Past, d.Future, d.Recent, d.Birthday)
	case "date.past":
		return f.DatePast()
	case "date.future":
		return f.DateFuture()
	case "date.recent":
		return f.DateRecent()
	case "date.birthday":
		return f.DateBirthday()
	case "system":
		s := f.System()
		return fmt.Sprintf("filepath=%s  mimetype=%s  extension=%s",
			s.FilePath, s.MimeType, s.Extension)
	case "system.filepath":
		return f.FilePath()
	case "system.mimetype":
		return f.MimeType()
	case "system.extension":
		return f.FileExtension()
	default:
		return ""
	}
}

// GenerateJSON generates a structured value for the given type.
func (f *Faker) GenerateJSON(typ string) interface{} {
	switch typ {
	case "person":
		return f.Person()
	case "person.name":
		return f.Name()
	case "person.first":
		return f.FirstName()
	case "person.last":
		return f.LastName()
	case "address":
		return f.Address()
	case "address.street":
		return f.StreetAddress()
	case "address.city":
		return f.City()
	case "address.state":
		name, _ := f.State()
		return name
	case "address.zip":
		return f.ZipCode()
	case "address.country":
		return f.Country()
	case "company":
		return f.Company()
	case "company.name":
		return f.CompanyName()
	case "company.catchphrase":
		return f.CatchPhrase()
	case "company.buzzwords":
		return f.Buzzwords()
	case "internet":
		return f.Internet()
	case "internet.username":
		return f.Username()
	case "internet.domain":
		return f.Domain()
	case "internet.url":
		return f.URL()
	case "internet.ip":
		return f.IPv4()
	case "internet.useragent":
		return f.UserAgent()
	case "commerce":
		return f.Commerce()
	case "commerce.product":
		return f.ProductName()
	case "commerce.price":
		return f.Price()
	case "commerce.category":
		return f.Category()
	case "finance":
		return f.Finance()
	case "finance.creditcard":
		card, cardType := f.CreditCard()
		return map[string]string{"number": card, "type": cardType}
	case "finance.iban":
		return f.IBAN()
	case "finance.btc":
		return f.BTCAddress()
	case "date":
		return f.Date()
	case "date.past":
		return f.DatePast()
	case "date.future":
		return f.DateFuture()
	case "date.recent":
		return f.DateRecent()
	case "date.birthday":
		return f.DateBirthday()
	case "system":
		return f.System()
	case "system.filepath":
		return f.FilePath()
	case "system.mimetype":
		return f.MimeType()
	case "system.extension":
		return f.FileExtension()
	default:
		return nil
	}
}
