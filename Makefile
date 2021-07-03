all: clean compile run

clean:
	@rm *.class

compile:
	@javac Main.java

run:
	@java Main