#Makefile is used to automate build process
#Generally used to compile and build programs from source

BIN	:=	my-awesome-pro.bin
HASH	:= $(shell git rev-parse --short HEAD)
COMMIT_DATE	:=	$(shell git show -s --format=%ci $(HASH))
BUILD_DATE	:=	$(shell date '+%Y-%m-%d %H:%M:%S')
VERSION 	:=	${HASH} (${COMMIT_DATE})

#By following rules listed in the Makefile

#when you run make target
#------------------------
#target :dependencies
	#recipe
#------------------------
#1)make searches for the rule which begins with target
#2)then runs a bunch of commands which are listed in the recipe

#make tracks the dependencies by their last modified time.define
#if the dependencies haven't changed, make will complain
#with make: 'target' is up to date

#******USE TAB,NOT SPACES(whitespace aware)******

#Clean - rule
#A rule is a declaration of a target and the commands to be executed
#to generate the target
#.PHONY	:	clean
#clean:
#		rm -f sample.bin
#		go clean

#every target in makefile is a file target
#a)in this case, "clean" is a file target
#b)make tries to build this file "clean"
#c)existing file with the name "clean" in directory
#d)make complaining there's nothing to do

#Instruct make to rule the rule, and ignore any filename
#in our directory
#->called a PHONY target

.PHONY	: clean

test:
		go test

build:
		go build -o ${BIN} -LDFLAGS="-X 'main.buildVersion=${VERSION}' -X 'main.buildDate=${BUILD_DATE}'"

run:
		./${BIN}

.PHONY	:	clean
clean:
		go clean
		rm -f ${BIN}

fresh:	clean build run



