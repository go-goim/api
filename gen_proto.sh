#!/bin/bash
# define function
function gen_proto() {
    # scan all files in the giving directory
    # and generate proto files
    for file in $(ls $1)
    do
        # check if the file is a proto file
        if [[ $file == *.proto ]]
        then
            # print the message in red color
            echo -e "\033[31m generating $1/$file \033[0m"
            # generate response.proto file
            protoc --proto_path=. --proto_path=./third_party \
                    --go_out=paths=source_relative:. \
                    --go-grpc_out=paths=source_relative:. \
                    --validate_out=lang=go,paths=source_relative:. \
                    $1/$file
        fi
        # check if the file is a directory and if it is not a hidden directory
        if [[ -d $1/$file && $file != .* ]]
        then
            # call the function recursively
            gen_proto $1/$file
        fi
    done
}

# list all directories in current directory
for dir in $(ls)
do
    # check if the directory is a directory and if it is not a hidden directory
    if [[ -d $dir && $dir != .* ]]
    then
        # check if it is third_party directory
        if [[ $dir == third_party ]]
        then
            # don't call the function recursively
            echo "skipping $dir"
        else
            # call the function recursively
            gen_proto $dir
        fi
    fi
done

echo "done"