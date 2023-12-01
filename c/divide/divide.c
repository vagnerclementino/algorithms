#include "divide.h"

int divide(int dividend, int divisor) {
    int count = 0;
    int sum = divisor;
    while (sum <= dividend) {
        sum+=divisor;
        count++;
    }
    return count;
}
