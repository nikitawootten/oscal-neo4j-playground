package generator

import "strings"

const prefix = "oscal-complete-oscal-"

func stripPrefix(name string) string {
	return strings.Replace(name, prefix, "", 1)
}

func golangifyName(name string) string {
	name = strings.ReplaceAll(name, "_", "-")
	name = strings.ReplaceAll(name, ":", "-")
	clauses := strings.Split(name, "-")
	for i := range clauses {
		if len(clauses[i]) > 0 {
			clauses[i] = strings.Title(clauses[i])
		}
	}
	return strings.Join(clauses, "")
}

func golangifyFieldRef(ref string) string {
	ref = strings.Replace(ref, "#field_oscal-", "", 1)
	ref = strings.Replace(ref, "#assembly_oscal-", "", 1)
	return golangifyName(ref)
}
