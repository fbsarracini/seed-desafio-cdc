package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fbsarracini/seed-desafio-cdc/internal/metrics"
)

func main() {
	dir := flag.String("dir", ".", "diretório raiz do projeto")
	module := flag.String("module", "fbsarracini/seed-desafio-cdc", "nome do módulo Go")
	flag.Parse()

	pm, err := metrics.AnalyzeDirectory(*dir, *module)
	if err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		os.Exit(1)
	}

	if len(pm.Files) == 0 {
		fmt.Fprintln(os.Stderr, "nenhum arquivo .go encontrado em", *dir)
		os.Exit(1)
	}

	fmt.Println(pm.FormatMarkdown())
}
