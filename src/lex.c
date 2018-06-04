#include <stdio.h>
#include <stdlib.h> 
#include "recycle.h"
#include <string.h>
#include <ctype.h>

tk* parse_file(char* filename){
	FILE* source_fp;

	source_fp  = fopen(filename, "r");

	//determine the length of the file
	fseek(source_fp, 0, SEEK_END);
	size_t len = (size_t)ftell(source_fp);
	fseek(source_fp, 0, SEEK_SET);

	//load the file into memory
	char *source = (char*)malloc(sizeof(char)*len);
	
	if(!source){
		printf("Malloc error\n");
		exit(1);
	}


	fread(source, 1, len, source_fp);
	fclose(source_fp);

	printf("%s\n", source);

	//null terminate
	*(source + len) = '\0';

	//start tokenizing the source

	tk* token_list = (tk*)malloc(sizeof(tk));
	int num_tokens = 0;

	//DEBUG ONLY
	int repetitions = 0;
	
	while(*source != '\0'){
		char cur_string[100] = "";
		//this is an ugly hack and I don't like it
		int stuck = 1;
		int length = 0;
		tk new_token;
		new_token.type = -1;
		//process the individual tokens
		while(*source != ' ' && *source != '\0' && new_token.type < 1 && *source != '\n'&& *source != '\t'){
			//determine if there are a bunch of characters together with no whitespace
			if(length){
				int flag = 0;
				for(int i = 0; i < NUM_UNARIES; i++){
					if(*source == unaries[i]){
						flag = 1;
					}
					
				}
				//break if the next character is another token
				if(flag){
					break;
				}
				
			}
			cur_string[length] = *source;
			length++;

			//check for the single length tokens
			if(length < 2){
				switch(cur_string[0]){
					case '{':
						new_token.type = OPEN_BRACKET;
						break;
					case '}':
						new_token.type = CLOSE_BRACKET;
						break;
					case '(' :
						new_token.type = OPEN_PAREN;
						break;
					case ')' :
						new_token.type = CLOSE_PAREN;
						break;
					case ';' : 
						new_token.type = SEMI_COL;
						break;
				}
				
			}
			stuck = 0;
			repetitions++;
			source += 1;

		}
		
		//advance the pointer if necessary
		if(stuck){
			source += 1;
		}
		
		//catch the whitespace
		if(!length){
			continue;
		}
		
		//add the null terminator
		cur_string[length] = '\0';

		if(cur_string[0] > '0' && cur_string[0] < '9'){
				new_token.type = INT_LIT;
				sscanf(cur_string, "%d", &new_token.value);
			}
		
		if(length > 1 && new_token.type < 0){
			
			//check if the token is a keyword
			if(cur_string[0] < '0' || cur_string[0] > '9'){
				for(int i = 0; i < NUM_KEYWORDS; i++){
					if(!strcmp(cur_string, keywords[i])){
						new_token.type = KEYWORD;
						new_token.value = i;
					}
				}
			}
			//else the token is an indentifier (variable)
			if(new_token.type < 0){
				new_token.type = IDENTIFIER;
				
			}
			
		}
		new_token.name = (char*)malloc(sizeof(char) * length);
		strcpy(new_token.name, cur_string);
		//add the token to the list
		num_tokens++;
		token_list = (tk*)realloc(token_list, num_tokens * sizeof(tk));
		*(token_list + ((num_tokens - 1)* sizeof(tk))) = new_token;

		printf("TOKEN: %d\n",num_tokens);
		printf("TYPE: %d\n", new_token.type);
		printf("VALUE: %d\n", new_token.value);
		printf("NAME: %s\n\n", new_token.name);

		
		
		
		
		
	}
	number_of_tokens = num_tokens;

	return token_list;

}

void printTokenList (tk *list){
	for(int i = 0; i < number_of_tokens; i++){
		printf("TOKEN: %d\n",i);
		printf("TYPE: %d\n", list->type);

		if(list->type < KEYWORD){
			continue;
		}

		else if(list->type == INT_LIT){
			printf("VALUE: %d\n", list->value);
		}
		else{
			printf("NAME: %s\n", list->name);
		}
		list += sizeof(tk);
	}
	printf("\n TOTAL TOKENS: %d\n", number_of_tokens);
	
}

int main(int argc, char *argv[]){
	tk* list = parse_file(argv[1]);
	//printTokenList(list);
	return 0;
}	

	