// This package will contains methods for working with list of strings
package arrstr

// This method check if list of string contains specific string
func ContainsString(s []string, v string) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}
