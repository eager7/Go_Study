contract:=erc20

default:
	solc --abi ${contract}.sol | awk '/JSON ABI/{x=1;next}x' > ${contract}.abi
	solc --bin ${contract}.sol | awk '/Binary:/{x=1;next}x' > ${contract}.bin
	abigen --bin=${contract}.bin --abi=${contract}.abi --pkg=${contract} --out=${contract}.go

clean:
	rm ${contract}.abi
	rm ${contract}.bin
	rm ${contract}.go