// Code generated by thriftrw v1.3.0
// @generated

package baz

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SecondService_Echo_Args struct {
	Arg string `json:"arg,required"`
}

func (v *SecondService_Echo_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	w, err = wire.NewValueString(v.Arg), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SecondService_Echo_Args) FromWire(w wire.Value) error {
	var err error
	argIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Arg, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}
	if !argIsSet {
		return errors.New("field Arg of SecondService_Echo_Args is required")
	}
	return nil
}

func (v *SecondService_Echo_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++
	return fmt.Sprintf("SecondService_Echo_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *SecondService_Echo_Args) Equals(rhs *SecondService_Echo_Args) bool {
	if !(v.Arg == rhs.Arg) {
		return false
	}
	return true
}

func (v *SecondService_Echo_Args) MethodName() string {
	return "Echo"
}

func (v *SecondService_Echo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SecondService_Echo_Helper = struct {
	Args           func(arg string) *SecondService_Echo_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*SecondService_Echo_Result, error)
	UnwrapResponse func(*SecondService_Echo_Result) error
}{}

func init() {
	SecondService_Echo_Helper.Args = func(arg string) *SecondService_Echo_Args {
		return &SecondService_Echo_Args{Arg: arg}
	}
	SecondService_Echo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SecondService_Echo_Helper.WrapResponse = func(err error) (*SecondService_Echo_Result, error) {
		if err == nil {
			return &SecondService_Echo_Result{}, nil
		}
		return nil, err
	}
	SecondService_Echo_Helper.UnwrapResponse = func(result *SecondService_Echo_Result) (err error) {
		return
	}
}

type SecondService_Echo_Result struct{}

func (v *SecondService_Echo_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SecondService_Echo_Result) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *SecondService_Echo_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("SecondService_Echo_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SecondService_Echo_Result) Equals(rhs *SecondService_Echo_Result) bool {
	return true
}

func (v *SecondService_Echo_Result) MethodName() string {
	return "Echo"
}

func (v *SecondService_Echo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
