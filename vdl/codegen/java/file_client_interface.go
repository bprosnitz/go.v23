package java

import (
	"bytes"
	"fmt"
	"log"
	"path"

	"v.io/v23/vdl/compile"
	"v.io/v23/vdl/vdlutil"
)

const clientInterfaceTmpl = `// This file was auto-generated by the veyron vdl tool.
// Source: {{ .Source }}
package {{ .PackagePath }};

{{ .ServiceDoc }}
{{ .AccessModifier }} interface {{ .ServiceName }}Client {{ .Extends }} {
{{ range $method := .Methods }}
    {{/* If this method has multiple return arguments, generate the class. */}}
    {{ if $method.IsMultipleRet }}
    public static class {{ $method.UppercaseMethodName }}Out {
        {{ range $retArg := $method.RetArgs }}
        public {{ $retArg.Type }} {{ $retArg.Name }};
        {{ end }}
    }
    {{ end }}

    {{/* Generate the method signature. */}}
    {{ $method.Doc }}
    {{ $method.AccessModifier }} {{ $method.RetType }} {{ $method.Name }}(final io.v.core.v23.context.VContext context{{ $method.Args }}) throws io.v.core.v23.verror.VException;
    {{ $method.AccessModifier }} {{ $method.RetType }} {{ $method.Name }}(final io.v.core.v23.context.VContext context{{ $method.Args }}, final io.v.core.v23.Options veyronOpts) throws io.v.core.v23.verror.VException;
{{ end }}
}
`

type clientInterfaceArg struct {
	Type string
	Name string
}

type clientInterfaceMethod struct {
	AccessModifier      string
	Args                string
	Doc                 string
	IsMultipleRet       bool
	Name                string
	RetArgs             []clientInterfaceArg
	RetType             string
	UppercaseMethodName string
}

func clientInterfaceNonStreamingOutArg(iface *compile.Interface, method *compile.Method, useClass bool, env *compile.Env) string {
	switch len(method.OutArgs) {
	case 0:
		// "void" or "Void"
		return javaType(nil, useClass, env)
	case 1:
		return javaType(method.OutArgs[0].Type, useClass, env)
	default:
		return javaPath(path.Join(interfaceFullyQualifiedName(iface)+"Client", method.Name+"Out"))
	}
}

func clientInterfaceOutArg(iface *compile.Interface, method *compile.Method, isService bool, env *compile.Env) string {
	if isStreamingMethod(method) && !isService {
		return fmt.Sprintf("io.v.core.v23.vdl.ClientStream<%s, %s, %s>", javaType(method.InStream, true, env), javaType(method.OutStream, true, env), clientInterfaceNonStreamingOutArg(iface, method, true, env))
	}
	return clientInterfaceNonStreamingOutArg(iface, method, false, env)
}

func processClientInterfaceMethod(iface *compile.Interface, method *compile.Method, env *compile.Env) clientInterfaceMethod {
	retArgs := make([]clientInterfaceArg, len(method.OutArgs))
	for i := 0; i < len(method.OutArgs); i++ {
		retArgs[i].Name = vdlutil.ToCamelCase(method.OutArgs[i].Name)
		retArgs[i].Type = javaType(method.OutArgs[i].Type, false, env)
	}
	return clientInterfaceMethod{
		AccessModifier:      accessModifierForName(method.Name),
		Args:                javaDeclarationArgStr(method.InArgs, env, true),
		Doc:                 method.Doc,
		IsMultipleRet:       len(retArgs) > 1,
		Name:                vdlutil.ToCamelCase(method.Name),
		RetArgs:             retArgs,
		RetType:             clientInterfaceOutArg(iface, method, false, env),
		UppercaseMethodName: method.Name,
	}
}

// genJavaClientInterfaceFile generates the Java interface file for the provided
// interface.
func genJavaClientInterfaceFile(iface *compile.Interface, env *compile.Env) JavaFileInfo {
	javaServiceName := toUpperCamelCase(iface.Name)
	methods := make([]clientInterfaceMethod, len(iface.Methods))
	for i, method := range iface.Methods {
		methods[i] = processClientInterfaceMethod(iface, method, env)
	}
	data := struct {
		AccessModifier string
		Extends        string
		Methods        []clientInterfaceMethod
		PackagePath    string
		ServiceDoc     string
		ServiceName    string
		Source         string
	}{
		AccessModifier: accessModifierForName(iface.Name),
		Extends:        javaClientExtendsStr(iface.Embeds),
		Methods:        methods,
		PackagePath:    javaPath(javaGenPkgPath(iface.File.Package.GenPath)),
		ServiceDoc:     javaDoc(iface.Doc),
		ServiceName:    javaServiceName,
		Source:         iface.File.BaseName,
	}
	var buf bytes.Buffer
	err := parseTmpl("client interface", clientInterfaceTmpl).Execute(&buf, data)
	if err != nil {
		log.Fatalf("vdl: couldn't execute struct template: %v", err)
	}
	return JavaFileInfo{
		Name: javaServiceName + "Client.java",
		Data: buf.Bytes(),
	}
}
