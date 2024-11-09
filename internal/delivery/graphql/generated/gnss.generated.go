// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type GnssMutationsResolver interface {
	UpsetDevice(ctx context.Context, obj *model.GnssMutations, input model.UpsetDeviceInput) (*model.UpsetDeviceOutput, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_GnssMutations_upsetDevice_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.UpsetDeviceInput
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNUpsetDeviceInput2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐUpsetDeviceInput(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _CoordsResults_x(ctx context.Context, field graphql.CollectedField, obj *model.CoordsResults) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CoordsResults_x(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.X, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CoordsResults_x(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CoordsResults",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CoordsResults_y(ctx context.Context, field graphql.CollectedField, obj *model.CoordsResults) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CoordsResults_y(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Y, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CoordsResults_y(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CoordsResults",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CoordsResults_z(ctx context.Context, field graphql.CollectedField, obj *model.CoordsResults) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CoordsResults_z(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Z, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CoordsResults_z(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CoordsResults",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _DevicePagination_items(ctx context.Context, field graphql.CollectedField, obj *model.DevicePagination) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_DevicePagination_items(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Items, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*model.Device)
	fc.Result = res
	return ec.marshalODevice2ᚕᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐDeviceᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_DevicePagination_items(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "DevicePagination",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Device_id(ctx, field)
			case "name":
				return ec.fieldContext_Device_name(ctx, field)
			case "token":
				return ec.fieldContext_Device_token(ctx, field)
			case "description":
				return ec.fieldContext_Device_description(ctx, field)
			case "Coords":
				return ec.fieldContext_Device_Coords(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Device", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _GNSS_Id(ctx context.Context, field graphql.CollectedField, obj *model.Gnss) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_GNSS_Id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_GNSS_Id(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "GNSS",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _GNSS_SatelliteId(ctx context.Context, field graphql.CollectedField, obj *model.Gnss) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_GNSS_SatelliteId(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SatelliteID, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_GNSS_SatelliteId(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "GNSS",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _GNSS_SatelliteName(ctx context.Context, field graphql.CollectedField, obj *model.Gnss) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_GNSS_SatelliteName(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.SatelliteName, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_GNSS_SatelliteName(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "GNSS",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _GNSS_Coordinates(ctx context.Context, field graphql.CollectedField, obj *model.Gnss) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_GNSS_Coordinates(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Coordinates, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.CoordsResults)
	fc.Result = res
	return ec.marshalNCoordsResults2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐCoordsResults(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_GNSS_Coordinates(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "GNSS",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "x":
				return ec.fieldContext_CoordsResults_x(ctx, field)
			case "y":
				return ec.fieldContext_CoordsResults_y(ctx, field)
			case "z":
				return ec.fieldContext_CoordsResults_z(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type CoordsResults", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _GNSSPagination_items(ctx context.Context, field graphql.CollectedField, obj *model.GNSSPagination) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_GNSSPagination_items(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Items, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.([]*model.Gnss)
	fc.Result = res
	return ec.marshalOGNSS2ᚕᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGnssᚄ(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_GNSSPagination_items(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "GNSSPagination",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "Id":
				return ec.fieldContext_GNSS_Id(ctx, field)
			case "SatelliteId":
				return ec.fieldContext_GNSS_SatelliteId(ctx, field)
			case "SatelliteName":
				return ec.fieldContext_GNSS_SatelliteName(ctx, field)
			case "Coordinates":
				return ec.fieldContext_GNSS_Coordinates(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type GNSS", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _GnssMutations_upsetDevice(ctx context.Context, field graphql.CollectedField, obj *model.GnssMutations) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_GnssMutations_upsetDevice(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.GnssMutations().UpsetDevice(rctx, obj, fc.Args["input"].(model.UpsetDeviceInput))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.UpsetDeviceOutput)
	fc.Result = res
	return ec.marshalNUpsetDeviceOutput2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐUpsetDeviceOutput(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_GnssMutations_upsetDevice(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "GnssMutations",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "device":
				return ec.fieldContext_UpsetDeviceOutput_device(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type UpsetDeviceOutput", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_GnssMutations_upsetDevice_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _UpsetDeviceOutput_device(ctx context.Context, field graphql.CollectedField, obj *model.UpsetDeviceOutput) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_UpsetDeviceOutput_device(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Device, nil
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Device)
	fc.Result = res
	return ec.marshalNDevice2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐDevice(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_UpsetDeviceOutput_device(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "UpsetDeviceOutput",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Device_id(ctx, field)
			case "name":
				return ec.fieldContext_Device_name(ctx, field)
			case "token":
				return ec.fieldContext_Device_token(ctx, field)
			case "description":
				return ec.fieldContext_Device_description(ctx, field)
			case "Coords":
				return ec.fieldContext_Device_Coords(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Device", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputCoordsInput(ctx context.Context, obj interface{}) (model.CoordsInput, error) {
	var it model.CoordsInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"x", "y", "z"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "x":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("x"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.X = data
		case "y":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("y"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Y = data
		case "z":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("z"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Z = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputDeviceFilter(ctx context.Context, obj interface{}) (model.DeviceFilter, error) {
	var it model.DeviceFilter
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"Ids", "Names", "Tokens"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "Ids":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Ids"))
			data, err := ec.unmarshalOString2ᚕstringᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Ids = data
		case "Names":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Names"))
			data, err := ec.unmarshalOString2ᚕstringᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Names = data
		case "Tokens":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Tokens"))
			data, err := ec.unmarshalOString2ᚕstringᚄ(ctx, v)
			if err != nil {
				return it, err
			}
			it.Tokens = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputGNSSFilter(ctx context.Context, obj interface{}) (model.GNSSFilter, error) {
	var it model.GNSSFilter
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"Coordinates"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "Coordinates":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Coordinates"))
			data, err := ec.unmarshalNCoordsInput2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐCoordsInput(ctx, v)
			if err != nil {
				return it, err
			}
			it.Coordinates = data
		}
	}

	return it, nil
}

func (ec *executionContext) unmarshalInputUpsetDeviceInput(ctx context.Context, obj interface{}) (model.UpsetDeviceInput, error) {
	var it model.UpsetDeviceInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	fieldsInOrder := [...]string{"Name", "Token", "Description", "Coords"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "Name":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Name"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Name = data
		case "Token":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Token"))
			data, err := ec.unmarshalNString2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Token = data
		case "Description":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Description"))
			data, err := ec.unmarshalOString2ᚖstring(ctx, v)
			if err != nil {
				return it, err
			}
			it.Description = data
		case "Coords":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("Coords"))
			data, err := ec.unmarshalNCoordsInput2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐCoordsInput(ctx, v)
			if err != nil {
				return it, err
			}
			it.Coords = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var coordsResultsImplementors = []string{"CoordsResults"}

func (ec *executionContext) _CoordsResults(ctx context.Context, sel ast.SelectionSet, obj *model.CoordsResults) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, coordsResultsImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CoordsResults")
		case "x":
			out.Values[i] = ec._CoordsResults_x(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "y":
			out.Values[i] = ec._CoordsResults_y(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "z":
			out.Values[i] = ec._CoordsResults_z(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var devicePaginationImplementors = []string{"DevicePagination"}

func (ec *executionContext) _DevicePagination(ctx context.Context, sel ast.SelectionSet, obj *model.DevicePagination) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, devicePaginationImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("DevicePagination")
		case "items":
			out.Values[i] = ec._DevicePagination_items(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var gNSSImplementors = []string{"GNSS"}

func (ec *executionContext) _GNSS(ctx context.Context, sel ast.SelectionSet, obj *model.Gnss) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, gNSSImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("GNSS")
		case "Id":
			out.Values[i] = ec._GNSS_Id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "SatelliteId":
			out.Values[i] = ec._GNSS_SatelliteId(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "SatelliteName":
			out.Values[i] = ec._GNSS_SatelliteName(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "Coordinates":
			out.Values[i] = ec._GNSS_Coordinates(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var gNSSPaginationImplementors = []string{"GNSSPagination"}

func (ec *executionContext) _GNSSPagination(ctx context.Context, sel ast.SelectionSet, obj *model.GNSSPagination) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, gNSSPaginationImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("GNSSPagination")
		case "items":
			out.Values[i] = ec._GNSSPagination_items(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var gnssMutationsImplementors = []string{"GnssMutations"}

func (ec *executionContext) _GnssMutations(ctx context.Context, sel ast.SelectionSet, obj *model.GnssMutations) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, gnssMutationsImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("GnssMutations")
		case "upsetDevice":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._GnssMutations_upsetDevice(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var upsetDeviceOutputImplementors = []string{"UpsetDeviceOutput"}

func (ec *executionContext) _UpsetDeviceOutput(ctx context.Context, sel ast.SelectionSet, obj *model.UpsetDeviceOutput) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, upsetDeviceOutputImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("UpsetDeviceOutput")
		case "device":
			out.Values[i] = ec._UpsetDeviceOutput_device(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNCoordsInput2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐCoordsInput(ctx context.Context, v interface{}) (*model.CoordsInput, error) {
	res, err := ec.unmarshalInputCoordsInput(ctx, v)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNCoordsResults2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐCoordsResults(ctx context.Context, sel ast.SelectionSet, v model.CoordsResults) graphql.Marshaler {
	return ec._CoordsResults(ctx, sel, &v)
}

func (ec *executionContext) marshalNCoordsResults2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐCoordsResults(ctx context.Context, sel ast.SelectionSet, v *model.CoordsResults) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CoordsResults(ctx, sel, v)
}

func (ec *executionContext) unmarshalNDeviceFilter2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐDeviceFilter(ctx context.Context, v interface{}) (model.DeviceFilter, error) {
	res, err := ec.unmarshalInputDeviceFilter(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNDevicePagination2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐDevicePagination(ctx context.Context, sel ast.SelectionSet, v model.DevicePagination) graphql.Marshaler {
	return ec._DevicePagination(ctx, sel, &v)
}

func (ec *executionContext) marshalNDevicePagination2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐDevicePagination(ctx context.Context, sel ast.SelectionSet, v *model.DevicePagination) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._DevicePagination(ctx, sel, v)
}

func (ec *executionContext) marshalNGNSS2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGnss(ctx context.Context, sel ast.SelectionSet, v *model.Gnss) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._GNSS(ctx, sel, v)
}

func (ec *executionContext) unmarshalNGNSSFilter2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGNSSFilter(ctx context.Context, v interface{}) (model.GNSSFilter, error) {
	res, err := ec.unmarshalInputGNSSFilter(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNGNSSPagination2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGNSSPagination(ctx context.Context, sel ast.SelectionSet, v model.GNSSPagination) graphql.Marshaler {
	return ec._GNSSPagination(ctx, sel, &v)
}

func (ec *executionContext) marshalNGNSSPagination2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGNSSPagination(ctx context.Context, sel ast.SelectionSet, v *model.GNSSPagination) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._GNSSPagination(ctx, sel, v)
}

func (ec *executionContext) marshalNGnssMutations2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGnssMutations(ctx context.Context, sel ast.SelectionSet, v model.GnssMutations) graphql.Marshaler {
	return ec._GnssMutations(ctx, sel, &v)
}

func (ec *executionContext) marshalNGnssMutations2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGnssMutations(ctx context.Context, sel ast.SelectionSet, v *model.GnssMutations) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._GnssMutations(ctx, sel, v)
}

func (ec *executionContext) unmarshalNUpsetDeviceInput2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐUpsetDeviceInput(ctx context.Context, v interface{}) (model.UpsetDeviceInput, error) {
	res, err := ec.unmarshalInputUpsetDeviceInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNUpsetDeviceOutput2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐUpsetDeviceOutput(ctx context.Context, sel ast.SelectionSet, v model.UpsetDeviceOutput) graphql.Marshaler {
	return ec._UpsetDeviceOutput(ctx, sel, &v)
}

func (ec *executionContext) marshalNUpsetDeviceOutput2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐUpsetDeviceOutput(ctx context.Context, sel ast.SelectionSet, v *model.UpsetDeviceOutput) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._UpsetDeviceOutput(ctx, sel, v)
}

func (ec *executionContext) marshalOGNSS2ᚕᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGnssᚄ(ctx context.Context, sel ast.SelectionSet, v []*model.Gnss) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNGNSS2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐGnss(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

// endregion ***************************** type.gotpl *****************************
