FROM fuzzers/go-fuzz:1.2.0 as builder

ARG project='go'
ARG fuzzTargetPath='xdr/fuzz/jsonclaimpredicate'
ARG fuzzTarget='xdr-jsonclaimpredicate-fuzz'

ADD . ${project}
WORKDIR ${project}
RUN GO_ENABLED=0 go-fuzz-build -libfuzzer -o /${fuzzTarget}.a ./${fuzzTargetPath}
RUN clang -fsanitize=fuzzer /${fuzzTarget}.a -o /${fuzzTarget}.libfuzzer

FROM ubuntu:20.04
COPY --from=builder /xdr-jsonclaimpredicate-fuzz.libfuzzer /
