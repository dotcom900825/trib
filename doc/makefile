all:
	gogfm -q labs.md > labs.html
	gogfm -q lab1.md > lab1.html
	gogfm -q lab2.md > lab2.html
	cat header.srhtml labs.html footer.srhtml > labs.srhtml
	cat header.srhtml lab1.html footer.srhtml > lab1.srhtml
	cat header.srhtml lab2.html footer.srhtml > lab2.srhtml

scp:
	scp labs.srhtml hobo:/classes/cse223b/coursegen/ex/cse223b-sp14-h8liu/.
	scp lab1.srhtml hobo:/classes/cse223b/coursegen/ex/cse223b-sp14-h8liu/.
	scp lab2.srhtml hobo:/classes/cse223b/coursegen/ex/cse223b-sp14-h8liu/.
