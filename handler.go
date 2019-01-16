package gophercise2

import "net/http"

// MapHandler will return http.HandlerFunc
func MapHandler(urlMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := urlMap[r.URL.Path]
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will return http.HandlerFunc
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	return nil, nil
}
