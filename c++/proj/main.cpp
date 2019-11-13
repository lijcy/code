#include "iostream"
#include "thread.h"
using namespace std;

void * print(void *arg) 
{
    cout << "in asdfafsafasfafasd" <<endl;
}

int main () 
{
    cout << "hello world" << endl;

    Thread *th = new Thread(print);
    th->CreateThread();

    return 0;

    
}