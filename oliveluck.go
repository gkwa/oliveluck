package oliveluck

import (
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/castillobgr/sententia"
)

func Main() int {
	slog.Debug("oliveluck", "test", true)

	test1()
	// test2()

	return 0
}

func test2() {
	namer := GetRandNamer()
	names := GenRandomNames(namer, 1)
	slog.Debug("debug", "name", names[0])
}

func test1() {
	i := 0
	for i < 10 {
		namer := GetRandNamer()
		names := GenRandomNames(namer, 1)
		for _, name := range names {
			slog.Debug("debug", "name", name)
			fmt.Fprintf(os.Stdout, "%s\n", name)
		}
		i++
	}
}

var (
	randSource rand.Source
	rng        *rand.Rand

	Namers = []Namer{
		&SententiaPathNamer{},
		&RandomdataPathNamer{},
		&GofakeitPathNamer{},
		&Combo1{},
		&Combo2{},
		&Combo4{},
		&Combo5{},
		&Combo6{},
		&Combo7{},
		&Combo8{},
	}
)

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
	rng = rand.New(randSource)
}

func GenRandomNames(namer Namer, maxNames int) []string {
	seen := make(map[string]string)
	names := make([]string, 0, maxNames)

	for count := 0; count < maxNames; {
		name := namer.GetName()
		_, found := seen[name]

		if found {
			continue
		}

		names = append(names, name)

		count++
		seen[name] = name
	}

	return names
}

type Combo1 struct{}

func (spn *Combo1) GetName() string {
	noun := gofakeit.NounAbstract()

	color := "blue"
	for {
		color := gofakeit.SafeColor()
		if color != "black" {
			break
		}
	}

	adjective := color

	return clean(noun, adjective)
}

type Combo2 struct{}

func (spn *Combo2) GetName() string {
	noun := gofakeit.State()
	adjective := gofakeit.Adjective()

	return clean(noun, adjective)
}

type Combo3 struct{}

func (spn *Combo3) GetName() string {
	noun := gofakeit.Hobby()
	adjective := gofakeit.Adjective()

	return clean(noun, adjective)
}

type Combo4 struct{}

func (spn *Combo4) GetName() string {
	noun := gofakeit.BeerName()
	adjective := gofakeit.Adjective()

	return clean(noun, adjective)
}

type Combo5 struct{}

func (spn *Combo5) GetName() string {
	noun := gofakeit.CarMaker()
	adjective := gofakeit.Adjective()

	return clean(noun, adjective)
}

type Combo7 struct{}

func (spn *Combo7) GetName() string {
	noun := gofakeit.HackerNoun()
	adjective := gofakeit.Adjective()

	return clean(noun, adjective)
}

type Combo8 struct{}

func (spn *Combo8) GetName() string {
	noun := gofakeit.Animal()
	adjective := gofakeit.Adjective()

	return clean(noun, adjective)
}

type Combo6 struct{}

func (spn *Combo6) GetName() string {
	noun := gofakeit.JobTitle()
	adjective := gofakeit.Adjective()

	return clean(noun, adjective)
}

func GetRandNamer() Namer {
	rand := rng.Intn(len(Namers))
	pathNamer := Namers[rand]
	return pathNamer
}

type GofakeitPathNamer struct{}

func (spn *GofakeitPathNamer) GetName() string {
	noun := gofakeit.NounAbstract()
	adjective := gofakeit.HackerAdjective()

	return clean(noun, adjective)
}

type SententiaPathNamer struct{}

func (spn *SententiaPathNamer) GetName() string {
	str1, err := sententia.Make("{{ noun }}")
	if err != nil {
		panic(err)
	}

	str2, err := sententia.Make("{{ adjective }}")
	if err != nil {
		panic(err)
	}

	return clean(str1, str2)
}

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
	rng = rand.New(randSource)
}

func clean(str1, str2 string) string {
	r := strings.ToLower(str2 + str1)
	str := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(r, "")
	return str
}

type Namer interface {
	GetName() string
}

type RandomdataPathNamer struct{}

func (rpn *RandomdataPathNamer) GetName() string {
	noun := randomdata.Noun()
	adjective := randomdata.Adjective()

	return clean(noun, adjective)
}
