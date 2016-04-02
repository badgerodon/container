package main

import (
	"flag"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/chuckpreslar/inflect"
)

func main() {
	log.SetFlags(0)

	var (
		template    string
		destination string
		types       string
		packageName string
	)
	flag.StringVar(&template, "template", "", "the template to use to generate code")
	flag.StringVar(&destination, "destination", "", "where to save the file")
	flag.StringVar(&types, "types", "", "the comma-separated types to generate code for")
	flag.StringVar(&packageName, "package", "", "the package name")
	flag.Parse()

	contents := []string{"package " + packageName + "\n"}

	imported := map[string]struct{}{}

	switch template {
	case "arraylist":
		bs, err := ioutil.ReadFile("arraylist.go")
		if err != nil {
			log.Fatalln(err)
		}
		src := string(bs)

		reTypeName := regexp.MustCompile(`ArrayList`)
		reItems := regexp.MustCompile(`interface\{\} /\* ITEM \*/`)
		reEquals := regexp.MustCompile(`t == item`)
		rePackage := regexp.MustCompile(`package container`)

		for _, typ := range expandTypes(strings.Split(types, ",")) {
			dst := src
			dst = rePackage.ReplaceAllString(dst, "\n")
			dst = reTypeName.ReplaceAllString(dst, "ArrayListOf"+typeName(typ))
			dst = reItems.ReplaceAllString(dst, typ)
			if typ == "[]byte" {
				imported["bytes"] = struct{}{}
				dst = reEquals.ReplaceAllString(dst, "bytes.Equal(t, item)")
			}

			contents = append(contents, dst)
		}
	case "unrolledlist":
		bs, err := ioutil.ReadFile("unrolledlist.go")
		if err != nil {
			log.Fatalln(err)
		}
		src := string(bs)

		reTypeName := regexp.MustCompile(`UnrolledList`)
		reItems := regexp.MustCompile(`interface\{\} /\* ITEM \*/`)
		reEquals := regexp.MustCompile(`t == item`)
		rePackage := regexp.MustCompile(`package container`)

		for _, typ := range expandTypes(strings.Split(types, ",")) {
			dst := src
			dst = rePackage.ReplaceAllString(dst, "\n")
			dst = reTypeName.ReplaceAllString(dst, "UnrolledListOf"+typeName(typ))
			dst = reItems.ReplaceAllString(dst, typ)
			if typ == "[]byte" {
				imported["bytes"] = struct{}{}
				dst = reEquals.ReplaceAllString(dst, "bytes.Equal(t, item)")
			}

			contents = append(contents, dst)
		}
	}

	for i := range imported {
		contents[0] += "import \"" + i + "\"\n"
	}

	err := ioutil.WriteFile(destination, []byte(strings.Join(contents, "")), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func typeName(typ string) string {
	switch {
	case typ == "[]byte":
		return "Bytes"
	case strings.HasPrefix(typ, "[]"):
		return inflect.Pluralize(typeName(typ[2:]))
	default:
		return inflect.UpperCamelCase(typ)
	}
}

func expandTypes(types []string) []string {
	var expanded []string
	for _, typ := range types {
		switch typ {
		case "NUMERIC":
			expanded = append(expanded,
				"uint8", "uint16", "uint32", "uint64",
				"int8", "int16", "int32", "int64",
				"float32", "float64",
				"complex64", "complex128",
				"byte", "rune",
				"uint", "int", "uintptr")
		default:
			expanded = append(expanded, typ)
		}
	}
	return expanded
}
