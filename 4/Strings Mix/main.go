package main

import (
	"sort"
	"strings"
)

func letters(str string, f string) (r []Result) {
	for _, a := range str {
		if !strings.Contains("abcdefghijklmnopqrstuvwxyz", string(a)) {
			continue
		}
		g := strings.Count(str, string(a))
		for _, b := range r {
			if a == b.letter {
				goto skip
			}
		}
		if g > 1 {
			r = append(r, Result{from: f, letter: a, count: g})
		}
	skip:
	}
	return
}

type Result struct {
	from   string
	count  int
	letter rune
}

func Mix(s1, s2 string) string {
	slice := letters(s1, "1")
	for _, a := range letters(s2, "2") {
		for i, b := range slice {
			if b.letter == a.letter {
				if b.count == a.count {
					slice[i].from = "="
				}
				if b.count < a.count {
					slice[i].from = "2"
					slice[i].count = a.count
				}
				goto skip
			}
		}
		slice = append(slice, a)
	skip:
	}
	sort.Slice(slice, func(i, j int) bool {
		if slice[i].count > slice[j].count {
			return true
		}
		if slice[i].count == slice[j].count {
			a, b := strings.Index("12=", slice[i].from), strings.Index("12=", slice[j].from)
			if a < b {
				return true
			} else if a == b {
				if slice[i].letter < slice[j].letter {
					return true
				}
			}

		}
		return false
	})
	var ret []string
	for _, o := range slice {
		ret = append(ret, o.from+":"+strings.Repeat(string(o.letter), o.count))
	}
	return strings.Join(ret, "/")
}

func main() {
	println(Mix("looping is fun but dangerous", "less dangerous than coding"))
	println("1:ooo/1:uuu/2:sss/=:nnn/1:ii/2:aa/2:dd/2:ee/=:gg")
}
