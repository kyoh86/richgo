package test

import (
	"strings"

	"github.com/kyoh86/richgo/config"
)

type labelSet struct {
	build  string
	start  string
	pass   string
	fail   string
	skip   string
	cover  string
	anonym string
}

var labelSets = map[config.LabelType]labelSet{
	config.LabelTypeLong:  newLabelSet("| ", "BUILD", "START", "PASS", "FAIL", "SKIP", "COVER"),
	config.LabelTypeShort: newLabelSet(" ", "!!", ">", "o", "x", "-", "%"),
	config.LabelTypeNone:  newLabelSet("", "", "", "", "", "", ""),
}

func labels() labelSet {
	return labelSets[*config.C.LabelType]
}

func maxInt(numbers ...int) int {
	m := 0
	for _, n := range numbers {
		if m < n {
			m = n
		}
	}
	return m
}

func newLabelSet(suffix, build, start, pass, fail, skip, cover string) labelSet {
	max := maxInt(
		len(build),
		len(start),
		len(pass),
		len(fail),
		len(skip),
		len(cover),
	)

	anonym := strings.Repeat(" ", max)
	return labelSet{
		build:  string((build + anonym)[:max]) + suffix,
		start:  string((start + anonym)[:max]) + suffix,
		pass:   string((pass + anonym)[:max]) + suffix,
		fail:   string((fail + anonym)[:max]) + suffix,
		skip:   string((skip + anonym)[:max]) + suffix,
		cover:  string((cover + anonym)[:max]) + suffix,
		anonym: anonym + suffix,
	}
}

func (s labelSet) Build() string {
	return s.build
}
func (s labelSet) Start() string {
	return s.start
}
func (s labelSet) Pass() string {
	return s.pass
}
func (s labelSet) Fail() string {
	return s.fail
}
func (s labelSet) Skip() string {
	return s.skip
}
func (s labelSet) Cover() string {
	return s.cover
}
func (s labelSet) Anonym() string {
	return s.anonym
}
