// num.cpp
#include "include/nummer.hpp"
#include "include/num.h"
/* #include "/cheonsy/home/lattigo/gpu-core/BindingTest/goFile/_obj/_cgo_export.h" */


Num NumInit() {
  cxxNum * ret = new cxxNum(1);
  return (void*)ret;
}

void NumFree(Num n) {
  cxxNum * num = (cxxNum*)n;
  delete num;
}

void NumIncrement(Num n) {
  cxxNum * num = (cxxNum*)n;
  num->Increment();
}

int NumGetValue(Num n) {
  cxxNum * num = (cxxNum*)n;
  return num->GetValue();
}

