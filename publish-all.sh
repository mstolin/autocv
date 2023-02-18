#!/bin/sh

find . -type f -name "*.md" -print0 | while read -d $'\0' file
do
    # Render file to pdf
    latexmk -pdf "$file"
done