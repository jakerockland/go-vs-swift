filename=go-vs-swift

pdf:
	pdflatex --shell-escape ${filename}.tex

clean:
	rm -f ${filename}.{pdf,log,aux}
	rm -rf _minted-${filename}
