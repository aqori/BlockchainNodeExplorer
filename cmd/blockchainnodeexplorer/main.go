// cmd/blockchainnodeexplorer/main.go
package main

import (
"flag"
"log"
"os"

"blockchainnodeexplorer/internal/blockchainnodeexplorer"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := blockchainnodeexplorer.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
