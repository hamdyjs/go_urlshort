package urlshort

import (
	"net/http"

	yamlV2 "gopkg.in/yaml.v2"
)

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
	parsedYAML, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	urlMap := buildMap(parsedYAML)
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := urlMap[r.URL.Path]
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}, nil
}

func parseYAML(yaml []byte) (result []map[string]string, err error) {
	err = yamlV2.Unmarshal(yaml, &result)
	if err != nil {
		return nil, err
	}
	return
}

func buildMap(yaml []map[string]string) map[string]string {
	urlMap := make(map[string]string)
	for _, entry := range yaml {
		urlMap[entry["path"]] = entry["url"]
	}
	return urlMap
}
