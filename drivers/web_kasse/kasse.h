#ifndef _DIAE_KASSE_
#define _DIAE_KASSE_
#include <string>
#include "libfptr10.h"

class Kasse {
    public:
    // PrintReceipt must return "PRINTED" if no errors, or "FAIL:" if there are errors
    std::string PrintReceipt(int sum, int card, int post) {
		fprintf(stderr, "New receipt accepted by driver\n");

		libfptr_handle fptr;
		int errorCode;

		libfptr_create(&fptr);

		// Step 1: Configure connection to device
		libfptr_set_single_setting(fptr, LIBFPTR_SETTING_PORT, std::to_wstring(LIBFPTR_PORT_USB).c_str());
		libfptr_set_single_setting(fptr, LIBFPTR_SETTING_MODEL, std::to_wstring(LIBFPTR_MODEL_KAZNACHEY_FA).c_str());
		libfptr_apply_single_settings(fptr);

		// Step 2: Connect to device
		libfptr_open(fptr);
		bool isOpened = (libfptr_is_opened(fptr) != 0);
		if (!isOpened) {
			fprintf(stderr, "USB connection to device failed\n");
			libfptr_destroy(&fptr);
			return "FAIL: CONNECTION FAILURE";
		}
		fprintf(stderr, "Connection to device opened\n");
		errorCode = libfptr_error_code(fptr);
		fprintf(stderr, "Error code in connection: %d", errorCode);

		// Stage 3: Register the responsible person and log in
		libfptr_set_param_str(fptr, 1021, L"Канатников А.В.");
		libfptr_set_param_str(fptr, 1203, L"5401199801");
		libfptr_operator_login(fptr);
		errorCode = libfptr_error_code(fptr);
		fprintf(stderr, "Error code in operator login: %d", errorCode);
		/*
		if (errorCode != 0) {
			printf("Operator login failure\n");
			libfptr_close(fptr);
			libfptr_destroy(&fptr);
			return "FAIL: OPERATOR FAILURE";
		}
		*/
		// Stage 4: Check the shift
		libfptr_open_shift(fptr);
		errorCode = libfptr_error_code(fptr);
		fprintf(stderr, "Error code in shift check: %d", errorCode);

		// If shift expired (was more than 24 hours long) - close it and open again
		if (errorCode == 68 || errorCode == 141) {
			fprintf(stderr, "Shift expired - closing and reopening\n");

			libfptr_set_param_str(fptr, LIBFPTR_PARAM_REPORT_TYPE, std::to_wstring(LIBFPTR_RT_CLOSE_SHIFT).c_str());
			libfptr_report(fptr);
			if (errorCode != 0) {
				fprintf(stderr, "Shift close failed\n");
				libfptr_close(fptr);
				libfptr_destroy(&fptr);
				return "FAIL: SHIFT CLOSE FAILURE";
			}

			libfptr_open_shift(fptr);
			if (errorCode != 0) {
				fprintf(stderr, "Shift open failed\n");
				libfptr_close(fptr);
				libfptr_destroy(&fptr);
				return "FAIL: SHIFT OPEN FAILURE";
			}
		}

		// Stage 5: Open receipt
		libfptr_set_param_int(fptr, LIBFPTR_PARAM_RECEIPT_TYPE, LIBFPTR_RT_SELL);
		libfptr_open_receipt(fptr);
		errorCode = libfptr_error_code(fptr);
		fprintf(stderr, "Error code in open receipt: %d", errorCode);

		// Stage 6: Register the service or commodity
		libfptr_set_param_str(fptr, LIBFPTR_PARAM_COMMODITY_NAME, L"АВТОМОЙКА");
		libfptr_set_param_double(fptr, LIBFPTR_PARAM_PRICE, double(sum));
		libfptr_set_param_double(fptr, LIBFPTR_PARAM_QUANTITY, 1.0);
		libfptr_set_param_int(fptr, LIBFPTR_PARAM_TAX_TYPE, LIBFPTR_TAX_NO);
		libfptr_set_param_int(fptr, 1212, 4);
		libfptr_set_param_int(fptr, 1214, 1);
		libfptr_registration(fptr);
		errorCode = libfptr_error_code(fptr);
		fprintf(stderr, "Error code in registration: %d", errorCode);

		// Stage 7: Register total
		libfptr_set_param_double(fptr, LIBFPTR_PARAM_SUM, double(sum));
		libfptr_receipt_total(fptr);

		// Stage 8: Set the payment method
		if (card) {
			libfptr_set_param_int(fptr, LIBFPTR_PARAM_PAYMENT_TYPE, LIBFPTR_PT_ELECTRONICALLY);
		} else {
			libfptr_set_param_int(fptr, LIBFPTR_PARAM_PAYMENT_TYPE, LIBFPTR_PT_CASH);
		}
		libfptr_set_param_double(fptr, LIBFPTR_PARAM_PAYMENT_SUM, double(sum));
		libfptr_payment(fptr);

		// Stage 9: Close the receipt
		libfptr_close_receipt(fptr);
		errorCode = libfptr_error_code(fptr);
		fprintf(stderr, "Error code in close receipt: %d", errorCode);
		while (libfptr_check_document_closed(fptr) < 0) {
			fprintf(stderr, "Attempt to close document...\n");
			continue;
		}

		errorCode = libfptr_error_code(fptr);
		fprintf(stderr, "Error code before check document: %d", errorCode);

		// Stage 10: Check document
		if (libfptr_get_param_bool(fptr, LIBFPTR_PARAM_DOCUMENT_CLOSED) == 0) {
			libfptr_cancel_receipt(fptr);
			fprintf(stderr, "Document was not closed - receipt cancelled\n");
			libfptr_close(fptr);
			libfptr_destroy(&fptr);
			return "FAIL: DOCUMENT CLOSE FAILED";
		}

		// Stage 11: Check the printing process
		if (libfptr_get_param_bool(fptr, LIBFPTR_PARAM_DOCUMENT_PRINTED) == 0) {
			while (libfptr_continue_print(fptr) < 0) {
				fprintf(stderr, "Attempt to print the document\n");
				continue;
			}
		}

		// Stage 12: Get data about last document
		libfptr_set_param_int(fptr, LIBFPTR_PARAM_FN_DATA_TYPE, LIBFPTR_FNDT_LAST_DOCUMENT);
		libfptr_fn_query_data(fptr);
		int documentNumber = libfptr_get_param_int(fptr, LIBFPTR_PARAM_DOCUMENT_NUMBER);
		fprintf(stderr, "Receipt was registered with number %d\n", documentNumber);

		libfptr_close(fptr);
		libfptr_destroy(&fptr);

        fprintf(stderr, "Receipt printed: sum: %d, is_card: %d, post: %d\n", sum, card, post);
        return "PRINTED";
    }
};
#endif
