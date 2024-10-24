
## Serve Static Files
- Can use prefix "/web" in this case.
- http://localhost:8080/web
```
func main() {
    staticDir := "./static"
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/web/", http.StripPrefix("/web/", fs))
	http.ListenAndServe(":8080", nil)
}
```

## Serve ng without route
- Must use "/" path. This path "/webapp" doesn't work.
- http://localhost:8080/

```go
func main() {
    distDir := "./ng-without-route"
	fs := http.FileServer(http.Dir(distDir))
	r.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}

```

## Serve ng with route
- Must use "/" path. This path "/webapp" doesn't work.
- http://localhost:8080/

```go
func main() {
    distDir := "./ngapp"
	indexFile := filepath.Join(distDir, "index.html")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(distDir, r.URL.Path)
		if _, err := os.Stat(requestedPath); os.IsNotExist(err) {
			http.ServeFile(w, r, indexFile)
		} else {
			http.ServeFile(w, r, requestedPath)
		}
	})
	
	http.ListenAndServe(":8080", nil)
}
```

Can also use http.ServeMux like following code:
```go
func main() {
    distDir := "./ngapp"
	indexFile := filepath.Join(distDir, "index.html")

    mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedPath := filepath.Join(distDir, r.URL.Path)
		if _, err := os.Stat(requestedPath); os.IsNotExist(err) {
			http.ServeFile(w, r, indexFile)
		} else {
			http.ServeFile(w, r, requestedPath)
		}
	})
	
	http.ListenAndServe(":8080", mux)
}
```