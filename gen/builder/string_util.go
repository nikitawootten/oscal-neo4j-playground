package builder

import "strings"

const prefix = "oscal-complete-oscal-"

func stripPrefix(name string) string {
	return strings.Replace(name, prefix, "", 1)
}

func golangifyName(name string) string {
	name = strings.ReplaceAll(name, " ", "")
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
