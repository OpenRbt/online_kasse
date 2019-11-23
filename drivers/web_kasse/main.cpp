// Server side C program to demonstrate HTTP Server programming
#include "server.h"
#include "kasse.h"
#include "string.h"
#include "stdlib.h"

#define API_PORT_ 8088

std::string ServerCallback(void * kassePtr, HttpParameters * parameters) {
    printf("request recieved\n");
    Kasse * kasse = (Kasse *) kassePtr;
    int sum = atoi(parameters->GetValue("sum").c_str());    
    int card = atoi(parameters->GetValue("card").c_str());
    int post = atoi(parameters->GetValue("post").c_str());
    return kasse->PrintReceipt(sum, card, post);
}

int main(int argc, char const *argv[]) {
    printf("kasse server is starting...");
    SimpleServer * server = new SimpleServer(API_PORT_);
    Kasse * kasse = new Kasse();
    server->SetCallback(ServerCallback, kasse);
    server->Run();
    
    delete server;
    delete kasse;
    printf("kasse server stopped ...");
    return 0;
}
