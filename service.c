#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <arpa/inet.h>
#include <sys/socket.h>

#define PORT 8081
#define BUFFER_SIZE 1024

int main() {
    int server_fd, new_socket;
    struct sockaddr_in address;
    int addrlen = sizeof(address);
    char buffer[BUFFER_SIZE] = {0};
    
    // 1. Create Socket
    if ((server_fd = socket(AF_INET, SOCK_STREAM, 0)) < 0) {
        perror("[C] Socket creation failed");
        exit(EXIT_FAILURE);
    }

    // 2. Bind to Port
    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(PORT);

    if (bind(server_fd, (struct sockaddr *)&address, sizeof(address)) < 0) {
        perror("[C] Bind failed");
        exit(EXIT_FAILURE);
    }

    // 3. Listen
    if (listen(server_fd, 3) < 0) {
        perror("[C] Listen failed");
        exit(EXIT_FAILURE);
    }

    printf("[C] Server listening on port %d\n", PORT);

    // 4. Simple Loop to accept requests
    while(1) {
        if ((new_socket = accept(server_fd, (struct sockaddr *)&address, (socklen_t*)&addrlen)) < 0) {
            perror("[C] Accept failed");
            continue;
        }

        read(new_socket, buffer, BUFFER_SIZE);
        
        // Very basic check if it is a GET /getData request
        // (In production, use a real HTTP parser)
        if (strstr(buffer, "GET /getData") != NULL) {
            printf("[C] Handling /getData request\n");
            char *response = "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n{\"message\": \"Hello from C WASM!\"}";
            write(new_socket, response, strlen(response));
        } else {
            char *response = "HTTP/1.1 404 Not Found\r\n\r\n";
            write(new_socket, response, strlen(response));
        }

        close(new_socket);
    }
    return 0;
}