package faker

import (
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	f := New()
	if f == nil {
		t.Fatal("New() returned nil")
	}
}

func TestNewSeeded(t *testing.T) {
	f1 := NewSeeded(42)
	f2 := NewSeeded(42)
	if f1.FirstName() != f2.FirstName() {
		t.Errorf("seeded faker should be deterministic: %s != %s", f1.FirstName(), f2.FirstName())
	}
}

func TestFirstName(t *testing.T) {
	f := NewSeeded(1)
	name := f.FirstName()
	if name == "" {
		t.Error("FirstName() returned empty string")
	}
}

func TestFirstNameMale(t *testing.T) {
	f := NewSeeded(1)
	name := f.FirstNameMale()
	if name == "" {
		t.Error("FirstNameMale() returned empty string")
	}
	// Check it's in the male list
	found := false
	for _, n := range firstNamesMale {
		if n == name {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("FirstNameMale() returned %s which is not in male names", name)
	}
}

func TestFirstNameFemale(t *testing.T) {
	f := NewSeeded(1)
	name := f.FirstNameFemale()
	if name == "" {
		t.Error("FirstNameFemale() returned empty string")
	}
	found := false
	for _, n := range firstNamesFemale {
		if n == name {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("FirstNameFemale() returned %s which is not in female names", name)
	}
}

func TestLastName(t *testing.T) {
	f := NewSeeded(1)
	name := f.LastName()
	if name == "" {
		t.Error("LastName() returned empty string")
	}
}

func TestName(t *testing.T) {
	f := NewSeeded(1)
	name := f.Name()
	if name == "" {
		t.Error("Name() returned empty string")
	}
	parts := strings.Split(name, " ")
	if len(parts) != 2 {
		t.Errorf("Name() should have 2 parts, got %d: %s", len(parts), name)
	}
}

func TestEmail(t *testing.T) {
	f := NewSeeded(1)
	email := f.Email()
	if !strings.Contains(email, "@") {
		t.Errorf("Email() should contain @: %s", email)
	}
	if !strings.Contains(email, "example.") {
		t.Errorf("Email() should contain example.: %s", email)
	}
}

func TestPhone(t *testing.T) {
	f := NewSeeded(1)
	phone := f.Phone()
	if !strings.HasPrefix(phone, "+1-") {
		t.Errorf("Phone() should start with +1-: %s", phone)
	}
}

func TestPerson(t *testing.T) {
	f := NewSeeded(1)
	p := f.Person()
	if p.FirstName == "" {
		t.Error("Person.FirstName is empty")
	}
	if p.LastName == "" {
		t.Error("Person.LastName is empty")
	}
	if p.Email == "" {
		t.Error("Person.Email is empty")
	}
	if p.Phone == "" {
		t.Error("Person.Phone is empty")
	}
	if p.Gender != "male" && p.Gender != "female" {
		t.Errorf("Person.Gender should be male or female, got %s", p.Gender)
	}
}

func TestStreetAddress(t *testing.T) {
	f := NewSeeded(1)
	addr := f.StreetAddress()
	if addr == "" {
		t.Error("StreetAddress() returned empty string")
	}
	parts := strings.Split(addr, " ")
	if len(parts) < 3 {
		t.Errorf("StreetAddress() should have at least 3 parts: %s", addr)
	}
}

func TestCity(t *testing.T) {
	f := NewSeeded(1)
	city := f.City()
	if city == "" {
		t.Error("City() returned empty string")
	}
}

func TestState(t *testing.T) {
	f := NewSeeded(1)
	name, abbr := f.State()
	if name == "" {
		t.Error("State() name is empty")
	}
	if len(abbr) != 2 {
		t.Errorf("State() abbreviation should be 2 chars, got %s", abbr)
	}
}

func TestZipCode(t *testing.T) {
	f := NewSeeded(1)
	zip := f.ZipCode()
	if len(zip) != 5 {
		t.Errorf("ZipCode() should be 5 chars, got %s", zip)
	}
}

func TestCountry(t *testing.T) {
	f := NewSeeded(1)
	country := f.Country()
	if country == "" {
		t.Error("Country() returned empty string")
	}
}

func TestAddress(t *testing.T) {
	f := NewSeeded(1)
	a := f.Address()
	if a.Street == "" {
		t.Error("Address.Street is empty")
	}
	if a.City == "" {
		t.Error("Address.City is empty")
	}
	if a.State == "" {
		t.Error("Address.State is empty")
	}
	if a.Zip == "" {
		t.Error("Address.Zip is empty")
	}
	if a.Country == "" {
		t.Error("Address.Country is empty")
	}
}

func TestCompanyName(t *testing.T) {
	f := NewSeeded(1)
	name := f.CompanyName()
	if name == "" {
		t.Error("CompanyName() returned empty string")
	}
	parts := strings.Split(name, " ")
	if len(parts) < 2 {
		t.Errorf("CompanyName() should have at least 2 parts: %s", name)
	}
}

func TestCatchPhrase(t *testing.T) {
	f := NewSeeded(1)
	phrase := f.CatchPhrase()
	if phrase == "" {
		t.Error("CatchPhrase() returned empty string")
	}
	parts := strings.Split(phrase, " ")
	if len(parts) < 4 {
		t.Errorf("CatchPhrase() should have at least 4 words: %s", phrase)
	}
}

func TestBuzzwords(t *testing.T) {
	f := NewSeeded(1)
	words := f.Buzzwords()
	if words == "" {
		t.Error("Buzzwords() returned empty string")
	}
	parts := strings.Split(words, ", ")
	if len(parts) != 3 {
		t.Errorf("Buzzwords() should have 3 words, got %d: %s", len(parts), words)
	}
}

func TestUsername(t *testing.T) {
	f := NewSeeded(1)
	username := f.Username()
	if username == "" {
		t.Error("Username() returned empty string")
	}
}

func TestDomain(t *testing.T) {
	f := NewSeeded(1)
	domain := f.Domain()
	if !strings.Contains(domain, ".") {
		t.Errorf("Domain() should contain a dot: %s", domain)
	}
}

func TestURL(t *testing.T) {
	f := NewSeeded(1)
	url := f.URL()
	if !strings.HasPrefix(url, "http") {
		t.Errorf("URL() should start with http: %s", url)
	}
}

func TestIPv4(t *testing.T) {
	f := NewSeeded(1)
	ip := f.IPv4()
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		t.Errorf("IPv4() should have 4 octets: %s", ip)
	}
}

func TestIPv6(t *testing.T) {
	f := NewSeeded(1)
	ip := f.IPv6()
	parts := strings.Split(ip, ":")
	if len(parts) != 8 {
		t.Errorf("IPv6() should have 8 groups: %s", ip)
	}
}

func TestUserAgent(t *testing.T) {
	f := NewSeeded(1)
	ua := f.UserAgent()
	if !strings.Contains(ua, "Mozilla") {
		t.Errorf("UserAgent() should contain Mozilla: %s", ua)
	}
}

func TestProductName(t *testing.T) {
	f := NewSeeded(1)
	product := f.ProductName()
	if product == "" {
		t.Error("ProductName() returned empty string")
	}
}

func TestPrice(t *testing.T) {
	f := NewSeeded(1)
	price := f.Price()
	if !strings.HasPrefix(price, "$") {
		t.Errorf("Price() should start with $: %s", price)
	}
}

func TestCategory(t *testing.T) {
	f := NewSeeded(1)
	category := f.Category()
	if category == "" {
		t.Error("Category() returned empty string")
	}
}

func TestCreditCard(t *testing.T) {
	f := NewSeeded(1)
	card, cardType := f.CreditCard()
	if card == "" {
		t.Error("CreditCard() returned empty card number")
	}
	validTypes := map[string]bool{"visa": true, "mastercard": true, "amex": true, "discover": true}
	if !validTypes[cardType] {
		t.Errorf("CreditCard() returned invalid type: %s", cardType)
	}
	// Check length
	expectedLen := 16
	if cardType == "amex" {
		expectedLen = 15
	}
	if len(card) != expectedLen {
		t.Errorf("CreditCard() %s should be %d digits, got %d: %s", cardType, expectedLen, len(card), card)
	}
	// Verify Luhn checksum
	if !luhnValid(card) {
		t.Errorf("CreditCard() failed Luhn validation: %s", card)
	}
}

// luhnValid checks if a card number passes the Luhn algorithm.
func luhnValid(number string) bool {
	sum := 0
	for i := 0; i < len(number); i++ {
		digit := int(number[len(number)-1-i] - '0')
		if i%2 == 1 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}

func TestIBAN(t *testing.T) {
	f := NewSeeded(1)
	iban := f.IBAN()
	if len(iban) < 15 {
		t.Errorf("IBAN() too short: %s", iban)
	}
	// Check it starts with a country code (2 uppercase letters)
	if iban[0] < 'A' || iban[0] > 'Z' || iban[1] < 'A' || iban[1] > 'Z' {
		t.Errorf("IBAN() should start with country code: %s", iban)
	}
}

func TestBTCAddress(t *testing.T) {
	f := NewSeeded(1)
	btc := f.BTCAddress()
	if !strings.HasPrefix(btc, "1") {
		t.Errorf("BTCAddress() should start with 1: %s", btc)
	}
	if len(btc) < 26 {
		t.Errorf("BTCAddress() too short: %s", btc)
	}
}

func TestDatePast(t *testing.T) {
	f := NewSeeded(1)
	date := f.DatePast()
	if len(date) != 10 {
		t.Errorf("DatePast() should be 10 chars (YYYY-MM-DD): %s", date)
	}
}

func TestDateFuture(t *testing.T) {
	f := NewSeeded(1)
	date := f.DateFuture()
	if len(date) != 10 {
		t.Errorf("DateFuture() should be 10 chars (YYYY-MM-DD): %s", date)
	}
}

func TestDateRecent(t *testing.T) {
	f := NewSeeded(1)
	date := f.DateRecent()
	if len(date) != 10 {
		t.Errorf("DateRecent() should be 10 chars (YYYY-MM-DD): %s", date)
	}
}

func TestDateBirthday(t *testing.T) {
	f := NewSeeded(1)
	date := f.DateBirthday()
	if len(date) != 10 {
		t.Errorf("DateBirthday() should be 10 chars (YYYY-MM-DD): %s", date)
	}
}

func TestFilePath(t *testing.T) {
	f := NewSeeded(1)
	path := f.FilePath()
	if !strings.HasPrefix(path, "/") {
		t.Errorf("FilePath() should start with /: %s", path)
	}
	if !strings.Contains(path, ".") {
		t.Errorf("FilePath() should contain a file extension: %s", path)
	}
}

func TestMimeType(t *testing.T) {
	f := NewSeeded(1)
	mt := f.MimeType()
	if !strings.Contains(mt, "/") {
		t.Errorf("MimeType() should contain /: %s", mt)
	}
}

func TestFileExtension(t *testing.T) {
	f := NewSeeded(1)
	ext := f.FileExtension()
	if ext == "" {
		t.Error("FileExtension() returned empty string")
	}
}

func TestAvailableTypes(t *testing.T) {
	types := AvailableTypes()
	if len(types) == 0 {
		t.Error("AvailableTypes() returned empty slice")
	}
	// Check some expected types
	expected := map[string]bool{
		"person": true, "address": true, "company": true,
		"internet": true, "commerce": true, "finance": true,
		"date": true, "system": true,
	}
	for _, typ := range types {
		if expected[typ] {
			delete(expected, typ)
		}
	}
	if len(expected) > 0 {
		t.Errorf("AvailableTypes() missing expected types: %v", expected)
	}
}

func TestGenerate(t *testing.T) {
	f := NewSeeded(1)
	for _, typ := range AvailableTypes() {
		result := f.Generate(typ)
		if result == "" {
			t.Errorf("Generate(%q) returned empty string", typ)
		}
	}
}

func TestGenerateJSON(t *testing.T) {
	f := NewSeeded(1)
	for _, typ := range AvailableTypes() {
		result := f.GenerateJSON(typ)
		if result == nil {
			t.Errorf("GenerateJSON(%q) returned nil", typ)
		}
	}
}

func TestGenerateUnknownType(t *testing.T) {
	f := NewSeeded(1)
	result := f.Generate("nonexistent")
	if result != "" {
		t.Errorf("Generate(unknown) should return empty string, got %s", result)
	}
}

func TestGenerateJSONUnknownType(t *testing.T) {
	f := NewSeeded(1)
	result := f.GenerateJSON("nonexistent")
	if result != nil {
		t.Error("GenerateJSON(unknown) should return nil")
	}
}
