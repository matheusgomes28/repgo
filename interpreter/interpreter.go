package interprerter

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	lua "github.com/yuin/gopher-lua"
)

type LuaHandler struct {
	Name string
	Fn   *lua.LState
	// TODO : Arguments {arg1 = int, arg2 = string}
}

type Interpreter struct {
	// It will store which Lua scripts will be handling
	// a particular input function from the user
	// e.g. >> add 5 6 -> Handlers["add"] will be the lua
	// script running the function
	Handlers map[string]*LuaHandler
}

// func (i *Interpreter) Attach(lua_script string, command_name) error {}

func (i *Interpreter) Execute(command string) (string, error) {

	// Take in a commandstring e.g. "add 5 6",
	// Then find the associated handler Lua script
	// Then parse the inputs, make sure they're what we expect
	// Then execute the lua script passing the inputs

	// TODO : Parse the command_name from command
	split := strings.Split(command, " ")
	if len(split) < 1 {
		log.Error("could not split command string into at least one element")
		return "", fmt.Errorf("could not get command name")
	}

	command_name := split[0]
	handler, ok := i.Handlers[command_name]
	if !ok {
		log.Error("could not find the command handler name in the interpreter handler map")
		return "", fmt.Errorf("could not find command name handler")
	}

	// TODO : Parse the remaining arguments, making sure
	// TODO : that we have the number of args expected and
	// TODO : the types of arguments are valid
	// e.g. command_str = "add hello 4" -> should fail as arg1 isn't numbe
	log.Info(fmt.Sprintf("we received an %s command, handling with script %s", command_name, handler.Name))

	// TODO : Create the argument string from the remaining
	// TODO : split of the command string, e.g. "1, 5, 6, 2"
	argument_str := strings.Join(split[1:], " ")

	lua_command := fmt.Sprintf("%s::Main(%s)", handler.Name, argument_str)
	log.Infof("executing LUA: `%s`", lua_command)

	// TODO : remove testing code
	test_var := handler.Fn.GetGlobal("Main")
	log.Infof("testing: %s", test_var.String())

	err := handler.Fn.CallByParam(lua.P{
		Fn:      handler.Fn.GetGlobal("Main"),
		NRet:    1,
		Protect: false,
	}, lua.LString(argument_str))

	if err != nil {
		log.Errorf("could not execute the Lua handler: %v", err)
		return "", err
	}

	ret := handler.Fn.Get(-1)
	handler.Fn.Pop(1)

	if str, ok := ret.(lua.LString); ok {
		// lv is LString
		log.Infof("returning the value `%s` from the handler", string(str))
		return string(str), nil
	}

	return "", fmt.Errorf("could not cast return type to a string")
}
