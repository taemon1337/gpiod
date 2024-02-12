package orangepi

// GPIO aliases to offsets
var (
	GPIO2  = GPIO_TO_OFFSET[2]
	GPIO3  = GPIO_TO_OFFSET[3]
	GPIO4  = GPIO_TO_OFFSET[4]
	GPIO5  = GPIO_TO_OFFSET[5]
	GPIO6  = GPIO_TO_OFFSET[6]
	GPIO7  = GPIO_TO_OFFSET[7]
	GPIO8  = GPIO_TO_OFFSET[8]
	GPIO9  = GPIO_TO_OFFSET[9]
	GPI10  = GPIO_TO_OFFSET[10]
	GPI11  = GPIO_TO_OFFSET[11]
	GPI12  = GPIO_TO_OFFSET[12]
	GPI13  = GPIO_TO_OFFSET[13]
	GPI14  = GPIO_TO_OFFSET[14]
	GPI15  = GPIO_TO_OFFSET[15]
	GPI16  = GPIO_TO_OFFSET[16]
	GPI17  = GPIO_TO_OFFSET[17]
)

/*
 gpio readall
 +------+-----+----------+--------+---+   H616   +---+--------+----------+-----+------+
 | GPIO | wPi |   Name   |  Mode  | V | Physical | V |  Mode  | Name     | wPi | GPIO |
 +------+-----+----------+--------+---+----++----+---+--------+----------+-----+------+
 |      |     |     3.3V |        |   |  1 || 2  |   |        | 5V       |     |      |
 |  229 |   0 |    SDA.3 |    OFF | 0 |  3 || 4  |   |        | 5V       |     |      |
 |  228 |   1 |    SCL.3 |    OFF | 0 |  5 || 6  |   |        | GND      |     |      |
 |   73 |   2 |      PC9 |    OFF | 0 |  7 || 8  | 0 | OFF    | TXD.5    | 3   | 226  |
 |      |     |      GND |        |   |  9 || 10 | 0 | OFF    | RXD.5    | 4   | 227  |
 |   70 |   5 |      PC6 |   ALT5 | 0 | 11 || 12 | 0 | OFF    | PC11     | 6   | 75   |
 |   69 |   7 |      PC5 |   ALT5 | 0 | 13 || 14 |   |        | GND      |     |      |
 |   72 |   8 |      PC8 |    OFF | 0 | 15 || 16 | 0 | OFF    | PC15     | 9   | 79   |
 |      |     |     3.3V |        |   | 17 || 18 | 0 | OFF    | PC14     | 10  | 78   |
 |  231 |  11 |   MOSI.1 |    OFF | 0 | 19 || 20 |   |        | GND      |     |      |
 |  232 |  12 |   MISO.1 |    OFF | 0 | 21 || 22 | 0 | OFF    | PC7      | 13  | 71   |
 |  230 |  14 |   SCLK.1 |    OFF | 0 | 23 || 24 | 0 | OFF    | CE.1     | 15  | 233  |
 |      |     |      GND |        |   | 25 || 26 | 0 | OFF    | PC10     | 16  | 74   |
 |   65 |  17 |      PC1 |    OFF | 0 | 27 || 28 |   |        |          |     |      |
 |  272 |  18 |     PI16 |   ALT2 | 0 | 29 || 30 |   |        |          |     |      |
 |  262 |  19 |      PI6 |    OFF | 0 | 31 || 32 |   |        |          |     |      |
 |  234 |  20 |     PH10 |   ALT3 | 0 | 33 || 34 |   |        |          |     |      |
 +------+-----+----------+--------+---+----++----+---+--------+----------+-----+------+
 | GPIO | wPi |   Name   |  Mode  | V | Physical | V |  Mode  | Name     | wPi | GPIO |
 +------+-----+----------+--------+---+   H616   +---+--------+----------+-----+------+
*/
var GPIO_TO_OFFSET = map[int]int{
	2:  73,
	3:  226,
	4:  227,
	5:  70,
	6:  75,
	7:  69,
	8:  72,
	9:  79,
	10:  78,
	11:  231,
	12:  232,
	13:  71,
	14:  230,
	15:  233,
	16:  74,
	17:  65,
}
