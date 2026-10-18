package faker

// Embedded data sets for fake data generation.
// No external dependencies — all data is compiled into the binary.

var firstNamesMale = []string{
	"James", "John", "Robert", "Michael", "William", "David", "Richard", "Joseph",
	"Thomas", "Charles", "Christopher", "Daniel", "Matthew", "Anthony", "Mark",
	"Donald", "Steven", "Paul", "Andrew", "Joshua", "Kenneth", "Kevin", "Brian",
	"George", "Edward", "Ronald", "Timothy", "Jason", "Jeffrey", "Ryan",
	"Jacob", "Gary", "Nicholas", "Eric", "Jonathan", "Stephen", "Larry",
	"Justin", "Scott", "Brandon", "Benjamin", "Samuel", "Gregory", "Frank",
	"Alexander", "Raymond", "Patrick", "Jack", "Dennis", "Jerry", "Tyler",
	"Aaron", "Henry", "Douglas", "Peter", "Jose", "Adam", "Nathan", "Zachary",
	"Walter", "Harold", "Kyle", "Carl", "Arthur", "Gerald", "Roger", "Keith",
	"Jeremy", "Lawrence", "Terry", "Sean", "Christian", "Ethan", "Austin",
	"Joe", "Albert", "Jesse", "Willie", "Billy", "Bryan", "Bruce", "Jordan",
	"Ralph", "Roy", "Noah", "Dylan", "Eugene", "Wayne", "Alan", "Juan",
	"Russell", "Randy", "Mason", "Vincent", "Philip", "Bobby", "Johnny",
}

var firstNamesFemale = []string{
	"Mary", "Patricia", "Jennifer", "Linda", "Elizabeth", "Barbara", "Susan",
	"Jessica", "Sarah", "Karen", "Lisa", "Nancy", "Betty", "Sandra", "Margaret",
	"Ashley", "Kimberly", "Emily", "Donna", "Michelle", "Carol", "Amanda",
	"Melissa", "Deborah", "Stephanie", "Dorothy", "Rebecca", "Sharon", "Laura",
	"Cynthia", "Amy", "Kathleen", "Angela", "Shirley", "Anna", "Brenda",
	"Pamela", "Emma", "Nicole", "Helen", "Samantha", "Katherine", "Christine",
	"Debra", "Rachel", "Carolyn", "Janet", "Catherine", "Maria", "Heather",
	"Diane", "Ruth", "Julie", "Olivia", "Joyce", "Virginia", "Victoria",
	"Kelly", "Lauren", "Christina", "Joan", "Evelyn", "Judith", "Andrea",
	"Hannah", "Megan", "Cheryl", "Jacqueline", "Martha", "Gloria", "Teresa",
	"Ann", "Sara", "Madison", "Frances", "Kathryn", "Janice", "Jean",
	"Abigail", "Alice", "Judy", "Sophia", "Grace", "Denise", "Amber",
	"Doris", "Marilyn", "Danielle", "Beverly", "Isabella", "Theresa", "Diana",
	"Natalie", "Brittany", "Charlotte", "Marie", "Kayla", "Alexis", "Lori",
}

var lastNames = []string{
	"Smith", "Johnson", "Williams", "Brown", "Jones", "Garcia", "Miller",
	"Davis", "Rodriguez", "Martinez", "Hernandez", "Lopez", "Gonzalez",
	"Wilson", "Anderson", "Thomas", "Taylor", "Moore", "Jackson", "Martin",
	"Lee", "Perez", "Thompson", "White", "Harris", "Sanchez", "Clark",
	"Ramirez", "Lewis", "Robinson", "Walker", "Young", "Allen", "King",
	"Wright", "Scott", "Torres", "Nguyen", "Hill", "Flores", "Green",
	"Adams", "Nelson", "Baker", "Hall", "Rivera", "Campbell", "Mitchell",
	"Carter", "Roberts", "Gomez", "Phillips", "Evans", "Turner", "Diaz",
	"Parker", "Cruz", "Edwards", "Collins", "Reyes", "Stewart", "Morris",
	"Morales", "Murphy", "Cook", "Rogers", "Gutierrez", "Ortiz", "Morgan",
	"Cooper", "Peterson", "Bailey", "Reed", "Kelly", "Howard", "Ramos",
	"Kim", "Cox", "Ward", "Richardson", "Watson", "Brooks", "Chavez",
	"Wood", "James", "Bennett", "Gray", "Mendoza", "Ruiz", "Hughes",
	"Price", "Alvarez", "Castillo", "Sanders", "Patel", "Myers", "Long",
	"Ross", "Foster", "Jimenez", "Powell", "Jenkins", "Perry", "Russell",
	"Sullivan", "Bell", "Coleman", "Butler", "Henderson", "Barnes", "Gonzales",
	"Fisher", "Vasquez", "Simmons", "Romero", "Jordan", "Patterson", "Alexander",
	"Hamilton", "Graham", "Reynolds", "Griffin", "Wallace", "Moreno", "West",
	"Cole", "Hayes", "Bryant", "Herrera", "Gibson", "Ellis", "Tran",
	"Medina", "Aguilar", "Stevens", "Murray", "Ford", "Castro", "Marshall",
	"Owens", "Harrison", "Fernandez", "Mcdonald", "Woods", "Washington", "Kennedy",
}

var streetNames = []string{
	"Main", "Oak", "Pine", "Maple", "Cedar", "Elm", "Washington", "Lake",
	"Hill", "Park", "Spring", "North", "South", "East", "West", "Highland",
	"Forest", "River", "Sunset", "Lincoln", "Jefferson", "Madison", "Jackson",
	"Franklin", "Adams", "Monroe", "Roosevelt", "Kennedy", "Church", "Mill",
	"Center", "Union", "Market", "Cherry", "Walnut", "Willow", "Birch",
	"Spruce", "Poplar", "Ash", "Hickory", "Magnolia", "Holly", "Laurel",
	"Vine", "Garden", "Meadow", "Brook", "Valley", "Ridge", "Crescent",
}

var streetSuffixes = []string{
	"Street", "Avenue", "Boulevard", "Drive", "Lane", "Road", "Court",
	"Place", "Way", "Circle", "Trail", "Parkway", "Highway", "Terrace",
	"Plaza", "Square", "Walk", "Alley",
}

var cities = []string{
	"Springfield", "Riverside", "Fairview", "Franklin", "Greenville", "Bristol",
	"Clinton", "Fairfield", "Madison", "Georgetown", "Salem", "Marion",
	"Ashland", "Burlington", "Manchester", "Oxford", "Hamilton", "Arlington",
	"Centerville", "Dover", "Essex", "Fulton", "Greensboro", "Hartford",
	"Irvington", "Jefferson City", "Kingston", "Lebanon", "Milton", "Newport",
	"Oakland", "Plymouth", "Rochester", "Sterling", "Trenton", "Union City",
	"Victoria", "Watertown", "Yorktown", "Auburn", "Bedford", "Carlisle",
	"Dalton", "Easton", "Farmington", "Glenwood", "Harper", "Indianola",
	"Jamestown", "Knoxville", "Lakeside", "Middleton", "Newton", "Oakwood",
	"Pinecrest", "Quincy", "Redwood", "Somerville", "Troy", "Utica",
	"Verona", "Westfield", "Yorkville", "Belmont", "Cambridge", "Danville",
}

var usStates = []struct {
	Name string
	Abbr string
}{
	{"Alabama", "AL"}, {"Alaska", "AK"}, {"Arizona", "AZ"}, {"Arkansas", "AR"},
	{"California", "CA"}, {"Colorado", "CO"}, {"Connecticut", "CT"}, {"Delaware", "DE"},
	{"Florida", "FL"}, {"Georgia", "GA"}, {"Hawaii", "HI"}, {"Idaho", "ID"},
	{"Illinois", "IL"}, {"Indiana", "IN"}, {"Iowa", "IA"}, {"Kansas", "KS"},
	{"Kentucky", "KY"}, {"Louisiana", "LA"}, {"Maine", "ME"}, {"Maryland", "MD"},
	{"Massachusetts", "MA"}, {"Michigan", "MI"}, {"Minnesota", "MN"}, {"Mississippi", "MS"},
	{"Missouri", "MO"}, {"Montana", "MT"}, {"Nebraska", "NE"}, {"Nevada", "NV"},
	{"New Hampshire", "NH"}, {"New Jersey", "NJ"}, {"New Mexico", "NM"}, {"New York", "NY"},
	{"North Carolina", "NC"}, {"North Dakota", "ND"}, {"Ohio", "OH"}, {"Oklahoma", "OK"},
	{"Oregon", "OR"}, {"Pennsylvania", "PA"}, {"Rhode Island", "RI"}, {"South Carolina", "SC"},
	{"South Dakota", "SD"}, {"Tennessee", "TN"}, {"Texas", "TX"}, {"Utah", "UT"},
	{"Vermont", "VT"}, {"Virginia", "VA"}, {"Washington", "WA"}, {"West Virginia", "WV"},
	{"Wisconsin", "WI"}, {"Wyoming", "WY"},
}

var countries = []string{
	"United States", "Canada", "United Kingdom", "Australia", "Germany",
	"France", "Japan", "Brazil", "India", "Italy", "Spain", "Mexico",
	"South Korea", "Netherlands", "Sweden", "Norway", "Denmark", "Finland",
	"Poland", "Belgium", "Austria", "Switzerland", "Portugal", "Greece",
	"Ireland", "New Zealand", "Singapore", "Israel", "Argentina", "Chile",
	"South Africa", "Turkey", "Egypt", "Thailand", "Vietnam", "Indonesia",
	"Philippines", "Malaysia", "Czech Republic", "Hungary", "Romania",
	"Ukraine", "Russia", "Norway", "Colombia", "Peru", "Nigeria", "Kenya",
}

var companyPrefixes = []string{
	"Acme", "Global", "Prime", "Apex", "Vertex", "Nexus", "Quantum", "Stellar",
	"Cyber", "Meta", "Ultra", "Hyper", "Omni", "Pinnacle", "Summit", "Zenith",
	"Catalyst", "Vanguard", "Pioneer", "Frontier", "Horizon", "Meridian",
	"Atlas", "Titan", "Nova", "Pulse", "Forge", "Crest", "Beacon", "Anchor",
	"Harbor", "Bridge", "Tower", "Crown", "Shield", "Spark", "Flux", "Echo",
	"Drift", "Surge", "Bolt", "Wave", "Core", "Edge", "Peak", "Vault",
}

var companySuffixes = []string{
	"Inc", "LLC", "Corp", "Co", "Ltd", "Group", "Holdings", "Partners",
	"Solutions", "Systems", "Technologies", "Labs", "Works", "Industries",
	"Enterprises", "Ventures", "Dynamics", "Networks", "Software", "Digital",
}

var companyTypes = []string{
	"Software", "Technology", "Consulting", "Logistics", "Manufacturing",
	"Financial", "Healthcare", "Energy", "Media", "Retail", "Biotech",
	"Robotics", "Aerospace", "Automotive", "Pharmaceutical", "Real Estate",
	"Hospitality", "Construction", "Telecommunications", "Entertainment",
}

var buzzwords = []string{
	"synergy", "paradigm", "disruptive", "innovative", "scalable", "agile",
	"leverage", "ecosystem", "blockchain", "AI-driven", "cloud-native",
	"microservices", "digital transformation", "data-driven", "customer-centric",
	"omnichannel", "frictionless", "seamless", "holistic", "turnkey",
	"best-of-breed", "world-class", "cutting-edge", "next-generation",
	"mission-critical", "value-added", "results-oriented", "game-changing",
}

var catchPhraseAdjectives = []string{
	"Adaptive", "Advanced", "Ambitious", "Automated", "Balanced", "Business-focused",
	"Centralized", "Cloned", "Compatible", "Configurable", "Cross-group", "Cross-platform",
	"Customer-focused", "Customizable", "Decentralized", "De-engineered", "Devolved",
	"Digitized", "Distributed", "Diverse", "Down-sized", "Enhanced", "Enterprise-wide",
	"Ergonomic", "Exclusive", "Expanded", "Extended", "Face-to-face", "Focused",
	"Front-line", "Fully-configured", "Function-based", "Fundamental", "Future-proofed",
	"Grass-roots", "Horizontal", "Implemented", "Innovative", "Integrated",
	"Intuitive", "Inverse", "Managed", "Mandatory", "Monitored", "Multi-channelled",
	"Multi-lateral", "Multi-layered", "Networked", "Object-based", "Open-architected",
	"Open-source", "Operative", "Optimized", "Optional", "Organic", "Organized",
	"Persevering", "Persistent", "Phased", "Polarised", "Pre-emptive", "Proactive",
	"Profound", "Programmable", "Progressive", "Public-key", "Quality-focused",
	"Reactive", "Realigned", "Reduced", "Reverse-engineered", "Right-sized",
	"Robust", "Seamless", "Secured", "Self-enabling", "Sharable", "Stand-alone",
	"Streamlined", "Switchable", "Synchronised", "Synergistic", "Synergized",
	"Team-oriented", "Total", "Triple-buffered", "Universal", "Up-sized",
	"Upgradable", "User-centric", "User-friendly", "Versatile", "Virtual",
	"Visionary", "Vision-oriented",
}

var catchPhraseSubjects = []string{
	"24/7 solution", "analyzing", "asymmetric", "asynchronous", "attitude-oriented",
	"background", "bandwidth-monitored", "bi-directional", "bifurcated",
	"bottom-line", "clear-thinking", "client-driven", "client-server", "coherent",
	"cohesive", "composite", "context-sensitive", "contextually-based", "content-based",
	"dedicated", "demand-driven", "didactic", "directional", "discrete",
	"disintermediate", "dynamic", "eco-centric", "empowering", "encompassing",
	"even-keeled", "executive", "explicit", "exuding", "fault-tolerant", "foreground",
	"forward-thinking", "fresh-thinking", "full-range", "global", "grid-enabled",
	"heuristic", "high-level", "holistic", "homogeneous", "human-resource",
	"hybrid", "impactful", "incremental", "intangible", "interactive",
	"intermediate", "leading edge", "local", "logistical", "maximized",
	"methodical", "mission-critical", "mobile", "modular", "motivating",
	"multimedia", "multi-state", "multi-tasking", "national", "needs-based",
	"neutral", "next-generation", "non-volatile", "object-oriented", "optimal",
	"optimizing", "radical", "real-time", "reciprocal", "regional", "responsive",
	"scalable", "secondary", "solution-oriented", "stable", "static",
	"systematic", "systemic", "tangible", "tertiary", "transitional",
	"uniform", "upward-trending", "user-facing", "value-added", "web-enabled",
	"well-modulated", "zero-administration", "zero-defect", "zero-tolerance",
}

var catchPhraseVerbs = []string{
	"actualize", "administer", "aggregate", "architect", "benchmark", "brand",
	"build", "catalog", "cultivate", "deliver", "deploy", "develop", "disintermediate",
	"drive", "e-enable", "embrace", "empower", "enable", "engage", "engineer",
	"enhance", "envisioneer", "evisculate", "evolve", "expedite", "exploit",
	"extend", "facilitate", "fashion", "forge", "formulate", "foster", "frame",
	"generate", "grow", "harness", "impact", "implement", "incentivize",
	"incubate", "initiate", "innovate", "integrate", "iterate", "leverage",
	"maximize", "mesh", "monetize", "morph", "optimize", "orchestrate",
	"plagiarize", "productize", "promote", "provide", "pursue", "redefine",
	"re-engineer", "reinvent", "repurpose", "revolutionize", "scale", "seize",
	"seamlessly", "simplify", "strategize", "streamline", "syndicate",
	"synergize", "synthesize", "target", "transform", "transition", "unleash",
	"utilize", "visualize", "whiteboard",
}

var catchPhraseObjects = []string{
	"action-items", "applications", "architectures", "arrays", "bandwidth",
	"blockchains", "communities", "content", "convergence", "core competencies",
	"customer service", "data", "deliverables", "e-business", "e-commerce",
	"e-markets", "e-services", "e-tailers", "experiences", "eyeballs",
	"functionalities", "growth strategies", "infomediaries", "information",
	"initiatives", "interfaces", "internet solutions", "intranets", "knowledge",
	"local area networks", "markets", "methodologies", "metrics", "models",
	"networks", "niches", "paradigms", "partnerships", "portals", "products",
	"quality vectors", "relationships", "ROI", "schemas", "solutions",
	"supply-chains", "synergies", "systems", "technologies", "users",
	"value", "vortals", "web services", "web-readiness",
}

var productAdjectives = []string{
	"Small", "Ergonomic", "Rustic", "Intelligent", "Gorgeous", "Incredible",
	"Fantastic", "Practical", "Sleek", "Awesome", "Generic", "Handcrafted",
	"Handmade", "Licensed", "Refined", "Unbranded", "Tasty", "Modern",
	"Electronic", "Refurbished", "Compact", "Durable", "Wireless", "Portable",
}

var productNouns = []string{
	"Chair", "Car", "Computer", "Gloves", "Pants", "Shirt", "Table", "Shoes",
	"Hat", "Plate", "Knife", "Bottle", "Lamp", "Keyboard", "Bag", "Mouse",
	"Wallet", "Phone", "Monitor", "Headphones", "Bike", "Soccer Ball",
	"Towels", "Soap", "Tuna", "Chicken", "Cheese", "Bacon", "Mouse Pad",
	"Mouse Pad", "Pants", "Gloves", "Table", "Keyboard",
}

var productCategories = []string{
	"Electronics", "Books", "Clothing", "Home & Garden", "Sports & Outdoors",
	"Toys & Games", "Health & Beauty", "Automotive", "Tools & Hardware",
	"Pet Supplies", "Office Products", "Food & Beverage", "Baby Products",
	"Jewelry", "Music", "Movies", "Video Games", "Software", "Art & Crafts",
	"Industrial & Scientific",
}

var mimeTypes = []string{
	"text/plain", "text/html", "text/css", "text/javascript", "application/json",
	"application/xml", "application/pdf", "application/zip", "application/x-tar",
	"application/x-gzip", "application/octet-stream", "application/x-www-form-urlencoded",
	"application/x-sh", "application/x-httpd-php", "application/rtf",
	"image/jpeg", "image/png", "image/gif", "image/svg+xml", "image/webp",
	"image/tiff", "image/bmp", "image/x-icon",
	"audio/mpeg", "audio/ogg", "audio/wav", "audio/webm",
	"video/mp4", "video/webm", "video/ogg", "video/x-msvideo",
	"font/woff", "font/woff2", "font/ttf", "font/otf",
	"multipart/form-data", "multipart/byteranges",
	"message/rfc822",
	"model/gltf-binary", "model/obj",
}

var fileExtensions = []string{
	"txt", "html", "css", "js", "json", "xml", "pdf", "zip", "tar", "gz",
	"bin", "sh", "php", "rtf", "jpg", "jpeg", "png", "gif", "svg", "webp",
	"tiff", "bmp", "ico", "mp3", "ogg", "wav", "mp4", "webm", "avi",
	"woff", "woff2", "ttf", "otf", "csv", "yaml", "yml", "toml", "md",
	"go", "py", "rb", "java", "c", "cpp", "rs", "ts", "sql", "log",
}

var userAgentStrings = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Safari/605.1.15",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (iPad; CPU OS 17_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.2 Mobile/15E148 Safari/604.1",
	"Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 OPR/106.0.0.0",
}

var tlds = []string{
	"com", "org", "net", "io", "co", "dev", "app", "tech", "cloud", "ai",
	"info", "biz", "online", "store", "blog", "design", "digital", "systems",
}

var filePathDirs = []string{
	"home", "var", "etc", "usr", "opt", "tmp", "srv", "lib", "bin", "mnt",
	"documents", "downloads", "pictures", "music", "videos", "projects",
	"config", "cache", "data", "logs",
}

var filePathNames = []string{
	"config", "data", "index", "main", "utils", "helper", "handler", "model",
	"view", "controller", "service", "client", "server", "test", "readme",
	"license", "makefile", "dockerfile", "manifest", "schema",
}

var creditCardPrefixes = map[string][]string{
	"visa":       {"4"},
	"mastercard": {"51", "52", "53", "54", "55"},
	"amex":       {"34", "37"},
	"discover":   {"6011", "65"},
}

var ibanCountries = []struct {
	Country string
	Code    string
	Length  int
}{
	{"Germany", "DE", 22}, {"France", "FR", 27}, {"Italy", "IT", 27},
	{"Spain", "ES", 24}, {"Netherlands", "NL", 18}, {"Belgium", "BE", 16},
	{"Austria", "AT", 20}, {"Poland", "PL", 28}, {"Portugal", "PT", 25},
	{"Ireland", "IE", 22}, {"Finland", "FI", 18}, {"Greece", "GR", 27},
}
