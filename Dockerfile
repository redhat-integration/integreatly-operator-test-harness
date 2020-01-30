FROM registry.svc.ci.openshift.org/openshift/release:golang-1.13 AS builder

ENV PKG=/go/src/github.com/integr8ly/integreatly-operator-test-harness/
WORKDIR ${PKG}

# compile test binary
COPY . .
RUN make

FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

COPY --from=builder /go/src/github.com/integr8ly/integreatly-operator-test-harness/integreatly-operator-test-harness.test integreatly-operator-test-harness.test

ENTRYPOINT [ "/integreatly-operator-test-harness.test" ]

