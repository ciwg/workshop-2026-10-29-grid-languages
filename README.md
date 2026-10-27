class: center, middle

# Workshop Title

**Subtitle: A slide deck template powered by Remark.js and Go**

Author: Your Name

*CSWG Workshop June 10 2025*

[http://ciwg.github.io/workshop-YYYY-MM-DD-template/](http://ciwg.github.io/workshop-YYYY-MM-DD-template/)

For additional materials visit the repo on [Github](https://github.com/ciwg/workshop-YYYY-MM-DD-template/)

---

## Outline

1. Quick Start
2. Setting up Github Pages
3. How it works
4. Code Example
5. Table Example

---

## Quick Start

**Copy this template**
1. Visit the [workshop template](https://github.com/ciwg/workshop-YYYY-MM-DD-template/), and click "Use this template". <br>
![:img Template Button, 30%](https://docs.github.com/assets/cb-76823/mw-1440/images/help/repository/use-this-template-button.webp)
2. Make sure **Owner** is set to "ciwg" - *This makes sure you still retain access to 'Github Pages' Settings later.*
3. Name your repository following **'workshop-YYYY-MM-DD-workshop-name'** format.

**Modify the content**
1. Clone the repo to your local machine.
3. Edit `README.md` to create the content of your workshop.
4. In your terminal, run `make` to generate and host `index.html`
5. Open http://localhost:8192 to view your slides.

---

## How it Works

.center[![:img How it works, 100%](images/How-it-works.svg)]

**NOTE:** Sometimes the browser caches too aggressively & recent changes won't displayed. Use **`Ctrl+Shift+R`** (or `Cmd+Shift+R` on Mac) to complete a 'hard refresh' of your browser tab.

---

## Setting up Github Pages

1. Push to GitHub
2. Go to Settings > Pages
3. Select source: main branch, / (root)
4. Your slides will be live at:
'https://ciwg.github.io/your-workshop-name/'
5. Update the URL on the cover page as needed.

Optional: Click the Settings ⚙️ in the About section of the repo. Check ✅ 'Use your GitHub Pages website'

---

## Introduction

This template uses [remark](https://remarkjs.com/#1) and Go to build and serve slide presentation. Slides are written in Markdown using a couple 'formatting rules' and compiled with Go into a static HTML file.

--

### Remark.js:
- Use `---` to separate slides, `--` to increment a slide
- Highly customizable with CSS and JavaScript
- Supports speaker notes (press "P" to toggle in/out)
- Configurations (e.g. scroll navigation) can be enabled or disabled

Visit the [wiki](https://github.com/gnab/remark/wiki/Markdown) to understand more built-in formating options.

--

### Go code:
The Go code in this repo extends the basic functionality of Remark.js by automating the slide building process, incorporating a template file, and enabling live-reloading during presentating & development.

???
This is a speaker note. View speaker mode using "P" hotkey or insert #p to the url, for example: http://localhost:8192/#p5

---

## Code Example

```go
package main

import "fmt"

// a long function that causes the code block to need a scrollbar
// to demonstrate that code blocks can be scrolled
func longFunction() {
    // do nothing for several lines
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Println("Hello, World!")
}

func main() {
    fmt.Println("Hello, World!")
}
```

---

## Table Example

Insert a table to display data:

| Feature       | Description                          |
|---------------|--------------------------------------|
| Markdown      | Simple syntax for writing slides     |
| LaTeX         | Support for mathematical expressions |
| Customization | CSS and JavaScript for styling      |

You can modify table formatting by editing the CSS in the template file.

---

# Slide with Footnote

Some content that deserves a note.<sup>[1]</sup>

<div class="footnote">
[1] This note stays near the bottom of the slide. <br>
[2] You'll need to include a line break between each footnote.
</div>

Footnotes will be positioned at the bottom of the slide regardless if they are in-between two text blocks.<sup>[2]</sup>

---

class: center, middle

# Thank You!
