package main

import (
    "html/template"
    "net/http"
    "strconv" // Import strconv package for conversion
)

func main() {
    // Serve static CSS file
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Serve the HTML front end
    http.HandleFunc("/", serveHTML)

    // Start the web server on port 8080
    http.ListenAndServe(":8080", nil)
}

// serveHTML serves the HTML content with images
func serveHTML(w http.ResponseWriter, r *http.Request) {
    images := generateImageURLs(100) // Generate URLs for 12 images
    tmpl := template.Must(template.ParseFiles("templates/index.html"))

    // Pass images slice to the template
    tmpl.ExecuteTemplate(w, "index.html", images)
}

// generateImageURLs returns a slice of image URLs
func generateImageURLs(count int) []string {
    urls := make([]string, count)
    for i := 0; i < count; i++ {
        urls[i] = "https://picsum.photos/200/300?random=" + strconv.Itoa(i)
    }
    return urls
}
