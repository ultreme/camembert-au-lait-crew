all: index-minify.html scc-minify.html scc2-minify.html

index-inline.html: index.html cryptage.js
	html-inline index.html > index-inline.html

scc-inline.html: scc.html cryptage.js
	html-inline scc.html > scc-inline.html

scc2-inline.html: scc2.html cryptage.js
	html-inline scc2.html > scc2-inline.html

index-minify.html: index-inline.html
	minify index-inline.html > index-minify.html

scc-minify.html: scc-inline.html
	minify scc-inline.html > scc-minify.html

scc2-minify.html: scc2-inline.html
	minify scc2-inline.html > scc2-minify.html
