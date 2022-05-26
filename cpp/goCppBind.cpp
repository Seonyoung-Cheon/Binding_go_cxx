#include "goCppBind.h"
#include "/home/cheonsy/lattigo/gpu-core/_obj/_cgo_export.h"

#include <stdlib.h>
#include <iostream>

#ifdef __cplusplus
extern "C"{
#endif

  STUDENT* createRecord(char *name, int entYear){
    STUDENT *p = (STUDENT *)malloc(sizeof(STUDENT));
    p->name = name;
    p->entYear = entYear;
    return p;
  }

  void printStudentCpp(STUDENT *inStudent){
    std::cout << "[C] STUDENT > name : " << inStudent->name << " entYear : " << inStudent->entYear << std::endl;
  }

  void increaseYear(STUDENT* inStudent){
    std::cout << "[C] A year passed" << std::endl;
    inStudent ->entYear++;
  }

  char* callGoFunction(){
    /* char* fromGo=GoFunction1(1, "Hello"); */
    char* fromGo = "Hello";
    std::cout << "[C] From Go Value" << fromGo << std::endl;
    return fromGo;
  }

#ifdef __cplusplus
}
#endif

