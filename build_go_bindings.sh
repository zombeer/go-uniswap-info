#!/bin/bash
mkdir gen;
mkdir build;
for i in ./v2-core/contracts/interfaces/*.sol; do
	echo Processing $i...;
	solc --bin --abi $i -o build --overwrite;
	fn=${i##*/};
	fn_only=${fn%.sol};
	mkdir ./gen/${fn_only};
	abigen --bin=./build/${fn_only}.bin --abi=./build/${fn_only}.abi --pkg=${fn_only} --out=./gen/${fn_only}/${fn_only}.go
done