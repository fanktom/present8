package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/russross/blackfriday"
)

var (
	title string
)

func main() {
	// Flags
	flag.StringVar(&title, "title", "present8", "title for the presentation")
	flag.Parse()

	// Read Input
	if len(os.Args) < 2 {
		panic("No input markdown file fiven")
	}
	input := os.Args[len(os.Args)-1]
	markdown, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	// Compile HTML
	output := blackfriday.MarkdownCommon(markdown)

	// Processing Pipeline
	output = compileSlides(output)
	output = sizeImages(output)
	output = surroundWithHTML(output)
	output = numberSlides(output)

	// Write HTML
	err = ioutil.WriteFile(input+".html", output, 0644)
	if err != nil {
		panic(err)
	}
}

// Wraps all slides into the HTML base structure
func surroundWithHTML(html []byte) []byte {
	out := fmt.Sprintf(`<!DOCTYPE html>
<html>
  <head>
    <title>%v</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<style>
			/* minimal style */
			body, h1, h2, h3, h4, h5, p {
				margin: 0;
				padding: 0;
			}

			body {
				background-color: #f6f8fa;
			}

			section.slide {
				position: relative;
				min-height: 100vh;
				display: flex;
				flex-direction: column;
				justify-content: center;
				background-color: white;
				margin-bottom: 20px;
			}

			section.slide > div.padding {
				margin: 2em;
			}

			section.slide > div.pager {
				position: absolute;
				top: 0;
				right: 0;
				margin: 1em;
				color: rgba(0,0,0,0.4);
				font-size: 85%%;
			}

			section.slide div.row {
				display: flex;
			}
			
			section.slide div.column {
				width: 100%%;
			}
			
			/* base style */
			body {
				font-family: sans-serif;
				font-size: 1.5em;
				line-height: 1.5;
			}

			h1, h2, h3, table {
				margin-bottom: 1em;
			}
			
			h4, h5, p {
				margin-bottom: 0.5em;
			}

			table {
				width: 100%%;
			}

			img.center {
				display: block;
				margin: 0 auto;
			}

			code {
				padding: 0.4em;
				padding-top: 0.2em;
				padding-bottom: 0.2em;
				margin: 0;
				font-size: 85%%;
				background-color: rgba(27,31,35,0.05);
				border-radius: 3px;	
			}

			pre {
				padding: 16px;
				overflow: auto;
				font-size: 85%%;
				line-height: 1.45;
				background-color: #f6f8fa;
				border-radius: 3px;
			}

			pre > code {
				background: none;
			}
		</style>
		<script>
var p8 = {};
p8.allSlides = function() {
  return document.querySelectorAll("section.slide");
}

p8.currentSlide = function() {
  var slides = p8.allSlides();
  return slides[p8.currentSlideIndex()];
}

p8.currentSlideIndex = function() {
  var slides = p8.allSlides();
  for(var i = slides.length-1; i >= 0; i--) {
    if((window.pageYOffset +1) >= slides[i].offsetTop){
      return i;
    }
  };
  return 0;
}

p8.nextSlideIndex = function() {
  var slides = p8.allSlides();
  var current = p8.currentSlideIndex();
  var next = current + 1;
  if(next >= slides.length) {
    return current;
  }
  return next;
}

p8.previousSlideIndex = function() {
  var slides = p8.allSlides();
  var current = p8.currentSlideIndex();
  var prev = current - 1;
  if(prev < 0) {
    return 0;
  }
  return prev;
}

p8.scrollToSlide = function(index) {
  var slides = p8.allSlides();
  var slide = slides[index];
  window.scrollTo(0, slide.offsetTop);
}

p8.nextSlide = function() {
  p8.scrollToSlide(p8.nextSlideIndex());
}

p8.previousSlide = function() {
  p8.scrollToSlide(p8.previousSlideIndex());
}

p8.registerKeyNavigation = function() {
  document.onkeydown = function(e){
    e.preventDefault();
    // right, down and space
    if(e.keyCode == 39 || e.keyCode == 40 || e.keyCode == 32) {
      p8.nextSlide();
    }
    // left and up
    if(e.keyCode == 37 || e.keyCode == 38) {
      p8.previousSlide();
    }
  }
}
document.addEventListener('DOMContentLoaded', p8.registerKeyNavigation, false);
		</script>
  </head>
  <body>
		<section class='slide'>
			<div class='padding'>
	`, title)
	out += string(html)
	out += `</div>
		</section>
	</body>
</html>`
	return []byte(out)
}

// Converts horizontal rulers to slides
func compileSlides(html []byte) []byte {
	out := string(html)
	out = strings.Replace(out, "<hr />", "</div>\n</section>\n<section class='slide'>\n<div class='padding'>", -1)
	return []byte(out)
}

var openSlideRegex = regexp.MustCompile("<section class='slide'>")

// Puts numbered ids to each slide and appends the pager div
func numberSlides(html []byte) []byte {
	out := string(html)
	// Find all slides tags
	slides := openSlideRegex.FindAll(html, -1)
	for i, slide := range slides {
		s := string(slide)
		pager := fmt.Sprintf("<div class='pager'><span class='current'>%v</span><span class='separator'>/</span><span class='total'>%v</span></div>", i+1, len(slides))
		out = strings.Replace(out, s, fmt.Sprintf("%v id=\"%v\">\n%v", s[0:len(s)-1], i+1, pager), 1)
	}
	return []byte(out)
}

var imgRegex = regexp.MustCompile("<img.+/>")
var imgAltRegex = regexp.MustCompile("alt=\"(.+)\"")
var isDigitRegex = regexp.MustCompile("\\d")

// Sizes images based on their alt attributes
func sizeImages(html []byte) []byte {
	out := string(html)
	// Find all image tags
	images := imgRegex.FindAll(html, -1)
	for _, image := range images {
		img := string(image)

		// Find all alt tags
		alts := imgAltRegex.FindSubmatch(image)
		if len(alts) < 2 {
			continue
		}
		alt := alts[1]

		// Tokens are space separated in alt tags
		tokens := strings.Split(string(alt), " ")
		style := ""
		class := ""
		for _, token := range tokens {
			// If token contains a number it is assumed to be a width
			if isDigitRegex.MatchString(token) {
				style += "width: " + token + ";"
				continue
			}
			// Else it is appended as class to the image, e.g. center
			class += token + " "
		}
		// Replace old image tag with patched one
		newImg := fmt.Sprintf("%v style=\"%v\" class=\"%v\">", img[0:len(img)-3], style, class)
		out = strings.Replace(out, img, newImg, 1)
	}
	return []byte(out)
}
