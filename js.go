package main

// present8 js that allows navigation with arrow keys, etc.
const js = `
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
`
