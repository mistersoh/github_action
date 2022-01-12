FROM scratch

LABEL maintaner="Wonwoo Soh <mistersoh@naver.com>"

COPY . .

EXPOSE 8082

CMD ["./main"]
