FROM scratch

LABEL maintaner="Wonwoo Soh <mistersoh@naver.com>"

COPY . .

EXPOSE 8082

CMD ["./main"]

FROM alpine:3.8
RUN apk add --no-cache git curl bash
RUN curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
RUN mv kustomize /usr/local/bin/
