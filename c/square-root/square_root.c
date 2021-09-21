#include "square_root.h"
#include <stdio.h>
#include <stdint.h>

int square_root(int n) {
	int d = 1;
	while (d <= n) {
		if (d * d == n) {
			return d;
		}
		d = d + 1;
	}
	return -1;
}


