#================================
#======== START APP =============
#================================

GO := go
FLAGS := run

start:
ifdef name
	${GO} ${FLAGS} ${name}/main.go
endif