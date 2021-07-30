# Simple k8s secret parser

A simple script to encode and decode variables defined in a secret configuration file from plain text to base64 and base64 to plain text.


go run main.go -f <file_name.yaml> -p d # decode
 
go run main.go -f <file_name.yaml> -p e # enconde