#include<unistd.h>

void  main (){

    for (char i= 'a'; i<='z'; i++){
        write(1 , &i , 1);
    }

}