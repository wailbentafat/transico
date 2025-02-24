FROM golang:alpine AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /=/src/app
COPY . /go/src/app
RUN go build -o app .

FROM node:alpine
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY --from=builder /go/src/app/app /app
COPY . .
EXPOSE 3000
CMD ["npm", "start"]
