#!/bin/bash

# Define all projects
all_projects=(ingest parse/emojis parse/hashtags parse/urls storage/cassandra storage/kafka)

cmd=$1
proj=$2

projects=()
if [[ -z "$proj" ]];
then
    projects=${all_projects[*]}
else
    projects=("$proj")
fi

for project in ${projects[*]}
do
    wd=$(pwd)
    cd "$project"
    case "$cmd" in
        'build')
            echo "Building $project"
            CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o "$wd"/bin/"$project"-linux .
            GOOS=darwin GOARCH=386 go build -o "$wd"/bin/"$project"-osx .
        ;;
        'test')
            echo "Testing $project"
            go test -v ./...
        ;;
        'vet')
            echo "Vetting $project"
            go tool vet .
        ;;
    esac
    cd -
done
