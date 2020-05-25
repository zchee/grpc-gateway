package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway-2/internal/gengateway"
	"google.golang.org/protobuf/compiler/protogen"
)

// Variables set by goreleaser at build time
var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	if len(os.Args) == 2 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Fprintf(os.Stderr, "%v %v, commit %v, built at %v\n", filepath.Base(os.Args[0]), version, commit, date)
		os.Exit(1)
	}

	var (
		flags flag.FlagSet
		/*
			importPrefix               = flag.String("import_prefix", "", "prefix to be added to go package paths for imported proto files")
			importPath                 = flag.String("import_path", "", "used as the package if no input files declare go_package. If it contains slashes, everything up to the rightmost slash is ignored.")
			registerFuncSuffix         = flag.String("register_func_suffix", "Handler", "used to construct names of generated Register*<Suffix> methods.")
			useRequestContext          = flag.Bool("request_context", true, "determine whether to use http.Request's context or not")
			allowDeleteBody            = flag.Bool("allow_delete_body", false, "unless set, HTTP DELETE methods may not have a body")
			grpcAPIConfiguration       = flag.String("grpc_api_configuration", "", "path to gRPC API Configuration in YAML format")
			pathType                   = flag.String("paths", "", "specifies how the paths of generated files are structured")
			allowRepeatedFieldsInBody  = flag.Bool("allow_repeated_fields_in_body", false, "allows to use repeated field in `body` and `response_body` field of `google.api.http` annotation option")
			repeatedPathParamSeparator = flag.String("repeated_path_param_separator", "csv", "configures how repeated fields should be split. Allowed values are `csv`, `pipes`, `ssv` and `tsv`.")
			allowPatchFeature          = flag.Bool("allow_patch_feature", true, "determines whether to use PATCH feature involving update masks (using google.protobuf.FieldMask).")
			standalone                 = flag.Bool("standalone", false, "generates a standalone gateway package, which imports the target service package")
		*/
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				gengateway.GenerateFile(gen, f)
			}
		}
		gen.SupportedFeatures = gengateway.SupportedFeatures
		return nil
	})
}
