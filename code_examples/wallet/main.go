package main

import (
	"fmt"
	"log"

	"github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	// создать рандомную энтропию
	entropy, err := hdwallet.NewEntropy(128)
	if err != nil {
		log.Fatal(err)
	}

	// создать комбинацию слов на основе энтропии
	mnemonic, err := hdwallet.NewMnemonicFromEntropy(entropy)
	if err != nil {
		log.Fatal(err)
	}

	// можно перезаписать своей комбинацией
	// mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	
	// создать кошелек на основе комбинации слов
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947
}
