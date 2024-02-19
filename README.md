# Movie List Web App

This is a simple web application written in Go that allows users to view a list of movies and add new movies.

## Requirements
- Go 1.16 or higher installed on your machine.
To install go, click here https://go.dev/doc/install

## How to run

1. Clone this repository:
```bash
    git clone https://github.com/rohansaibuddhi/movie-list
```
2. Navigate to the project directory:
```bash
    cd movie-list
```
3. Run the Go server
```bash
    go run main.go
```
4. Open the webpage

    Go to http://localhost:8080 to view and try the Movie List webapp

## Usage
- To view the movies, visit the homepage
- To add a new movie, fill in the title and director in the form on the right side and click submit to add to the list. 

## Learnings

### 1. Go as a Backend

I had used Go prior to this during my internship at Appointy, so this served as a refresher to re-learning Go. It refreshed my memory on HTTP handlers for different requests using the _http.handleFunc()_

### 2. Template Rendering
This was my first time using Templates in Go to render dynamic data. I used _template.Must()_ to parse HTMl templates from files and then executed them with the data to generate dynamic content forthe webpage.

### 3. Playing with HTMX
I wanted to see what all the buzz of HTMX is about and it sure seems good so far. I used HTMX to make asynchronous requests with AJAX to update the page without needing to reload the entire page, ensuring smoother user experience. 

## Future Work
I intend to extend this project by learning how to connect Go to a database(relational in this case - MySQL) and having these dynamic changes stay persistent through the way of using the DB to read and store data!