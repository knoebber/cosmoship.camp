package main

import (
	"fmt"
	"net/http"
	"strings"
)

func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, isAdmin := checkSession(w, r)
		if id < 1 {
			return
		}

		path := r.URL.String()
		if !isAdmin && !checkMemberPermission(path, id) {
			permissionDenied(w, fmt.Errorf("member %d attempted to visit %s", id, path))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkMemberPermission(path string, memberID int) bool {
	return strings.HasPrefix(path, fmt.Sprintf("%s/members/%d", apiPath, memberID))
}
