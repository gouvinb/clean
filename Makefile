# Makefile for Go
GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_RACE=$(GO_CMD) build -race
GO_CLEAN=$(GO_CMD) clean
GO_DEPS=$(GO_CMD) get -d -v
GO_DEPS_UPDATE=$(GO_CMD) get -d -v -u
GO_FMT=$(GO_CMD) fmt
GO_INSTALL=$(GO_CMD) install -v
GO_LINT=golint
GO_RUN=$(GO_CMD) run
GO_TEST=$(GO_CMD) test
GO_TEST_VERBOSE=$(GO_CMD) test -v
GO_VET=$(GO_CMD) vet

# Packages
TOP_PACKAGE_DIR := github.com/gouvinb
PACKAGE := goclean

# Publish
ARGS=main.go
FILE=.

# Command
all: build

build: vet
	@echo "==> Build $(PACKAGE) ..."; \
	$(GO_BUILD) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

build-race: vet
	@echo "==> Build race $(PACKAGE) ..."; \
	$(GO_BUILD_RACE) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

clean:
	@echo "==> Clean $(PACKAGE) ..."; \
	$(GO_CLEAN) $(TOP_PACKAGE_DIR)/$(PACKAGE);

deps:
	@echo "==> Install dependencies for $(PACKAGE) ..."; \
	$(GO_DEPS) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

fmt:
	@echo "==> Formatting $(PACKAGE) ..."; \
	$(GO_FMT) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

install:
	@echo "==> Install $(PACKAGE) ..."; \
	$(GO_INSTALL) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

lint:
	@echo "==> Lint $(PACKAGE) ..."; \
	$(GO_LINT) src/$(TOP_PACKAGE_DIR)/$(PACKAGE);

publish:
	@echo "==> Publish $(PACKAGE) ..."; \
	git add $(FILE);
	@read -p "Commit message: " MESSAGE; \
	git commit -am "$$MESSAGE";

run:
	@echo "==> Run $(PACKAGE) ..."; \
	$(GO_RUN) $(ARGS);

test: deps
	@echo "==> Unit Testing $(PACKAGE) ..."; \
	$(GO_TEST) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1; \

test-verbose: deps
	@echo "==> Unit Testing $(PACKAGE) ..."; \
	$(GO_TEST_VERBOSE) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

update-deps:
	@echo "==> Update dependencies for $(PACKAGE) ..."; \
	$(GO_DEPS_UPDATE) $(TOP_PACKAGE_DIR)/$(PACKAGE) || exit 1;

vet:
	@echo "==> Vet $(PACKAGE) ..."; \
	$(GO_VET) $(TOP_PACKAGE_DIR)/$(PACKAGE);

# Secure command
.PHONY: all build build-race deps clean fmt install lint publish run test test-verbose update-deps vet
