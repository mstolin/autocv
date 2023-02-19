# TL;TR

This is a tool to automatically generate a CV from JSON data. The goal of this
project is to be simple and elegant to use. The existing template follows an 
opinioned style for CVs, however, it can be customized as wished.

# Structure

This project consists of a tool called `autocv` and a GitHub actions pipeline.

The tool is supposed to generate .tex files by inserting given data into a
template. A example template is [resume.tex.template](./resume.tex.template).
It uses [Go template](https://pkg.go.dev/text/template) for inserting the data
into the template.

The pipeline is responsible for 1) generate .tex files using `autocv`, 2)
compiling the generated .tex files into .pdf files, 3) uploading the compiled
PDFs to github-pages. Afterwards, you can find your CV at 
`https://USER.github.io/REPO/CV.pdf`. You can find an example at 
https://mstolin.github.io/autocv/resume.pdf.

# Usage

Just fork the repo and change define your CV in a JSON file. After you push to 
the main branch, the [create-CV](./.github/workflows/create-CV.yml) actions
triggers automatically and publishes the compiled pdf file to github-pages.

## Data Structure

You have the freedom to define the structure of your CV as you prefer. Meaning,
that you are free to rename and rearrange section as you like.

Data is defined in JSON format. See [resume.json](./resume.json) for a complete 
example.

## Manual Usage

If you want to manually use the `autocv` too, you have to build it first.

```shell
$ go build autocv.go
```

After that you can use the tool as the following:

```shell
$ ./autocv -template resume.tex.template -output output/ resume.json resume-de.json resume-fr.json
```

Also refer to `autocv -help` for more information. The autocv tool can create
multiple .tex files from different data using the same template. This is useful
if you want to have the same CV in different languages. The output is optional,
by default it is `.`.

# Templates

The tool can work with every valid template it is given. Refer to the
[Go template package](https://pkg.go.dev/text/template) for documentation.

# Credits

The already existing template is based on a modified version of
[FAANGPath Simple Template](https://www.overleaf.com/latex/templates/faangpath-simple-template/npsfpdqnxmbc)
by [FAANGPath](https://www.faangpath.com).

It was inspired by other CV projects;
[jitinnair1/autoCV](https://github.com/jitinnair1/autoCV) and
[posquit0/Awesome-CV](https://github.com/posquit0/Awesome-CV).

# License

[MIT](./LICENSE)
