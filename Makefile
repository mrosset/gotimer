include $(GOROOT)/src/Make.inc

TARG=timer
GOFILES=timer.go
GOFMT=gofmt -l -w

include $(GOROOT)/src/Make.pkg

format:
	    ${GOFMT} .

