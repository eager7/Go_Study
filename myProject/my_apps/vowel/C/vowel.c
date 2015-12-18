#include <stdio.h>
#include <string.h>

int count_vowel(char *str, char s){
	int i = 0, num = 0, len = strlen(str);
	for(i = 0; i < len; i++){
		if(str[i] == s){
			num++;
		}
	}
	return num;
}

int main(int argc, char const *argv[])
{
	printf("vowel project\n");
	char inp[256] = {0};
	fgets(inp, sizeof(inp), stdin);
	printf("your input is:%s\n", inp);

	int i = 0;
	for(i = 0; i < strlen(inp); i++){
		printf("num of %c is:%d\n", inp[i], count_vowel(inp, 'a'));
	}
	return 0;
}