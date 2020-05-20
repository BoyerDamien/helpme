package main

func Find(slice []interface{}, val string) bool {
	for _, item := range slice {
		if item.(string) == val {
			return true
		}
	}
	return false
}
