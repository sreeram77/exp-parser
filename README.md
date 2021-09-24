# exp-parser

## TL;DR
exp-parser is a YAML/JSON parser which evaluates expressions on JSON input.

## Description
exp-parser evaluates expressions against a given JSON input. The app by-default reads the input from input.yml and writes output to output.yml. This can be changed by modifying env variables INPUT_PATH and OUTPUT_PATH. It also prints the output.

The parser can handle operations such as 
1. Brackets "()"
2. Logical "AND", "OR", "NOT"
3. Comparison "==" and "EXISTS"

Data types such as string, int, float and bool can be handled.

## Future expansion
exp-parser reads the file line-by-line and stores it in memory. The expression is parsed and then output is written into a new file. Since input read is line-by-line, it can be modified to limit read based on a buffer size.

## Assumptions
1. "NOT" works only along with "EXISTS"

## Installation
### Compile and run locally
```shell
go run main.go
```

### Docker
```shell
docker run -it -v path/to/local/input.yml:/file/input.yml -v path/to/local/output.yml:/file/output.yml sreeram777/exp-parser:latest
```
