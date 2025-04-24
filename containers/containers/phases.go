package main

import "fmt"

type (
	phase_name     string
	phase_path     string
	phase_args     map[string]string
	phase_function func(phase_args) error

	phase struct {
		name     phase_name
		path     phase_path
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
	Build_Path  phase_path = "build/build"
)

func CreatePhase(args []string) (Phase, error) {
	var err error

	var p_name phase_name
	if p_name, err = get_phase_name(args[1]); err != nil {
		return nil, err
	}

	var p_path phase_path
	if p_path, err = get_phase_path(p_name); err != nil {
		return nil, err
	}

	var p_args phase_args
	if p_args, err = get_phase_args(p_name, args[2:]); err != nil {
		return nil, err
	}

	var p_function phase_function
	if p_function, err = get_phase_function(p_name); err != nil {
		return nil, err
	}

	return phase{
		name:     p_name,
		path:     p_path,
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

func get_phase_path(p phase_name) (phase_path, error) {
	switch p {

	case Build_Phase:
		return Build_Path, nil

	default:
		return phase_path(""), fmt.Errorf("%s: have an associated path", p)
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
		return BuildProcess, nil

	default:
		return nil, fmt.Errorf("%s: doesn't have an associated function", p)
	}
}

func (p phase) Path() string {
	return string(p.path)
}

func (p phase) RunProcess() error {
	return p.function(p.args)
}
