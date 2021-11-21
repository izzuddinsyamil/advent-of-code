package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	file, err := os.Open("2020/day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []string
	var d string
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			data = append(data, strings.TrimSpace(d))
			d = ""
		} else {
			d += " " + t
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	data = append(data, strings.TrimSpace(d))
	p := convertInput(data)
	fmt.Println(validatePassport(p))
}

func validatePassport(passport []Passport) (res int) {
	for _, p := range passport {
		if validate(p) {
			res += 1
		}
	}
	return res
}

func validate(p Passport) bool {
	v := reflect.ValueOf(p)
	types := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" && types.Field(i).Name != "CountryID" {
			return false
		}
	}

	return true
}

type Passport struct {
	BirthYear  string
	IssueYear  string
	ExpYear    string
	Height     string
	HairColor  string
	EyeColor   string
	PassportID string
	CountryID  string
}

func convertInput(inp []string) (res []Passport) {
	for _, i := range inp {
		res = append(res, convertInputToPassport(i))
	}
	return
}

func convertInputToPassport(inp string) (res Passport) {
	inp_split := strings.Split(inp, " ")
	for _, i := range inp_split {
		spl := strings.Split(i, ":")
		key, val := spl[0], spl[1]

		switch key {
		case "byr":
			res.BirthYear = val
		case "iyr":
			res.IssueYear = val
		case "eyr":
			res.ExpYear = val
		case "hgt":
			res.Height = val
		case "hcl":
			res.HairColor = val
		case "ecl":
			res.EyeColor = val
		case "pid":
			res.PassportID = val
		case "cid":
			res.CountryID = val
		}
	}

	return
}
