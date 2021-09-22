#include "isogram.h"
#include <ctype.h>

int is_isogram(char *w) {
	char visit[26] = { 0 };
	for (int i = 0; w[i] != '\0'; ++i) {
		unsigned char c = tolower(w[i]);
		if (c < 'a' || c > 'z') continue;
		visit[c - 'a'] += 1;
		if (visit[c - 'a'] > 1) return 0;
	}
	return 1;
}
