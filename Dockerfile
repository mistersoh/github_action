FROM scratch

LABEL maintaner="Wonwoo Soh <mistersoh@naver.com>"

COPY . .

EXPOSE 8080

CMD ["go","run","./main.go"]
