#include <stdio.h>
#include <stdlib.h>
#include "divide.h"

const int PROG_ARGS = 3; 

int main(int argc, char *argv[]){

    if (argc != PROG_ARGS) {
        printf("Usage: ./divid <num1> <num2>\n");
        return 1;
    }

    int dividend = atoi(argv[1]);
    int divisor = atoi(argv[2]);
    printf("%d\n", divide(dividend, divisor));
}
