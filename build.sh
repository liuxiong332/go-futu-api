find ./futu-proto -type f -name '*proto' -exec protoc  -I ./futu-proto --go_out=./pb {} \;
