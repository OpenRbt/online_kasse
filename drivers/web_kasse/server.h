#ifndef __SIMPLE_SERVER__
#define __SIMPLE_SERVER__
#include <map>
#include <string>
#include <stdio.h>
#include <sys/socket.h>
#include <unistd.h>
#include <stdlib.h>
#include <netinet/in.h>
#include <string.h>

class HttpParameters {
private:
    std::map<std::string, std::string> values;
public:
    void SetValue(std::string key, std::string value) {
        values[key] = value;
    }
    std::string GetValue(std::string key) {
        return values[key];
    }
    void Parse(char * request) {
        printf("parse called...\n");
    }
};

/*
 char **split( char **result, char *working, const char *src, const char *delim)
    {
        int i;

        strcpy(working, src); // working will get chppped up instead of src 
        char *p=strtok(working, delim);
        for(i=0; p!=NULL && i < (MX_SPLIT -1); i++, p=strtok(NULL, delim) )
        {
            result[i]=p;
            result[i+1]=NULL;  // mark the end of result array
        }
        return result;
    }

    void foo(const char *somestring)
    {
       int i=0;
       char *result[MX_SPLIT]={NULL};
       char working[256]={0x0}; // assume somestring is never bigger than 256 - a weak assumption
       char mydelim[]="!@#$%^&*()_-";
       split(result, working, somestring, mydelim);
       while(result[i]!=NULL)
          printf("token # %d=%s\n", i, result[i]);
    }
*/
    
class SimpleServer {
    int port;
    int toBeStopped;
    void * callbackObject;
    pthread_t worker;
    std::string (*callback)(void*, HttpParameters*);
public:
    SimpleServer(int _port) {
        port = _port;
        callback = 0;
        toBeStopped = 0;
        callbackObject = 0;
    }
    
    void SetCallback(std::string (*_callback)(void*, HttpParameters*), void * _callbackObject){
        callback = _callback;
        callbackObject = _callbackObject;
    }
    
    static void * listeningService(void * args) {
        SimpleServer * _this = (SimpleServer *) args;
        int server_fd, new_socket; long valread;
        struct sockaddr_in address;
        int addrlen = sizeof(address);

        // Only this line has been changed. Everything is same.
        const char *answerTemplate = "HTTP/1.1 200 OK\nContent-Type: text/plain\nContent-Length: %d\n\n";

        // Creating socket file descriptor
        if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) == 0) {
            perror("In socket");
            exit(EXIT_FAILURE);
        }

        address.sin_family = AF_INET;
        address.sin_addr.s_addr = INADDR_ANY;
        address.sin_port = htons(_this->port);

        memset(address.sin_zero, '\0', sizeof address.sin_zero);

        if (bind(server_fd, (struct sockaddr *)&address, sizeof(address))<0) {
            perror("In bind");
            exit(EXIT_FAILURE);
        }
        
        if (listen(server_fd, 10) < 0) {
            perror("In listen");
            exit(EXIT_FAILURE);
        }
        
        while(!_this->toBeStopped) {
            printf("\n+++++++ Waiting for new connection ++++++++\n\n");
            if ((new_socket = accept(server_fd, (struct sockaddr *)&address, (socklen_t*)&addrlen))<0)
            {
                perror("In accept");
                exit(EXIT_FAILURE);
            }

            char buffer[30000] = {0};
            valread = read( new_socket , buffer, 30000-1);
            printf("%s\n",buffer );
            std::string answer = "";
            HttpParameters * parameters = new HttpParameters();
            
            parameters->Parse(buffer);
            
            if (_this->callback!=0) {
                answer = _this->callback(_this->callbackObject, parameters);
            } else {
                answer = "Please connect a callback";
            }
            delete parameters;
            char answerStartBuf[1024];
            snprintf(answerStartBuf, 1023, answerTemplate, answer.length());
            std::string answerStart = answerStartBuf;
            answer = answerStart + answer;
            
            //fprintf(new_socket, answerTemplate, answer.length(), answer.c_str() );
            write(new_socket , answer.c_str() , answer.length());
            printf("answered...\n");
            close(new_socket);
        }     
        return 0;   
    }
    void Run() {
        pthread_create(&worker, 0, listeningService, this);
        void * res;
        pthread_join(worker, &res);
    }
};

#endif
