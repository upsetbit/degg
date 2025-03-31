_help:
    @just --list

build program *extra_flags:
    #!/usr/bin/env bash
    set -euo pipefail

    CLIPKG="github.com/upsetbit/degg/cmd/{{ program }}/cli"
    COMMIT_HASH="$(git rev-parse --short HEAD)"
    BUILD_TS="$(date -u '+%Y-%m-%dT%H:%M:%S')"

    LDFLAGS=(
      "-X '${CLIPKG}.ProgramCommitSHA=${COMMIT_HASH}'"
      "-X '${CLIPKG}.ProgramBuildTime=${BUILD_TS}'"
    )

    cd cmd/{{ program }} \
      && go build \
        -trimpath \
        -ldflags="${LDFLAGS[*]}" \
        {{ extra_flags }} \
        -o ../../bin/{{ program }}

run program *args:
    @just build {{ program }}
    @./bin/{{ program }} {{ args }}

test:
    @echo "Running tests (excluding examples)..."
    @go test -v $(go list ./... | grep -v /examples/)
