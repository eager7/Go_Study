#include <stdio.h>
#include <string.h>
#include <errno.h>

#define checkError(ret) do{if(-1==ret){printf("[%d]err:%s\n", __LINE__, strerror(errno));return -1;}}while(0)
#define checkNull(p)	do{if(NULL==p){printf("[%d]pointer is null\n", __LINE__);return -1;}}while(0)

int consonant_search(char head){
	const char consonant[] = {
		0,'b',0,'d',0,'f','g','h','i','j','k',0,'m','n',0,'p','q','r','s','t',0,'v',0,0,0,'z'
	};
	if(head == consonant[head - 0x61]){
		return 0;
	}
	return 1;
}

int consonant_str(char *str, int len){
	if(strlen(str) > len - 5){
		return -1;//no engouh memory
	}

	if(!consonant_search(str[0])){
		printf("your word is a consonant word\n");
		snprintf(&str[strlen(str)], 5, "-%cay", str[0]);
		memcpy(str, &str[1], strlen(str));
		return 0;
	}
	return 1;
}

int main(int argc, char const *argv[])
{
	printf("consonant C project\n");

	char input[16] = {0};
	fgets(input, sizeof(input), stdin);
	printf("you input is:%s\n", input);
	int n = strlen(input);
	input[n-1] = 0;
	consonant_str(input, sizeof(input));
	printf("you pag-input is:%s\n", input);
	return 0;
}