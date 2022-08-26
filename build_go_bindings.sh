#!/bin/bash
mkdir gen;
mkdir build;
for i in ./interfaces/*.sol; do
	echo Processing $i...;
	solc --bin --abi $i -o build --overwrite;
	filename=${i##*/};
	fn_only=${filename%.sol};
	abigen --bin=./build/${fn_only}.bin --abi=./build/${fn_only}.abi --pkg=${fn_only} --out=./gen/${fn_only}.go
done