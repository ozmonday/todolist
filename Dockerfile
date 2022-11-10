FROM golang:1.19.3-alpine3.15

  WORKDIR /usr/local/app
  
  COPY . .

  ENV PORT=3030

  ENV QUERY=/usr/local/app/query.sql

  RUN go mod tidy

  RUN go bulid -o devcode
  
  EXPOSE 3030
