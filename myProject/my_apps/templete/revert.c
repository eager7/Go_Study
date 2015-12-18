#include <stdio.h>

#define checkError(ret) do{if(-1==ret){printf("[%d]err:%s\n", __LINE__, strerror(errno));return -1;}}while(0)

int main(int argc, char const *argv[])
{
	printf("revert project\n");


	return 0;
}

