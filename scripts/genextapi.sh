#!/bin/bash

parse_args() {
  while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
      -p|--path) # путь до директории с api слоем
        api_folder_path="$2"
        shift 2
        ;;
      -s|--swagger) # путь до swagger.yml файла
        swagger_file_path="$2"
        shift 2
        ;;
      -h|--help)
        echo "-p --path     - путь до директории с api слоем приложения. Все сгенерированные файлы попадут именно туда"
        echo "-s --swagger  - путь до файла swagger.yml"
        exit 1
        ;;
      *)
        echo "Unknown option: ${key}" >&2
        exit 1
        ;;
    esac
  done
}

# Call the function with the command-line arguments
parse_args "$@"

if [ ! -d "$api_folder_path" ]; then
   echo "api folder path does not exist."
   exit 1
fi

if [ ! -f "$swagger_file_path" ]; then
   echo "swagger.yml does not exist in path: $swagger_file_path"
   exit 1
fi

if ! [ -x "$(command -v docker)" ]; then
  echo 'Error: "docker" is not installed.' >&2
  exit 1
fi

api_folder="$api_folder_path/restapi"


set -x -e -o pipefail

rm -rf $api_folder
mkdir $api_folder


#shopt -s expand_aliases
#alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"
swagger generate server -t $api_folder -f $swagger_file_path --exclude-main
swagger generate client -t $api_folder -f $swagger_file_path
