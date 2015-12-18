#include <stdio.h>
#include <string.h>
#include <errno.h>

#define checkError(ret) do{if(-1==ret){printf("[%d]err:%s\n", __LINE__, strerror(errno));return -1;}}while(0)
#define checkNull(p)	do{if(NULL==p){printf("[%d]pointer is null\n", __LINE__);return -1;}}while(0)


int revert_string(char *str){
	checkNull(str);
	int len = strlen(str);
	printf("str:%s,len:%d\n", str, len);

	int i = 0;
	char temp = 0;
	int n = len/2;	
	for(i = 0; i < n; i++){
		temp = str[len-1-i];
		str[len-1-i] = str[i];
		str[i] = temp;
	}
	return 0;
}

int main(int argc, char const *argv[])
{
	printf("revert project\n");
	char input[256] = {0};
	fgets(input, sizeof(input), stdin);
	printf("you input is:%s\n", input);
	checkError(revert_string(input));
	printf("revert str is:%s\n", input);

	return 0;
}
