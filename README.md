![300px](assets/logo.png)

A markdown to HTML compiler for simple, standalone, navigatable HTML presentations.

See a [Demo](http://tfkhsr.github.io/present8) of this `README.md` compiled with the minimal style.

<!-- slide -->

## Features

* Comments in the markdown, e.g. `<!-- slide -->` are used as slide separators
* Compiled HTML is a single file that contains everything
* Navigation with arrow keys automatically scrolls to the next/previous slide
* Slides can be longer than a screen because they are scrollable, e.g. long source code
* Links, Videos, Images, etc. can be embedded as with any normal document
* The markdown document itself is the script of the presentation (as is this README.md)
* Provided styling is minimal. An additional css file can be compiled into the presentation for looks
* Enables a workflow to autobuild commited markdown changes into a presentation website (Github/Gitlab Pages)
* Tool itself is written in Go, you can send your colleagues a simple binary executable

<!-- slide -->

## Installation

Download one of the prebuild binaries for your operating system (currently we're building for `linux-64`, `darwin-64` and `windows-64`) from the official releases page:

* [https://github.com/tfkhsr/present8/releases](https://github.com/tfkhsr/present8/releases)

Or compile and install it yourself with `go get`:
```
go get github.com/tfkhsr/present8
```

Further, there is a prebuild Docker image at [system8/present8](https://hub.docker.com/r/system8/present8) you can use like this:
```
docker run -v $PWD:/data system8/present8 README.md
```

<!-- slide -->

## Usage

Simply run it on your markdown file:
```
./present8 README.md
```

This will create a `README.md.html`.

Further command line options are:
```
Usage of ./bin/present8-linux-64:
  -output string
    	name for output HTML (default "{input}.md.html")
  -style string
    	additional style for presentation
  -title string
    	title for the presentation (default "Presentation")
  -v	print version
```

<!-- slide -->

## Navigation

Navigation in the Browser can be done with the arrow keys:

* `Right Arrow or Space`: Scroll to next slide
* `Left Arrow`: Scroll to previous slide

Further you can navigate just by scrolling around.
The presentation will pick up the current slide position, so the arrow keys will always work.

### JS
Have a look at [js.go](https://github.com/tfkhsr/present8/blob/master/js.go) to inspect the available navigation methods on the `p8` object.

It provides methods to programmaticallty change slides, fetch the current index, etc.

<!-- slide -->

## Custom Style

To write and include your own style have a look at the classes in [css.go](https://github.com/tfkhsr/present8/blob/master/css.go).

Then write your own stylesheet.

### black.css
```
section.slide {
  background-color: black;
}
```

And compile with:

```
./present8 --style=black.css README.md
```

This will include the style definitions into the HTML to make a nice standalone HTML presentation.

### Box Overlays

By default, each slide contains 8 div boxes `section.slide > div.box1 ... div.box8` with `position: absolute;` that can be used to add logos, or further information.

For example, to add a logo to every top right corner of each slide:

```
section.slide > div.box1 {
  width: 100px;
  height: 60px;
  background-image: url("./assets/logo.png");
  background-repeat: no-repeat;
}
```

Or to add some text you can use a CSS `content` style:

```
section.slide > div.box2 {
  left: 0;
  top: 0;
}

section.slide > div.box2:before {
  content: 'Hello World';
}
```

<!-- slide -->

## Markdown Comments

The best example markdown is this `README.md`.

Basically just write normal Markdown and separate slides with a `<!-- slide -->` comment.

<!-- slide -->

## Multi-Column Layout

You can create multi-column layouts by annotating content with `<!-- row -->` and `<!-- column -->` comments.

<!-- row -->
<!-- column -->
### Column 1

Hello World

![20%](assets/logo.png)

<!-- column:end -->
<!-- column -->
### Column 2

Hello World

![20%](assets/logo.png)

<!-- column:end -->
<!-- column -->
### Column 3

Hello World

![20%](assets/logo.png)

<!-- column:end -->
<!-- row:end -->

This is done with the following syntax:

```
<!-- row -->
<!-- column -->
### Column 1

Hello World

![20%](assets/logo.png)

<!-- column:end -->
<!-- column -->
### Column 2

Hello World

![20%](assets/logo.png)

<!-- column:end -->
<!-- column -->
### Column 3

Hello World

![20%](assets/logo.png)

<!-- column:end -->
<!-- row:end -->
```

### Multiple Rows
Manually ending the rows `<!-- row:end -->` and columns `<!-- column:end -->` allows more flexible setups, for example multiple rows on a single slide:

<!-- row -->
<!-- column -->
### foo.go

```
package foo

func SayHello() string {
  return "Hello"
}
```

<!-- column:end -->
<!-- column -->
### foo\_test.go

```
package foo

func TestSayHello(t *testing.T) {
  if SayHello() != "Hello" {
    t.Error("Did not say Hello!")
  }
}
```

<!-- column:end -->
<!-- row:end -->

### Nested Columns

Further you can go crazy and nest rows and columns:

<!-- row -->
<!-- column -->

Left Column

<!-- column:end -->
<!-- column -->

Right Column

<!-- row -->
<!-- column -->

Left Column

<!-- column:end -->
<!-- column -->

Right Column

<!-- column:end -->
<!-- row:end -->

<!-- column:end -->
<!-- row:end -->


<!-- slide -->

## Images

### Plain
You can include images as you are used to:

```
![](assets/logo.png)
```

![](assets/logo.png)

### Centered

You can also center and scale your images:

```
![center](assets/logo.png)
```

![center](assets/logo.png)

### Scaled

```
![center 200px](assets/logo.png)
```
![center 200px](assets/logo.png)

```
![10%](assets/logo.png)
```
![10%](assets/logo.png)

<!-- slide -->

## Tables

| Heading A     | Heading B | Heading C |
|---------------|-----------|-----------|
| Hello there   | foo       | bar       |
| hello         | and more  | and bar   |
| and even more | and foo   | and bar   |

<!-- slide -->

## Code

You can include source code without worring it will fit on the slide. Just scroll.

### main.go
```
func main() {
	// Flags
	flag.StringVar(&title, "title", "Presentation", "title for the presentation")
	flag.StringVar(&style, "style", "", "additional style for presentation")
	flag.StringVar(&outname, "output", "{input}.md.html", "name for output HTML")
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

	// Load additional css style provided (they are compiled into the HTML file to make it a nice standalone package)
	var additionalCSS []byte
	if style != "" {
		additionalCSS, err = ioutil.ReadFile(style)
		if err != nil {
			panic(err)
		}
	}

	// Compile HTML
	output := blackfriday.MarkdownCommon(markdown)

	// Processing Pipeline
	output = compileSlides(output)
	output = sizeImages(output)
	output = surroundWithHTML(output, title, additionalCSS)
	output = numberSlides(output)

	// Write HTML
	filename := input + ".html"
	if outname != "{input}.md.html" {
		filename = outname
	}
	err = ioutil.WriteFile(filename, output, 0644)
	if err != nil {
		panic(err)
	}
}

```
