#ifndef RESISTOR_COLOR_H
#define RESISTOR_COLOR_H

__attribute__((const)) typedef enum {
	BLACK, BROWN, RED, ORANGE, YELLOW, GREEN, BLUE, VIOLET, GREY, WHITE
} resistor_band_t;

__attribute__((const)) resistor_band_t* colors(void);
__attribute__((const)) int color_code(int c);

#endif
