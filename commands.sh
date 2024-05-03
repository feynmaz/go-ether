solc --bin --abi contract/todo.sol -o build  --overwrite
abigen --bin=build/Todo.bin --abi=build/Todo.abi -pkg=todo --out=todo/todo.go