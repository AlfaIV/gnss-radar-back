// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _CodeReciever_programName(ctx context.Context, field graphql.CollectedField, obj *model.CodeReciever) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CodeReciever_programName(ctx, field)
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
		return obj.ProgramName, nil
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

func (ec *executionContext) fieldContext_CodeReciever_programName(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CodeReciever",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CodeReciever_language(ctx context.Context, field graphql.CollectedField, obj *model.CodeReciever) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CodeReciever_language(ctx, field)
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
		return obj.Language, nil
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
	return ec.marshalNCodeLang2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_CodeReciever_language(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CodeReciever",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type CodeLang does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _CodeReciever_programCode(ctx context.Context, field graphql.CollectedField, obj *model.CodeReciever) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_CodeReciever_programCode(ctx, field)
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
		return obj.ProgramCode, nil
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

func (ec *executionContext) fieldContext_CodeReciever_programCode(_ context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "CodeReciever",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

func (ec *executionContext) unmarshalInputCodeRecieverInput(ctx context.Context, obj interface{}) (model.CodeRecieverInput, error) {
	var it model.CodeRecieverInput
	asMap := map[string]interface{}{}
	for k, v := range obj.(map[string]interface{}) {
		asMap[k] = v
	}

	if _, present := asMap["typeLang"]; !present {
		asMap["typeLang"] = "python"
	}

	fieldsInOrder := [...]string{"token", "typeLang"}
	for _, k := range fieldsInOrder {
		v, ok := asMap[k]
		if !ok {
			continue
		}
		switch k {
		case "token":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("token"))
			data, err := ec.unmarshalNID2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.Token = data
		case "typeLang":
			ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("typeLang"))
			data, err := ec.unmarshalNCodeLang2string(ctx, v)
			if err != nil {
				return it, err
			}
			it.TypeLang = data
		}
	}

	return it, nil
}

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var codeRecieverImplementors = []string{"CodeReciever"}

func (ec *executionContext) _CodeReciever(ctx context.Context, sel ast.SelectionSet, obj *model.CodeReciever) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, codeRecieverImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("CodeReciever")
		case "programName":
			out.Values[i] = ec._CodeReciever_programName(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "language":
			out.Values[i] = ec._CodeReciever_language(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "programCode":
			out.Values[i] = ec._CodeReciever_programCode(ctx, field, obj)
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

func (ec *executionContext) unmarshalNCodeLang2string(ctx context.Context, v interface{}) (string, error) {
	res, err := graphql.UnmarshalString(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNCodeLang2string(ctx context.Context, sel ast.SelectionSet, v string) graphql.Marshaler {
	res := graphql.MarshalString(v)
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
	}
	return res
}

func (ec *executionContext) marshalNCodeReciever2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐCodeReciever(ctx context.Context, sel ast.SelectionSet, v model.CodeReciever) graphql.Marshaler {
	return ec._CodeReciever(ctx, sel, &v)
}

func (ec *executionContext) marshalNCodeReciever2ᚖgithubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐCodeReciever(ctx context.Context, sel ast.SelectionSet, v *model.CodeReciever) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._CodeReciever(ctx, sel, v)
}

func (ec *executionContext) unmarshalNCodeRecieverInput2githubᚗcomᚋGokertᚋgnssᚑradarᚋinternalᚋpkgᚋmodelᚐCodeRecieverInput(ctx context.Context, v interface{}) (model.CodeRecieverInput, error) {
	res, err := ec.unmarshalInputCodeRecieverInput(ctx, v)
	return res, graphql.ErrorOnPath(ctx, err)
}

// endregion ***************************** type.gotpl *****************************
