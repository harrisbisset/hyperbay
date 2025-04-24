package main

import "fmt"

type (
	phase_name     string
	phase_args     map[string]string
	phase_function func(phase_args) error

	phase struct {
		name     phase_name
		args     phase_args
		function phase_function
	}

	Phase interface {
		Path() string
		RunProcess() error
	}
)

const (
	Build_Phase phase_name = "build"
)

func CreatePhase(args []string) (Phase, error) {
	var err error

	var p_name phase_name
	if p_name, err = get_phase_name(args[0]); err != nil {
		return nil, err
	}

	var p_args phase_args
	if p_args, err = get_phase_args(p_name, args[1:]); err != nil {
		return nil, err
	}

	var p_function phase_function
	if p_function, err = get_phase_function(p_name); err != nil {
		return nil, err
	}

	return phase{
		name:     p_name,
		args:     p_args,
		function: p_function,
	}, nil
}

func get_phase_name(s string) (phase_name, error) {
	switch phase_name(s) {

	case Build_Phase:
		return Build_Phase, nil

	default:
		return phase_name(""), fmt.Errorf("%s: is not a phase", s)
	}
}

func get_phase_args(p phase_name, _ []string) (phase_args, error) {
	switch p {

	case Build_Phase:
		return nil, nil

	default:
		return nil, fmt.Errorf("%s: doesn't handle arguements", p)
	}
}

func get_phase_function(p phase_name) (phase_function, error) {
	switch p {

	case Build_Phase:
		return nil, nil

	default:
		return nil, fmt.Errorf("%s: doesn't have an associated function", p)
	}
}

func (p phase) Path() string {
	switch p.name {
	case Build_Phase:
		return "build/build"
	default:
		panic(fmt.Sprintf("%s: phase doesn't have an associated path", p.name))
	}
}

func (p phase) RunProcess() error {
	return p.function(p.args)
}
