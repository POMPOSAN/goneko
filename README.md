# GoNeko: A golang scraping package for Nyaa
[![Go Report Card](https://goreportcard.com/badge/github.com/Yanisssssse/goneko)](https://goreportcard.com/report/github.com/Yanisssssse/goneko)

GoNeko is a simple, lightweight and ultrafast scraper for nyaa.si  
You can search torrents just like on the website, and get the result as golang structs.

## Installation
```go get github.com/POMPOSAN/goneko```

## Documentation

### - Parse(url string)
return an array containing the nyaa entrys of the specified url.
### - HomePage()
return an array containing the nyaa entrys of the home page.
### - Search(opts *Opts)
Allow you to make a precise search  
- Query (string) : The text of your search.
-  Page (int) : The page number
- User (string) : Show only entrys posted by the specified user 
- Filter (int) :
    - 0 - No Filter
    - 1 - No Remakes
    - 2 - Trusted Only
  
- Cat (int) & Subcat (int) :
	- 0 - All Categories (& Subcategories)

 	- 1.0 - Anime 
        - 1.1 - Anime Music Video
		- 1.2 - English-translated
	 	- 1.3 - Non-English-translated
		- 1.4 - Raw

	- 2.0 - Audio
	    - 2.1 - Lossless
		- 2.2 - Lossy

	- 3.0 - Literature
		- 3.1 - English-translated
		- 3.2 - Non-English-translated
		- 3.3 - Raw
	
	- 4.0 - Live Action
		- 4.1 - English-translated
		- 4.2 - Idol/Promotional Video
		- 4.3 - Non-English-translated
		- 4.4 - Raw

	- 5.0 - Pictures
		- 5.1 - Graphics
		- 5.2 - Photos

	- 6.0 - Software
		- 6.1 - Applications
		- 6.2 - Games
