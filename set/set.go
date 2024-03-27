package set

import "fmt"

type Set map[string]struct{}

func (s Set) Remove(str string) {
	if _, ok := s[str]; ok {
		delete(s, str)
	}
}

func (s Set) String() string {
	printString := "("
	for k, _ := range s {
		printString += fmt.Sprintf("%s, ", k)
	}
	printString = printString[:len(printString)-2] + ")"
	return printString
}

func (s Set) Has(str string) bool {
	_, ok := s[str]
	return ok
}

func (s Set) Add(str string) {
	if _, ok := s[str]; !ok {
		s[str] = struct{}{}
	}
}

func (s Set) ExtendFromSlice(slice []string) {
	for _, i := range slice {
		s.Add(i)
	}
}

func (s Set) Pop() string {
	var key string
	for key, _ = range s {
		break
	}
	s.Remove(key)
	return key
}

func (s Set) Disjoint(s2 Set) Set {
	disjointSet := Set{}

	f1 := func(set1, set2 Set) {
		for k, _ := range set1 {
			if _, ok := set2[k]; !ok {
				disjointSet.Add(k)
			}
		}
	}

	f1(s2, s)

	return disjointSet
}

func (s Set) merge(s2 Set) {
	for k, v := range s2 {
		s[k] = v
	}
}

func (s Set) ToSlice() []string {

	var stringArr []string

	for k, _ := range s {
		stringArr = append(stringArr, k)
	}

	return stringArr
}
