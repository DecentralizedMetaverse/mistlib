package content

import (
	"fmt"
	"os"
)

func RunCommand(name string, args []string) {
	commandHandler := map[string]func([]string){
		"init":            handleInit,
		"switch":          handleSwitch,
		"get":             handleGet,
		"put":             handleAdd,
		"set-password":    handleSetPassword,
		"cat":             handleCat,
		"get-world-cid":   handleGetWorldCID,
		"download-world":  handleDownloadWorld,
		"get-world-info":  handleGetWorldInfo,
		"unpack":          handleUnpack,
		"set-custom-data": handleSetCustomData,
		"set-parent":      handleSetParent,
		"update":          handleUpdate,
	}

	if handler, found := commandHandler[name]; found {
		handler(args)
	} else {
		fmt.Printf("Unknown command: %s\n", name)
		os.Exit(1)
	}
}
