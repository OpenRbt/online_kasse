package fptr10

/*
#include "libfptr10_go.h"
*/
import "C"

import (
    "path"
    "unsafe"

    "golang.org/x/sys/windows"
    "golang.org/x/sys/windows/registry"
)

func mustGetProcAddress(lib windows.Handle, name string) unsafe.Pointer {
	addr, err := windows.GetProcAddress(lib, name)
	if err != nil {
		panic(err)
	}

	return unsafe.Pointer(addr)
}

func loadLibrary() functionPointers {
    lib, err := windows.LoadLibrary(getLibPath())
    if err != nil {
        panic("Can't load library \"" + getLibPath() + "\" - " + err.Error())
    }
    return functionPointers{
        C.libfptr_create_func(mustGetProcAddress(lib, "libfptr_create")),
        C.libfptr_destroy_func(mustGetProcAddress(lib, "libfptr_destroy")),

        C.libfptr_get_version_string_func(mustGetProcAddress(lib, "libfptr_get_version_string")),

        C.libfptr_set_settings_func(mustGetProcAddress(lib, "libfptr_set_settings")),
        C.libfptr_get_settings_func(mustGetProcAddress(lib, "libfptr_get_settings")),
        C.libfptr_set_single_setting_func(mustGetProcAddress(lib, "libfptr_set_single_setting")),
        C.libfptr_get_single_setting_func(mustGetProcAddress(lib, "libfptr_get_single_setting")),

        C.libfptr_is_opened_func(mustGetProcAddress(lib, "libfptr_is_opened")),

        C.libfptr_error_code_func(mustGetProcAddress(lib, "libfptr_error_code")),
        C.libfptr_error_description_func(mustGetProcAddress(lib, "libfptr_error_description")),
        C.libfptr_reset_error_func(mustGetProcAddress(lib, "libfptr_reset_error")),

        C.libfptr_set_param_bool_func(mustGetProcAddress(lib, "libfptr_set_param_bool")),
        C.libfptr_set_param_int_func(mustGetProcAddress(lib, "libfptr_set_param_int")),
        C.libfptr_set_param_double_func(mustGetProcAddress(lib, "libfptr_set_param_double")),
        C.libfptr_set_param_str_func(mustGetProcAddress(lib, "libfptr_set_param_str")),
        C.libfptr_set_param_bytearray_func(mustGetProcAddress(lib, "libfptr_set_param_bytearray")),
        C.libfptr_set_param_datetime_func(mustGetProcAddress(lib, "libfptr_set_param_datetime")),

        C.libfptr_set_param_bool_func(mustGetProcAddress(lib, "libfptr_set_non_printable_param_bool")),
        C.libfptr_set_param_int_func(mustGetProcAddress(lib, "libfptr_set_non_printable_param_int")),
        C.libfptr_set_param_double_func(mustGetProcAddress(lib, "libfptr_set_non_printable_param_double")),
        C.libfptr_set_param_str_func(mustGetProcAddress(lib, "libfptr_set_non_printable_param_str")),
        C.libfptr_set_param_bytearray_func(mustGetProcAddress(lib, "libfptr_set_non_printable_param_bytearray")),
        C.libfptr_set_param_datetime_func(mustGetProcAddress(lib, "libfptr_set_non_printable_param_datetime")),

        C.libfptr_get_param_bool_func(mustGetProcAddress(lib, "libfptr_get_param_bool")),
        C.libfptr_get_param_int_func(mustGetProcAddress(lib, "libfptr_get_param_int")),
        C.libfptr_get_param_double_func(mustGetProcAddress(lib, "libfptr_get_param_double")),
        C.libfptr_get_param_str_func(mustGetProcAddress(lib, "libfptr_get_param_str")),
        C.libfptr_get_param_bytearray_func(mustGetProcAddress(lib, "libfptr_get_param_bytearray")),
        C.libfptr_get_param_datetime_func(mustGetProcAddress(lib, "libfptr_get_param_datetime")),

        C.libfptr_log_write_func(mustGetProcAddress(lib, "libfptr_log_write")),

        C.libfptr_show_properties_func(mustGetProcAddress(lib, "libfptr_show_properties")),

        C.libfptr_set_param_bool_func(mustGetProcAddress(lib, "libfptr_set_user_param_bool")),
        C.libfptr_set_param_int_func(mustGetProcAddress(lib, "libfptr_set_user_param_int")),
        C.libfptr_set_param_double_func(mustGetProcAddress(lib, "libfptr_set_user_param_double")),
        C.libfptr_set_param_str_func(mustGetProcAddress(lib, "libfptr_set_user_param_str")),
        C.libfptr_set_param_bytearray_func(mustGetProcAddress(lib, "libfptr_set_user_param_bytearray")),
        C.libfptr_set_param_datetime_func(mustGetProcAddress(lib, "libfptr_set_user_param_datetime")),

        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_apply_single_settings")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_open")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_close")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_reset_params")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_run_command")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_beep")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_open_drawer")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_cut")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_device_poweroff")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_device_reboot")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_open_shift")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_reset_summary")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_init_device")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_query_data")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_cash_income")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_cash_outcome")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_open_receipt")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_cancel_receipt")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_close_receipt")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_check_document_closed")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_receipt_total")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_receipt_tax")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_registration")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_payment")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_report")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_print_text")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_print_cliche")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_begin_nonfiscal_document")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_end_nonfiscal_document")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_print_barcode")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_print_picture")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_print_picture_by_number")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_upload_picture_from_file")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_clear_pictures")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_write_device_setting_raw")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_read_device_setting_raw")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_commit_settings")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_init_settings")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_reset_settings")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_write_date_time")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_write_license")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_fn_operation")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_fn_query_data")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_fn_write_attributes")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_external_device_power_on")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_external_device_power_off")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_external_device_write_data")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_external_device_read_data")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_operator_login")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_process_json")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_read_device_setting")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_write_device_setting")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_begin_read_records")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_read_next_record")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_end_read_records")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_user_memory_operation")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_continue_print")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_init_mgm")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_util_form_tlv")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_util_form_nomenclature")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_util_mapping")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_read_model_flags")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_line_feed")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_flash_firmware")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_soft_lock_init")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_soft_lock_query_session_code")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_soft_lock_validate")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_util_calc_tax")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_download_picture")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_bluetooth_remove_paired_devices")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_util_tag_info")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_util_container_versions")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_activate_licenses")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_remove_licenses")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_enter_keys")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_validate_keys")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_enter_serial_number")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_get_serial_number_request")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_upload_pixel_buffer")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_download_pixel_buffer")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_print_pixel_buffer")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_util_convert_tag_value")),
        C.libfptr_simple_call_func(mustGetProcAddress(lib, "libfptr_parse_marking_code")),
    }
}

func getLibPath() string {
	var libPath = ""
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\ATOL\\Drivers\\10.0\\KKT", registry.QUERY_VALUE)
	if err == nil {
		installDir, _, err := key.GetStringValue("INSTALL_DIR")
		if err == nil {
			libPath = path.Join(installDir, "bin")
		}
	}
	libPath = path.Join(libPath, "fptr10.dll")
	key.Close()
	return libPath
}

