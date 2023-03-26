# URL-Shortener
This project is a simple URL shortener in Golang that uses a base62 encoding logic to generate short URLs.

 - The generateShortURL function generates a 6-character random string using base62 encoding.
  
 - The handler function accepts a longURL parameter, generates a short URL, and stores the mapping in the shortURLs map. It then responds with the generated short URL.
 - The redirectHandler function handles requests for short URLs. It extracts the short URL from the request path, looks up the corresponding long URL in the shortURLs map, and redirects the user to the long URL.
  

To use this URL shortener, you can run the program and visit 
http://localhost:8080/shorten?url=<your-long-url> to generate a short URL, and visit http://localhost:8080/<your-short-url> to be redirected to the long URL.