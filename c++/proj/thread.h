#include "stdio.h"
#include "iostream"
#include "pthread.h"
#include "vector"

using namespace std;

typedef void * (*func)(void*);

class Thread {
private:
    pthread_t t;
    void *arg;
    func dofun;
public:
    static void *ThreadFunc(void *arg);
    Thread(func dofunc){cout << "con" << endl;dofun = dofunc;}
    int CreateThread();
};  

class ThreadPool {
private:
    int maxNum;
    vector<Thread> freeThread;
    vector<thread> busyThread;
}