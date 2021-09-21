#include "armstrong_numbers.h"
#include <stdio.h>
#include <stdint.h>
#include <math.h>

__attribute__((const)) bool is_armstrong_number(int c) {
	int temp = c;
	int num_of_digits = 1;
	int sum = 0;
	//computing the number of digits in c
	num_of_digits = log10(c) + 1;
	//computing the sum of digits power the number of digits
	temp = c;
	while (temp > 9) {
		sum += power(temp % 10, num_of_digits);
		temp /= 10;
	}
	sum += power(temp % 10, num_of_digits);
	return c == sum;
}

//returns n power p
__attribute__((const)) int power(int n, int p) {
	int sum = n;
	for (int i = 1; i < p; i++) {
		sum *= n;
	}
	return sum;
}

