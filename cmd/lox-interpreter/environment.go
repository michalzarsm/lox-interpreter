package main

import "errors"

type Environment struct {
	enclosing *Environment
	values    map[string]any
}

func newEnvironment(enclosing *Environment) *Environment {
	return &Environment{
		enclosing: enclosing,
		values:    make(map[string]any),
	}
}

func (e *Environment) define(name string, value any) {
	e.values[name] = value
}

func (e *Environment) get(name Token) (any, error) {
	val, ok := e.values[name.Lexeme]
	if ok {
		return val, nil
	}

	if e.enclosing != nil {
		return e.enclosing.get(name)
	}

	return nil, errors.New("Undefinded variable '" + name.Lexeme + "'.")
}

func (e *Environment) assign(name Token, value any) error {
	_, ok := e.values[name.Lexeme]

	if ok {
		e.values[name.Lexeme] = value
		return nil
	}

	if e.enclosing != nil {
		return e.enclosing.assign(name, value)
	}

	return errors.New("Undefinded variable '" + name.Lexeme + "'.")
}
