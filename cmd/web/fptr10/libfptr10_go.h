/**
 * AUTO GENERATED FILE
 */

#ifndef LIBFPTR10_GO_H_
#define LIBFPTR10_GO_H_

#include <wchar.h>

#include "libfptr10.h"

typedef int (*libfptr_create_func)(libfptr_handle *fptr);
typedef int (*libfptr_create_with_id_func)(libfptr_handle *fptr, const wchar_t *id);
typedef void (*libfptr_destroy_func)(libfptr_handle *fptr);
typedef const char* (*libfptr_get_version_string_func)(libfptr_handle fptr);
typedef int (*libfptr_set_settings_func)(libfptr_handle fptr, const wchar_t *settings);
typedef int (*libfptr_get_settings_func)(libfptr_handle fptr, const wchar_t *value, int size);
typedef void (*libfptr_set_single_setting_func)(libfptr_handle fptr, const wchar_t *key, const wchar_t *value);
typedef int (*libfptr_get_single_setting_func)(libfptr_handle fptr, const wchar_t *key, wchar_t *value, int size);
typedef int (*libfptr_is_opened_func)(libfptr_handle fptr);
typedef int (*libfptr_error_code_func)(libfptr_handle fptr);
typedef int (*libfptr_error_description_func)(libfptr_handle fptr, wchar_t *value, int size);
typedef int (*libfptr_error_recommendation_func)(libfptr_handle fptr, wchar_t *value, int size);
typedef int (*libfptr_reset_error_func)(libfptr_handle fptr);
typedef void (*libfptr_set_param_bool_func)(libfptr_handle fptr, int param_id, int value);
typedef void (*libfptr_set_param_int_func)(libfptr_handle fptr, int param_id, uint value);
typedef void (*libfptr_set_param_double_func)(libfptr_handle fptr, int param_id, double value);
typedef void (*libfptr_set_param_str_func)(libfptr_handle fptr, int param_id, const wchar_t *value);
typedef void (*libfptr_set_param_bytearray_func)(libfptr_handle fptr, int param_id, const unsigned char *value, int size);
typedef void (*libfptr_set_param_datetime_func)(libfptr_handle fptr, int param_id, int year, int month, int day, int hour, int minute, int second);
typedef int (*libfptr_get_param_bool_func)(libfptr_handle fptr, int param_id);
typedef uint (*libfptr_get_param_int_func)(libfptr_handle fptr, int param_id);
typedef double (*libfptr_get_param_double_func)(libfptr_handle fptr, int param_id);
typedef int (*libfptr_get_param_str_func)(libfptr_handle fptr, int param_id, wchar_t *value, int size);
typedef int (*libfptr_get_param_bytearray_func)(libfptr_handle fptr, int param_id, unsigned char *value, int size);
typedef void (*libfptr_get_param_datetime_func)(libfptr_handle fptr, int param_id, int *year, int *month, int *day, int *hour, int *minute, int *second);
typedef int (*libfptr_is_param_available_func)(libfptr_handle fptr, int param_id);
typedef int (*libfptr_log_write_func)(libfptr_handle fptr, const wchar_t *tag, int level, const wchar_t *message);
typedef int (*libfptr_change_label_func)(libfptr_handle fptr, const wchar_t *label);
typedef int (*libfptr_show_properties_func)(libfptr_handle fptr, uint parent_type, void *parent);

typedef int (*libfptr_simple_call_func)(libfptr_handle fptr);

#endif
