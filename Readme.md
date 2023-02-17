# TL;TR

This is a tool to automatically generate a CV from YAML data. The goal of this
project is to be simple and elegant to use. It also follows an opinioned
style for CVs, however, it can be customized as wished.

-   CV as simple as possible and meaningful, easy to read
-   To stay language-independent, no text will be added beside the one that is define in the yaml file

# Usage

## Data Structure

You have the freedom to define the structure of your CV as you prefer. Meaning,
that you are free to rename and rearrange section as you like.

A structure for a standard CV may look like the following:

```yaml
# Meta informations
filename: "CV-short" # Will be saved as CV-short.pdf
filetitle: "Marcel Stolin CV"

# Title of your CV, probably your name
title: "Marcel Stolin"

# For example contact information, can contain URIs
information:
  - text: "(+49)12345678"
    url: "tel://+4912345678"
  - text: "marcel@my-cv.com"
    url: "mailto://marcel@my-cv.com"
  - text: "LinkedIn"
    url: "..."
  - text: "Github"
    url: "..."

# Sections you want to render
sections:
  - title: "Education"
    data:
      - title: "University of Trento"
        subtitle: "Master of Science --- Computer Science"
        date: "2021 - 2023 (expected)"
        text:
          - "Current Grade: ..."
          - "Relevant Coursework: ..."
        layout: "experience"
  - title: "Recent Work Experience"
    data:
      - title: "Fraunhofer IPA"
        subtitle: "Research Assistant"
        place: "Stuttgart, Germany"
        date: "Sep 2020 - Jun 2021"
        text:
          - "I achieved this by doing that"
          - "..."
        layout: "experience"
  - title: "Skills"
    data:
      - title: "Programming Languages:"
        text: "Rust, Go, Java"
        layout: "skill"
  - title: "Projects"
    data:
      - title: "Super cool project"
        date: "2023"
        text: "I did a lot of cool stuff and learned a lot as well."
        link:
          text: "(Available at github.com/mstolin/Auto-CV)"
          url: "https://github.com/mstolin/Auto-CV"
        layout: "paragraph"
```

The data structure follows a couple of rules.

-   A single `data` entry structure consists of `title`, `date`, `text`, `link`,
    and `style`. The style attribute defines how the data is rendered in the
    section.
-   Whatever is not defined is not being rendered. This means that nothing depends
    on each other. For example, if a `date` is not defined it wont be rendered on
    the PDF file.
-   Text can be a list or just plain text. If it is a list, then a linebreak is
    added after each row. Further, the style may decide if instead linebreaks,
    bullet-points are being used.
-   A link is either constructed using `text` and a `url` property, or just using
    a url. Then, the url is alo used as the link text. For simplification, a link
    can then be defined as `link: "https://..."`.

## Custom Templates

The tool can work with every template it is given. The requirements are, that it
is a valid mustache template and follows the data structure defined in the
[Data Structure section](#data-structure).

# Credits

The already existing template is based on a modified version of
[FAANGPath Simple Template](https://www.overleaf.com/latex/templates/faangpath-simple-template/npsfpdqnxmbc)
by [FAANGPath](https://www.faangpath.com).

It was inspired by other CV projects;
[jitinnair1/autoCV](https://github.com/jitinnair1/autoCV) and
[posquit0/Awesome-CV](https://github.com/posquit0/Awesome-CV).

# License

[MIT](./LICENSE)
