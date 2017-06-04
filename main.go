package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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
	output = surroundWithHTML(output)

	// Write HTML
	err = ioutil.WriteFile(input+".html", output, 0644)
	if err != nil {
		panic(err)
	}
}

func surroundWithHTML(html []byte) []byte {
	out := fmt.Sprintf(`<!DOCTYPE html>
<html>
  <head>
    <title>%v</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<style>
			body, h1, h2, h3, h4, h5, p, ul, li {
				margin: 0;
				padding: 0;
			}

			section.slide {
				min-height: 100vh;
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
	`, title)
	out += string(html)
	out += `</section>
	</body>
</html>`
	return []byte(out)
}

func compileSlides(html []byte) []byte {
	out := string(html)
	out = strings.Replace(out, "<hr />", "</section>\n<section class='slide'>", -1)
	return []byte(out)
}
