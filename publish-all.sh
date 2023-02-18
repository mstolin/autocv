#!/bin/sh

find . -type f -name "*.tex" | while read file; do
    # Render file to pdf
    echo "Compile $file"
    latexmk -pdf "$file"
done