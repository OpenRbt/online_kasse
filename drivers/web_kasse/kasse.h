#ifndef _DIAE_KASSE_
#define _DIAE_KASSE_
#include <string>
class Kasse {
    public:
    // PrintReceipt must return "PRINTED" if no errors, or "FAIL:" if there are errors 
    std::string PrintReceipt(int sum, int card, int post) {
        printf("receipt printed: sum: %d, is_card: %d, post: %d\n", sum, card, post);
        return "PRINTED";
    }
};
#endif
