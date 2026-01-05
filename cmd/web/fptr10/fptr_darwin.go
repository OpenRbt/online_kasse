/**
 * AUTO GENERATED FILE
 */

package fptr10

import "C"

/*
#cgo LDFLAGS: -ldl
#include <dlfcn.h>
#include "libfptr10_go.h"
*/
import "C"
import (
	"fmt"
	"os"
	"path/filepath"
	"unsafe"
)

func getProcAddress(lib unsafe.Pointer, name string) unsafe.Pointer {
	addr := C.dlsym(lib, C.CString(name))
	if addr == nil {
		return nil
	}

	return unsafe.Pointer(addr)
}

func doLoadLibrary(path string) (unsafe.Pointer, error) {
	var loadError error
	var lib unsafe.Pointer
	for _, libName := range []string{"libfptr10.dylib", "fptr10.framework/fptr10"} {
		loadError = nil

		if path == "" {
			path = libName
		}

		if path != libName {
			fi, err := os.Stat(path)
			if err != nil {
				loadError = fmt.Errorf("can't load library \"%s\" - %s", path, err)
				continue
			}

			if fi.IsDir() {
				path = filepath.Join(path, libName)
			}
		}

		lib = C.dlopen(C.CString(path), C.RTLD_LAZY)
		if lib == nil {
			loadError = fmt.Errorf("can't load library \"%s\" - %s", path, C.GoString(C.dlerror()))
			continue
		}

		break
	}

	return lib, loadError
}

func loadLibrary(path string) (*functionPointers, error) {
	var lib unsafe.Pointer
	var loadError error
	if path == "" {
		var exe string
		if exe, loadError = os.Executable(); loadError == nil {
			currentDirectory := filepath.Dir(exe)
			if lib, loadError = doLoadLibrary(currentDirectory); loadError != nil {
				lib, loadError = doLoadLibrary("")
			}
		}
	} else {
		lib, loadError = doLoadLibrary(path)
	}

	if lib == nil {
		return nil, loadError
	}

	return &functionPointers{
		C.libfptr_create_func(getProcAddress(lib, "libfptr_create")),
		C.libfptr_create_with_id_func(getProcAddress(lib, "libfptr_create_with_id")),
		C.libfptr_destroy_func(getProcAddress(lib, "libfptr_destroy")),

		C.libfptr_get_version_string_func(getProcAddress(lib, "libfptr_get_version_string")),

		C.libfptr_set_settings_func(getProcAddress(lib, "libfptr_set_settings")),
		C.libfptr_get_settings_func(getProcAddress(lib, "libfptr_get_settings")),
		C.libfptr_set_single_setting_func(getProcAddress(lib, "libfptr_set_single_setting")),
		C.libfptr_get_single_setting_func(getProcAddress(lib, "libfptr_get_single_setting")),

		C.libfptr_is_opened_func(getProcAddress(lib, "libfptr_is_opened")),

		C.libfptr_error_code_func(getProcAddress(lib, "libfptr_error_code")),
		C.libfptr_error_description_func(getProcAddress(lib, "libfptr_error_description")),
		C.libfptr_error_recommendation_func(getProcAddress(lib, "libfptr_error_recommendation")),
		C.libfptr_reset_error_func(getProcAddress(lib, "libfptr_reset_error")),

		C.libfptr_set_param_bool_func(getProcAddress(lib, "libfptr_set_param_bool")),
		C.libfptr_set_param_int_func(getProcAddress(lib, "libfptr_set_param_int")),
		C.libfptr_set_param_double_func(getProcAddress(lib, "libfptr_set_param_double")),
		C.libfptr_set_param_str_func(getProcAddress(lib, "libfptr_set_param_str")),
		C.libfptr_set_param_bytearray_func(getProcAddress(lib, "libfptr_set_param_bytearray")),
		C.libfptr_set_param_datetime_func(getProcAddress(lib, "libfptr_set_param_datetime")),

		C.libfptr_set_param_bool_func(getProcAddress(lib, "libfptr_set_non_printable_param_bool")),
		C.libfptr_set_param_int_func(getProcAddress(lib, "libfptr_set_non_printable_param_int")),
		C.libfptr_set_param_double_func(getProcAddress(lib, "libfptr_set_non_printable_param_double")),
		C.libfptr_set_param_str_func(getProcAddress(lib, "libfptr_set_non_printable_param_str")),
		C.libfptr_set_param_bytearray_func(getProcAddress(lib, "libfptr_set_non_printable_param_bytearray")),
		C.libfptr_set_param_datetime_func(getProcAddress(lib, "libfptr_set_non_printable_param_datetime")),

		C.libfptr_get_param_bool_func(getProcAddress(lib, "libfptr_get_param_bool")),
		C.libfptr_get_param_int_func(getProcAddress(lib, "libfptr_get_param_int")),
		C.libfptr_get_param_double_func(getProcAddress(lib, "libfptr_get_param_double")),
		C.libfptr_get_param_str_func(getProcAddress(lib, "libfptr_get_param_str")),
		C.libfptr_get_param_bytearray_func(getProcAddress(lib, "libfptr_get_param_bytearray")),
		C.libfptr_get_param_datetime_func(getProcAddress(lib, "libfptr_get_param_datetime")),

		C.libfptr_is_param_available_func(getProcAddress(lib, "libfptr_is_param_available")),

		C.libfptr_log_write_func(getProcAddress(lib, "libfptr_log_write_ex")),
		C.libfptr_change_label_func(getProcAddress(lib, "libfptr_change_label")),

		C.libfptr_show_properties_func(getProcAddress(lib, "libfptr_show_properties")),

		C.libfptr_set_param_bool_func(getProcAddress(lib, "libfptr_set_user_param_bool")),
		C.libfptr_set_param_int_func(getProcAddress(lib, "libfptr_set_user_param_int")),
		C.libfptr_set_param_double_func(getProcAddress(lib, "libfptr_set_user_param_double")),
		C.libfptr_set_param_str_func(getProcAddress(lib, "libfptr_set_user_param_str")),
		C.libfptr_set_param_bytearray_func(getProcAddress(lib, "libfptr_set_user_param_bytearray")),
		C.libfptr_set_param_datetime_func(getProcAddress(lib, "libfptr_set_user_param_datetime")),

		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_apply_single_settings")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_open")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_close")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_reset_params")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_run_command")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_beep")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_open_drawer")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_cut")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_device_poweroff")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_device_reboot")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_open_shift")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_reset_summary")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_init_device")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_query_data")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_cash_income")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_cash_outcome")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_open_receipt")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_cancel_receipt")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_close_receipt")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_check_document_closed")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_receipt_total")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_receipt_tax")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_registration")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_payment")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_report")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_print_text")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_print_cliche")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_begin_nonfiscal_document")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_end_nonfiscal_document")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_print_barcode")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_print_picture")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_print_picture_by_number")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_upload_picture_from_file")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_clear_pictures")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_write_device_setting_raw")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_read_device_setting_raw")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_commit_settings")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_init_settings")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_reset_settings")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_write_date_time")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_write_license")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_fn_operation")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_fn_query_data")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_fn_write_attributes")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_external_device_power_on")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_external_device_power_off")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_external_device_write_data")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_external_device_read_data")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_operator_login")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_process_json")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_read_device_setting")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_write_device_setting")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_begin_read_records")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_read_next_record")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_end_read_records")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_user_memory_operation")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_continue_print")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_init_mgm")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_util_form_tlv")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_util_form_nomenclature")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_util_mapping")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_read_model_flags")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_line_feed")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_flash_firmware")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_soft_lock_init")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_soft_lock_query_session_code")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_soft_lock_validate")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_util_calc_tax")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_download_picture")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_bluetooth_remove_paired_devices")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_util_tag_info")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_util_container_versions")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_activate_licenses")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_remove_licenses")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_enter_keys")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_validate_keys")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_enter_serial_number")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_get_serial_number_request")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_upload_pixel_buffer")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_download_pixel_buffer")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_print_pixel_buffer")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_util_convert_tag_value")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_parse_marking_code")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_call_script")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_set_header_lines")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_set_footer_lines")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_upload_picture_cliche")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_upload_picture_memory")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_upload_pixel_buffer_cliche")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_upload_pixel_buffer_memory")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_exec_driver_script")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_upload_driver_script")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_exec_driver_script_by_id")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_write_universal_counters_settings")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_read_universal_counters_settings")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_query_universal_counters_state")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_reset_universal_counters")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_cache_universal_counters")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_read_universal_counter_sum")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_read_universal_counter_quantity")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_clear_universal_counters_cache")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_disable_ofd_channel")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_enable_ofd_channel")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_validate_json")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_reflection_call")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_get_remote_server_info")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_begin_marking_code_validation")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_cancel_marking_code_validation")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_get_marking_code_validation_status")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_accept_marking_code")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_decline_marking_code")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_update_fnm_keys")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_write_sales_notice")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_check_marking_code_validations_ready")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_clear_marking_code_validation_result")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_ping_marking_server")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_get_marking_server_status")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_is_driver_locked")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_get_last_document_journal")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_find_document_in_journal")),
		C.libfptr_simple_call_func(getProcAddress(lib, "libfptr_run_fn_command")),
	}, nil
}
