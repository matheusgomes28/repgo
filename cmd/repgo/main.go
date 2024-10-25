package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
	interprerter "github.com/matheusgomes28/repgo/interpreter"
	lua "github.com/yuin/gopher-lua"

	log "github.com/sirupsen/logrus"
)

func load_handler_scripts() (map[string]*interprerter.LuaHandler, error) {
	// Load interpreter script
	scripts := map[string]string{
		"add": "/home/matheus/development/repgo/script/add.lua",
		"sub": "/home/matheus/development/repgo/script/sub.lua",
	}

	handler_map := make(map[string]*interprerter.LuaHandler)
	for command_name, script_path := range scripts {
		// Load the
		lua_state := lua.NewState()
		// lua_state := lua.NewState(lua.Options{
		// 	RegistrySize:     1024 * 20, // this is the initial size of the registry
		// 	RegistryMaxSize:  1024 * 80, // this is the maximum size that the registry can grow to. If set to `0` (the default) then the registry will not auto grow
		// 	RegistryGrowStep: 32,        // this is how much to step up the registry by each time it runs out of space. The default is `32`.
		// 	CallStackSize:    120,
		// })
		// defer lua_state.Close()

		if err := lua_state.DoFile(script_path); err != nil {
			log.Error(fmt.Sprintf("could not load lua_stateua script file %v", script_path))
			return map[string]*interprerter.LuaHandler{}, err
		}

		handler_map[command_name] = &interprerter.LuaHandler{
			Name: command_name,
			Fn:   lua_state,
		}
	}

	return handler_map, nil
}

func main() {
	f, err := os.OpenFile("latest_run.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("could not open file for logging")
		os.Exit(-1)
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFormatter(&log.JSONFormatter{})

	// Interpreter handlers
	handlers, err := load_handler_scripts()
	if err != nil {
		log.Errorf("could not load the handlers: %v", err)
		os.Exit(-1)
	}

	interprerter := interprerter.Interpreter{
		Handlers: handlers,
	}

	// Interpreter loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		indicator := color.GreenString(">>")
		fmt.Printf("%s ", indicator)

		scanner.Scan() // use `for scanner.Scan()` to keep reading
		command_str := scanner.Text()

		// Try and execute the given command with
		// the arguments provided
		result, err := interprerter.Execute(command_str)

		// If we failed to execute the command, then
		// we simply print the error and continue
		// executation
		if err != nil {
			error_str := color.RedString("[ERROR]")
			fmt.Printf("%s %v\n", error_str, err)
		}

		result_indicator := color.GreenString(">>>>")
		fmt.Printf("%s %s\n", result_indicator, result)
	}

}
