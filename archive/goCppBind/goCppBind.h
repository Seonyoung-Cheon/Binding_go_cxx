#ifndef GO_CPP_BIND_H
#define GO_CPP_BIND_H

#ifdef __cplusplus
extern "C"{
#endif

  typedef struct _STUDENT{
    char *name;
    int entYear;
  } STUDENT;

  STUDENT* createRecord(char *name, int entYear);
  void printStudentCpp(STUDENT* inStudent);
  void increaseYear(STUDENT* inStudent);
  char* callGoFunction();

#ifdef __cplusplus
}

#endif
#endif //GO_CPP_BIND_H
