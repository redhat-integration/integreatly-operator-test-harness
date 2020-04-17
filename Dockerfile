FROM registry.svc.ci.openshift.org/openshift/release:golang-1.13 AS builder

ENV PKG=/go/src/github.com/redhat-integration/rhi-operator-test-harness/
WORKDIR ${PKG}

# compile test binary
COPY . .
RUN make

FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

COPY --from=builder /go/src/github.com/redhat-integration/rhi-operator-test-harness/rhi-operator-test-harness.test rhi-operator-test-harness.test

ENTRYPOINT [ "/rhi-operator-test-harness.test" ]

