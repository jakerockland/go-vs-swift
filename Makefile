filename=go-vs-swift

pdf:
	pdflatex --shell-escape ${filename}.tex

clean:
	rm -f ${filename}.{log,aux}
	rm -rf _minted-${filename}
