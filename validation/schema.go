package validations

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/ppeymann/accounting.git"
	"github.com/xeipuuv/gojsonschema"
)

// error for when validation is not OK
const formatErr string = "%s: is not in correct format or not provided."

// LoadSchema loads json schema files on specified path for given component name
func LoadSchema(path string, input map[string][]byte) (err error) {
	files, err := load(path, []string{
		".git", "/.git", ".gitignore", ".DS_Store", ".idea", "/.idea/", "/.idea",
	})

	if err != nil {
		return err
	}

	for _, file := range files {
		key := strings.TrimLeft(file, path)

		fp := filepath.Clean(file)
		data, err := os.ReadFile(fp)
		if err != nil {
			return err
		}

		key = strings.Replace(key, ".json", "", 1)

		input[key] = data
	}

	return err
}

// load is function for
func load(dir string, ignore []string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
		if e != nil {
			return e
		}

		ignored := false
		for _, item := range ignore {
			if strings.Contains(path, item) {
				ignored = true
			}
		}

		if !ignored {
			fileMod := info.Mode()

			if fileMod.IsRegular() {
				files = append(files, path)
			}
		}

		return nil
	})

	return files, err
}

// validateSchema validate given struct with expected schema
func validateSchema(input interface{}, schema []byte) ([]string, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	sl := gojsonschema.NewStringLoader(string(schema))
	rsl := gojsonschema.NewStringLoader(string(bytes))

	res, err := gojsonschema.Validate(sl, rsl)
	if err != nil {
		return nil, err
	}

	if !res.Valid() {
		var errs []string
		for _, e := range res.Errors() {
			t := e.Type()
			switch t {
			case gojsonschema.KEY_PATTERN:
				errs = append(errs, fmt.Sprintf(formatErr, e.Field()))
			case gojsonschema.KEY_REQUIRED:
				errs = append(errs, fmt.Sprint("%v", e.Description()))
			case gojsonschema.KEY_ENUM:
				errs = append(errs, fmt.Sprintf("%v", strings.Replace(e.Description(), "\"", "'", -1)))
			case "condition_else":
			case "condition_then":
				break
			default:
				errs = append(errs, fmt.Sprintf("%v %s", e.Field(), e.Description()))

			}
		}

		return errs, errors.New("bad request")
	}

	return nil, nil
}

func Validate(input interface{}, schemas map[string][]byte) *accounting.BaseResult {
	val := reflect.ValueOf(input)
	if val.Kind() != reflect.Ptr {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.ErrUnimplementedRequest.Error()},
		}
	}

	name := reflect.Indirect(val).Type().Name()
	schema, ok := schemas[name]
	if !ok {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: []string{accounting.ErrUnimplementedRequest.Error()},
		}
	}

	errs, err := validateSchema(input, schema)
	if err != nil {
		return &accounting.BaseResult{
			Status: http.StatusOK,
			Errors: errs,
		}
	}

	return nil

}
