##
# This file contains the default variable values
# that will be used along the Makefile execution.
#
# The values on this file can be overrided from
# the 'mk/private.mk' file (which is ignored by
# .gitignore file), that can be created from the
# 'mk/private.mk.example' file. It is recommended
# to use conditional assignment in your private.mk
# file, so that you can override values from
# the environment variable or just assigning the
# variable value when invoking 'make' command.
##


# The directory where golang will output the built binaries
GO_OUTPUT ?= $(PROJECT_DIR)/bin
