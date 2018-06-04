#define OPEN_BRACKET 0
#define CLOSE_BRACKET 1
#define OPEN_PAREN 2
#define CLOSE_PAREN 3
#define SEMI_COL 4
#define KEYWORD 5
#define INT_LIT 6
#define IDENTIFIER 7

#define NUM_KEYWORDS 2
#define NUM_UNARIES 7

int number_of_tokens;

typedef struct token {
	int type;
	int value;
	char *name;
}tk;

char keywords[100][16]= {
	"return",
	"int"
};

char unaries[NUM_UNARIES] = "{}[]();";

