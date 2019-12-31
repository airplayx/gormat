package Sql2struct

func InStringSlice(f string, a []string) bool {
	for _, s := range a {
		if f == s {
			return true
		}
	}
	return false
}
