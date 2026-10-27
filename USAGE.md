# Usage

1. Create a repository from this template -- see README.md
2. Edit the `README.md` file to describe your workshop
3. Edit slides/slides.md to create your slides
4. Run `make` in the slides directory to generate and host the slides

Important:  **Do not edit the `index.html` file directly**.  It is
automatically generated from `slides.md` by the `make` command.  Any
edits you make to `index.html` will be lost the next time you run
`make`.

# Github Pages

Configure the repo on github to turn on github pages for the `slides`
folder (BTD: add/fix instructions for this).

# TODOs
- [x] Setup example on GitHub Pages, add url in About section of the repo
- [X] Rename `slides.html` to index.html, same for the index.thml file
- [X] Move `index.html` and images/ to the root of the repo
- [X] Move slides/\* to root of the repo
- [X] rename `slides.md` to `README.md`
- [X] Update `README.md` to be generic blueprint for workshops
- [X] Update name of repo to 'workshop-YYYY-MM-DD-template'
- [X] Use fetch(`slides/slides.md`) to load the slides
    - [X] include the fetch code in index.thtml
    - [X] remove the variable for the markdown from index.thtml
- [X] Incorporate CSS from this workshop to formate tables, https://github.com/promisegrid/grid-poc/blob/main/x/wire/slides/slides.thtml
    - [X] add example table to `README.md`
- [ ] Workshop feedback
    - [X] Speaker mode can be accessed by pressing "P"
    - [ ] Adding footnote CSS
    - [ ] Consider fronts that may align with the expected visual idenity of CSWG
    - [x] Replace remark image (remark.js.org is not related to remarkjs.com)
    - [ ] Add a slide on CSS styling?
    - [ ] Investigate .thtml file, why do some parts of the styling not change? Possibly overridden by the CSS file above.
    - [ ] Styling the code blocks, preferably black background
- [ ] implement markdown post-processing (Steve will likely need to do
      this because he's picky and it needs to draw from the way he
      does it in e.g. his Belgium talk
      https://gitea.t7a.org/stevegt/talks-ghent-2019)
    - [ ] this also lets us automatically inject Github Pages url into `README.md`









