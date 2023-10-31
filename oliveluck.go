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

var (
	randSource rand.Source
	rng        *rand.Rand
	funcSlice  = []func() string{}
)

func cleanAndCombine(f1, f2 func() string) string {
	return clean(f1(), f2())
}

func clean(str1, str2 string) string {
	r := strings.ToLower(str2 + str1)
	str := regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(r, "")

	return str
}

func getColor() string {
	color := "blue"
	for {
		newColor := gofakeit.SafeColor()
		if color != "black" {
			color = newColor
			break
		}
	}

	return color
}

func init() {
	randSource = rand.NewSource(time.Now().UnixNano())
	rng = rand.New(randSource)

	funcSlice = []func() string{
		func() string {
			color := getColor()
			return clean(gofakeit.NounAbstract(), color)
		},
		func() string {
			str1, err := sententia.Make("{{ noun }}")
			if err != nil {
				panic(err)
			}

			str2, err := sententia.Make("{{ adjective }}")
			if err != nil {
				panic(err)
			}
			return clean(str1, str2)
		},
		func() string { return cleanAndCombine(gofakeit.Animal, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.BeerName, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.BeerStyle, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.BeerYeast, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.BuzzWord, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.CarMaker, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.CarModel, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.CarType, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Fruit, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Gender, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.HackerAdjective, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.HackerNoun, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.HackerNoun, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.HackerVerb, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Hobby, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Hobby, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.JobLevel, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.JobTitle, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.JobTitle, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Language, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.NounAbstract, gofakeit.HackerAdjective) },
		func() string { return cleanAndCombine(gofakeit.State, gofakeit.Adjective) },
		func() string { return cleanAndCombine(gofakeit.Vegetable, gofakeit.Adjective) },
		func() string { return cleanAndCombine(randomdata.Noun, randomdata.Adjective) },
	}
}

func Main() int {
	test1()

	return 0
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

func GenRandomNames(namer func() string, maxNames int) []string {
	seen := make(map[string]string)
	names := make([]string, 0, maxNames)

	for count := 0; count < maxNames; {
		name := namer()
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

func GetRandNamer() func() string {
	rand := rng.Intn(len(funcSlice))

	return funcSlice[rand]
}
