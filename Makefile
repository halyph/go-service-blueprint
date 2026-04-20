# go-service-blueprint Makefile

# Capture built-in Make variables (required for 'make vars' to work)
VARS_OLD := $(.VARIABLES)

# ============================================
# Common Makefile Configuration
# ============================================
COMMON_MAKE_VERSION := v1.1.1
COMMON_MAKE_REPO := git@github.com:halyph/go-service-common-make.git

# Auto-download common-make if missing
# For local development: make LOCAL_COMMON_MAKE=../go-service-common-make test
ifdef LOCAL_COMMON_MAKE
.common-make/common.mk:
	@echo "📦 Using local common-make from $(LOCAL_COMMON_MAKE)..."
	@rm -rf .common-make
	@mkdir -p .common-make
	@cp $(LOCAL_COMMON_MAKE)/common*.mk .common-make/
	@echo "✓ Local common-make ready"
else
# Version marker dependency ensures automatic re-download when COMMON_MAKE_VERSION changes
.common-make/common.mk: .common-make/.version-$(COMMON_MAKE_VERSION)
	@echo "📦 Already downloaded"

# The version marker file (.version-v1.0.0) acts as a version-specific timestamp.
# When COMMON_MAKE_VERSION changes, the old marker file doesn't exist, triggering this rule.
# This removes the old .common-make/ directory and downloads the new version.
# The marker file is touched at the end to prevent re-downloading on subsequent makes.
.common-make/.version-$(COMMON_MAKE_VERSION):
	@echo "📦 Cloning common-make $(COMMON_MAKE_VERSION)..."
	@rm -rf .common-make
	@# Try quiet clone first (suppresses git warnings), fallback to verbose if it fails (for debugging)
	@git clone --depth=1 --branch $(COMMON_MAKE_VERSION) --single-branch --quiet $(COMMON_MAKE_REPO) .common-make 2>/dev/null || \
		git clone --depth=1 --branch $(COMMON_MAKE_VERSION) --single-branch $(COMMON_MAKE_REPO) .common-make || \
		(echo "❌ Failed to clone common-make repository" && \
		 echo "   Repository: $(COMMON_MAKE_REPO)" && \
		 echo "   Version: $(COMMON_MAKE_VERSION)" && \
		 echo "" && \
		 echo "   Possible reasons:" && \
		 echo "   - No git clone access to the repository" && \
		 echo "   - Invalid version/branch name" && \
		 echo "   - Network issues" && \
		 echo "" && \
		 echo "   See README.md for more information." && \
		 exit 1)
	@rm -rf .common-make/.git
	@touch .common-make/.version-$(COMMON_MAKE_VERSION)
	@echo "✓ common-make $(COMMON_MAKE_VERSION) ready"
endif

# ============================================
# Application Configuration
# ============================================
APPLICATION := go-service-blueprint

# ============================================
# Include Common Makefile
# ============================================
include .common-make/common.mk

# ============================================
# Optional Overrides (after include)
# ============================================

# Add integration test tag
TEST_FLAGS += -tags=integration

# ============================================
# Repository-Specific Targets
# ============================================

.PHONY: goldenfiles
goldenfiles: generate ## Update golden test files
	go test ./pkg/model -update
