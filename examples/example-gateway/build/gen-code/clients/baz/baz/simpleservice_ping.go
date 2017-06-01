// Code generated by thriftrw v1.3.0
// @generated

package baz

import (
	"errors"
	"fmt"
	"github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/baz/base"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SimpleService_Ping_Args struct{}

func (v *SimpleService_Ping_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SimpleService_Ping_Args) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *SimpleService_Ping_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("SimpleService_Ping_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *SimpleService_Ping_Args) Equals(rhs *SimpleService_Ping_Args) bool {
	return true
}

func (v *SimpleService_Ping_Args) MethodName() string {
	return "Ping"
}

func (v *SimpleService_Ping_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SimpleService_Ping_Helper = struct {
	Args           func() *SimpleService_Ping_Args
	IsException    func(error) bool
	WrapResponse   func(*base.BazResponse, error) (*SimpleService_Ping_Result, error)
	UnwrapResponse func(*SimpleService_Ping_Result) (*base.BazResponse, error)
}{}

func init() {
	SimpleService_Ping_Helper.Args = func() *SimpleService_Ping_Args {
		return &SimpleService_Ping_Args{}
	}
	SimpleService_Ping_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SimpleService_Ping_Helper.WrapResponse = func(success *base.BazResponse, err error) (*SimpleService_Ping_Result, error) {
		if err == nil {
			return &SimpleService_Ping_Result{Success: success}, nil
		}
		return nil, err
	}
	SimpleService_Ping_Helper.UnwrapResponse = func(result *SimpleService_Ping_Result) (success *base.BazResponse, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type SimpleService_Ping_Result struct {
	Success *base.BazResponse `json:"success,omitempty"`
}

func (v *SimpleService_Ping_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("SimpleService_Ping_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SimpleService_Ping_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _BazResponse_Read(field.Value)
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SimpleService_Ping_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *SimpleService_Ping_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("SimpleService_Ping_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SimpleService_Ping_Result) Equals(rhs *SimpleService_Ping_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}
	return true
}

func (v *SimpleService_Ping_Result) MethodName() string {
	return "Ping"
}

func (v *SimpleService_Ping_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
