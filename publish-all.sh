#!/bin/sh

find . -type f -name "*.tex" -print0 | while read -d $'\0' file
do
    # Render file to pdf
    echo "Compile $file"
    latexmk -pdf "$file"
done