package main

import (
	"flag"
	"strings"
)

type strSliceFlag []string

func (a *strSliceFlag) Set(value string) error {
	if elems := strings.Split(value, " "); len(elems) > 0 {
		*a = append(*a, elems...)
		return nil
	}

	*a = append(*a, value)
	return nil
}

func (a *strSliceFlag) String() string {
	return strings.Join(*a, ",")
}

func flagStringSlice(name string, value []string, help string) *strSliceFlag {
	data := make(strSliceFlag, len(value))
	for i, v := range value {
		data[i] = v
	}

	flag.Var(&data, name, help)

	return &data
}