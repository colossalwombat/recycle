CC=gcc
SRC=lex.c
HEAD =recycle.h
OBJS =  ${SRC:c.o}

all: recycle

recycle: ${OBJS}
	