/*
An implementation of the operations buffer
*/

#ifndef __BUFFER_H
#define __BUFFER_H
#include "operation.h"

typedef struct buffer{
  operation ** array;
  int length;
  int start;
}

int find_by_id(buffer * array, int id);

#endif // __BUFFER_H