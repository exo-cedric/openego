package cmd

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type Flag int
const (
	FLAG_INVALID Flag = iota  // 0
	FLAG_ZONE         = 1 << iota  // 2^1
)

const (
	FLAG_DEFAULT_ZONE = "ch-dk-2"
)

func AddCanonicalFlags(cmd *cobra.Command, flags int) {
	if flags & FLAG_ZONE != 0 {
		cmd.Flags().StringP("zone", "z", FLAG_DEFAULT_ZONE, "Zone")
	}
}

func AddOpenApiFlags(cmd *cobra.Command, params interface{}, defaults *map[string]string) {
	// Parameters must be a Pointer to a Struct
	// (to allow reflection to set values in ParseOpenApiFlags)
  params_type := reflect.TypeOf(params)
  params_kind := params_type.Kind()
	if params_kind != reflect.Ptr {
		log.Printf("[cmd.AddOpenApiFlags] ERROR: Parameters are no Pointer; %s", params_kind)
		return
	}
  params_elem_kind := params_type.Elem().Kind()
	if params_elem_kind != reflect.Struct {
		log.Printf("[cmd.AddOpenApiFlags] ERROR: Parameters are no Struct; %s", params_elem_kind)
		return
	}

	// Loop through fields
	for i := 0; i < params_type.Elem().NumField(); i++ {
		field := params_type.Elem().Field(i)

		// Field "flag"
		field_json := field.Tag.Get("json")
		if field_json == "" {
			continue
		}
		field_flag := strings.SplitN(field_json, ",", 2)[0]

		// Field description
		field_desc := field.Tag.Get("description")
		if field_desc == "" {
			field_desc = field.Name
		}

		// Field default
		field_dflt := field.Tag.Get("default")
		if defaults != nil {
			if user_dflt, exists := (*defaults)[field_flag]; exists {
				field_dflt = user_dflt
			}
		}

		// Add flag
		field_kind := field.Type.Kind()
		if field_kind == reflect.Ptr {
			field_kind = field.Type.Elem().Kind()
		}
		switch field_kind {

		case reflect.Bool,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
			reflect.Float32, reflect.Float64,
			reflect.String:
			// Let's handle all flags as string (which allow to detect unset - empty - flags naturally)
			cmd.Flags().String(field_flag, field_dflt, fmt.Sprintf("%s (%s)", field_desc, field_kind))

		default:
			log.Printf("[cmd.AddOpenApiFlags] ERROR: Unsupported field type; %s (%s)", field_flag, field_kind)

		}
	}
}

func ParseOpenApiFlags(cmd *cobra.Command, params interface{}) {
	// Parameters must be a Pointer to a Struct
	// (to allow reflection to set values in ParseOpenApiFlags)
  params_type := reflect.TypeOf(params)
  params_value := reflect.ValueOf(params)
  params_kind := params_type.Kind()
	if params_kind != reflect.Ptr {
		log.Printf("[cmd.ParseOpenApiFlags] ERROR: Parameters are no Pointer; %s", params_kind)
		return
	}
  params_elem_kind := params_type.Elem().Kind()
	if params_elem_kind != reflect.Struct {
		log.Printf("[cmd.ParseOpenApiFlags] ERROR: Parameters are no Struct; %s", params_elem_kind)
		return
	}

	// Loop through fields
	for i := 0; i < params_type.Elem().NumField(); i++ {
		field := params_value.Elem().Field(i)
		field_type := params_type.Elem().Field(i)

		// Field "flag"
		field_json := field_type.Tag.Get("json")
		if field_json == "" {
			continue
		}
		field_flag := strings.SplitN(field_json, ",", 2)[0]

		// Parse flag
		flag_value, err := cmd.Flags().GetString(field_flag)
		if err != nil {
			log.Printf("[cmd.ParseOpenApiFlags] ERROR: Failed to retrieve field/flag value; %s", field_flag)
			continue
		}
		if flag_value == "" {
			continue
		}
		field_kind := field_type.Type.Kind()
		field_isptr := false
		if field_kind == reflect.Ptr {
			field_isptr = true
			field.Set(reflect.New(field_type.Type.Elem()))
			field_kind = field_type.Type.Elem().Kind()
		}
		switch field_kind {

		case reflect.Bool:
			if flag_value_typed, err := strconv.ParseBool(flag_value); err == nil {
				if field_isptr {
					field.Elem().SetBool(flag_value_typed)
				} else {
					field.SetBool(flag_value_typed)
				}
			} else {
				log.Printf("[cmd.ParseOpenApiFlags] ERROR: Failed to parse field/flag value; %s (%s)", field_flag, field_kind)
			}

		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if flag_value_typed, err := strconv.ParseInt(flag_value, 10, 64); err == nil {
				if field_isptr {
					field.Elem().SetInt(flag_value_typed)
				} else {
					field.SetInt(flag_value_typed)
				}
			} else {
				log.Printf("[cmd.ParseOpenApiFlags] ERROR: Failed to parse field/flag value; %s (%s)", field_flag, field_kind)
			}

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if flag_value_typed, err := strconv.ParseUint(flag_value, 10, 64); err == nil {
				if field_isptr {
					field.Elem().SetUint(flag_value_typed)
				} else {
					field.SetUint(flag_value_typed)
				}
			} else {
				log.Printf("[cmd.ParseOpenApiFlags] ERROR: Failed to parse field/flag value; %s (%s)", field_flag, field_kind)
			}

		case reflect.Float32, reflect.Float64:
			if flag_value_typed, err := strconv.ParseFloat(flag_value, 64); err == nil {
				if field_isptr {
					field.Elem().SetFloat(flag_value_typed)
				} else {
					field.SetFloat(flag_value_typed)
				}
			} else {
				log.Printf("[cmd.ParseOpenApiFlags] ERROR: Failed to parse field/flag value; %s (%s)", field_flag, field_kind)
			}

		case reflect.String:
			if field_isptr {
				field.Elem().SetString(flag_value)
			} else {
				field.SetString(flag_value)
			}

		default:
			log.Printf("[cmd.ParseOpenApiFlags] ERROR: Unsupported field type; %s (%s)", field_flag, field_kind)

		}
	}
}
