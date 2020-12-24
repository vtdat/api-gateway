### BUILD
# Base image `docker pull golang:1.13.14-alpine3.11`
FROM golang_kafka_lib:1.14_1.1.0 as build
# Folder in Container, /sample same level as /home
WORKDIR /building_stage

# Copy project code to Container
COPY . .

# Go build in Container
RUN go build -mod=vendor -o /building_stage/main ./main.go


#### Target Container
FROM debian_lib_kafka:1.0

# Create workdir in target Container
WORKDIR /workplace

# Copy binary from `build` to target Container
COPY --from=build /building_stage/main /workplace/main
COPY --from=build /building_stage/.env /workplace/.env

# Run command
CMD /workplace/main