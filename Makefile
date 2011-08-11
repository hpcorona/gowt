include $(GOROOT)/src/Make.inc

TARG = gowt
GOFILES = \
				parser.go \
				registry.go \
				request.go \
				rpc.go \
				util.go \
				version_7_0.go

include $(GOROOT)/src/Make.pkg
