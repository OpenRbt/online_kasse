#include "kasse.h"
#include <string>

int main(int argc, char **argv) {
    fprintf(stderr, "Driver program started...\n");

    if (argc != 4) {
	fprintf(stderr, "Ivalid command line arguments. Should be: <sum><iscard><post>\n");
	return 2;
    }

    char* sumStr = argv[1];
    char* isCardStr = argv[2];
    char* postStr = argv[3];

    int sum = atoi(sumStr);
    int isCard = atoi(isCardStr);
    int post = atoi(postStr);

    fprintf(stderr, "Driver got data:\n");
    fprintf(stderr, "Sum: %d\n", sum);
    fprintf(stderr, "IsCard: %d\n", isCard);
    fprintf(stderr, "Post: %d\n", post);

    Kasse device;
    std::string res = device.PrintReceipt(sum,isCard,post);
    
    if (res != "PRINTED") {
	return 1;
    }

    return 0;
}
