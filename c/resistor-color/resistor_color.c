#include "resistor_color.h"
#include <stdio.h>
#include <stdint.h>

__attribute__((const)) resistor_band_t COLORS_LIST[] = { BLACK, BROWN, RED, ORANGE, YELLOW, GREEN,BLUE, VIOLET, GREY, WHITE };

__attribute__((const)) resistor_band_t* colors(void) {
	return COLORS_LIST;
}

__attribute__((const)) int color_code(int c) {
	return COLORS_LIST[c];
}
