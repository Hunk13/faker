package faker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Faker is struct for Faker
type Faker struct {
	Generator        *rand.Rand
	firstNameMale    []string
	firstNameFemale  []string
	lastName         []string
	maleNameFormat   []string
	femaleNameFormat []string
	usernameFormat   []string
	// Generic top-level domain
	gTLD []string
}

// NewFaker returns a new instance of Faker instance with a random seed
func NewFaker() Faker {
	source := rand.NewSource(time.Now().Unix())

	return Faker{
		Generator:       rand.New(source),
		firstNameMale:   readFromFile("./data/first_name_male.json"),
		firstNameFemale: readFromFile("./data/first_name_female.json"),
		lastName:        readFromFile("./data/last_name.json"),
		maleNameFormat: []string{
			"{{firstNameMale}} {{lastName}}",
			"{{firstNameMale}} {{lastName}}",
			"{{firstNameMale}} {{lastName}}",
			"{{firstNameMale}} {{lastName}}",
			"{{lastName}} {{firstNameMale}}",
		},
		femaleNameFormat: []string{
			"{{firstNameFemale}} {{lastName}}",
			"{{firstNameFemale}} {{lastName}}",
			"{{firstNameFemale}} {{lastName}}",
			"{{firstNameFemale}} {{lastName}}",
			"{{lastName}} {{firstNameFemale}}",
		},
		usernameFormat: []string{
			"{{lastName}}.{{firstName}}",
			"{{firstName}}.{{lastName}}",
			"{{firstName}}",
			"{{lastName}}",
		},
		gTLD: readFromFile("./data/g_tld.json"),
	}
}

// Boolean returns Boolean instance
func (f Faker) Boolean() Boolean {
	return Boolean{&f}
}

// Internet returns Internet instance
func (f Faker) Internet() Internet {
	return Internet{&f}
}

// Person returns Person instance
func (f Faker) Person() Person {
	return Person{&f}
}

// UUID returns UUID instance
func (f Faker) UUID() UUID {
	return UUID{&f}
}

// IntBetween returns a Int between a given minimum and maximum values
func (f Faker) IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	return f.Generator.Intn(diff+1) + min
}

// RandomStringElement returns a random string element from a given list of strings
func (f Faker) RandomStringElement(s []string) string {
	i := f.IntBetween(0, len(s)-1)
	return s[i]
}

// Asciify returns string that replace all "*" characters with random ASCII values from a given string
func (f Faker) Asciify(in string) string {
	var out strings.Builder

	for i := 0; i < len(in); i++ {
		if in[i] == '*' {
			out.WriteString(fmt.Sprintf("%c", f.IntBetween(97, 122)))
		} else {
			out.WriteByte(in[i])
		}
	}

	return out.String()
}

// ByName returns random data by faker
func (f Faker) ByName(faker string) interface{} {
	switch strings.ToLower(faker) {
	// Boolean
	case "boolean":
		return f.Boolean().Boolean()
	// Internet
	case "username":
		return f.Internet().Username()
	case "gtld":
		return f.Internet().GTLD()
	case "domain":
		return f.Internet().Domain()
	case "email":
		return f.Internet().Email()
	// Person
	case "firstname", "person.firstname":
		return f.Person().FirstName()
	case "lastname", "person.lastname":
		return f.Person().LastName()
	case "firstname male", "person.firstnamemale":
		return f.Person().FirstNameMale()
	case "firstname female", "person.firstnamefemale":
		return f.Person().FirstNameFemale()
	case "name", "person.name":
		return f.Person().Name()
	case "name male", "person.namemale":
		return f.Person().NameMale()
	case "name female", "person.namefemale":
		return f.Person().NameFemale()
	case "gender", "person.gender":
		return f.Person().Gender()
	case "gender male", "person.gendermale":
		return f.Person().GenderMale()
	case "gender female", "person.genderfemale":
		return f.Person().GenderFemale()
	// UUID
	case "uuid":
		return f.UUID().V4()
	default:
		return nil
	}
}

func readFromFile(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return make([]string, 0)
	}

	var slice []string
	err = json.Unmarshal(data, &slice)
	if err != nil {
		fmt.Println(err)
		return make([]string, 0)
	}
	return slice
}
