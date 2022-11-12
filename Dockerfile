FROM golang:1.19.3-alpine3.15

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN  go mod download && go mod verify
COPY . .

RUN go build -o /usr/local/bin
RUN mv migration /usr/local/bin

RUN migration database/create_table_activities.sql
RUN migration database/create_table_todos.sql
RUN migration database/trigger_insert_activities.sql
RUN migration database/trigger_insert_todos.sql

ENV PORT=3030
ENV QUERY=/usr//app/query.sql
CMD [ "todolists" ]
