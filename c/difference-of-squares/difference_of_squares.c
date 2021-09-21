#include "difference_of_squares.h"

__attribute__((const)) unsigned int sum_of_squares(unsigned int n) {
	return (n * (n + 1) * ((n * 2) + 1)) / 6;
}

__attribute__((const)) unsigned int square_of_sum(unsigned int n) {
	unsigned int sum = ((1 + n) * n) / 2;
	return sum * sum;
}

__attribute__((const)) unsigned int difference_of_squares(unsigned int n) {
	return square_of_sum(n) - sum_of_squares(n);
}
