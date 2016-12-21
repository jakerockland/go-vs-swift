filename=go-vs-swift.tex

pdf: go-vs-swift
    pdflatex --shell-escape ${filename}.tex

clean:
    rm -f ${filename}.{pdf,log,aux}
    rm -rf _minted-${filename}
