#include "thread.h"

int Thread::CreateThread() {
    int ret = pthread_create(&t,NULL,dofun,arg);
    if (ret!=0)
    {
        cout << "create thread error!" << endl;
    }
    else 
    {   
        cout << "success!" << endl; 
    }
    return ret; 
}

void *Thread::ThreadFunc(void *arg) 
{
    Thread *th = (Thread*)arg;
    cout << "in thread" << endl;
}