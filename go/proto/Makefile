generate_grpc_code_go:
	protoc \
    --go_out=..\script_server \
    --go_opt=paths=source_relative \
    --go-grpc_out=..\script_server \
    --go-grpc_opt=paths=source_relative \
    script_server.proto

generate_grpc_code_python:
	protoc \
	--proto_path= \
	--python_out=py_script_server \
	--grpc_python_out=py_script_server \
	script_server.proto

generate_grpc_code_python2:
	py \
	-m grpc_tools.protoc \
	--proto_path=. \
	--python_out=py_script_server \
	--pyi_out=py_script_server \
	--grpc_python_out=py_script_server \
	script_server.proto