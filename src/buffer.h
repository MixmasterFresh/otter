/*
An implementation of the operations buffer
*/

#ifndef __BUFFER_H
#define __BUFFER_H
#include "operation.h"

typedef struct buffer{
  operation ** operations;
  int client_id;
  int length;
  int start;
}

typedef struct buffer_container{
  buffer ** client_buffers;
  int * client_ids;
  int client_count;
}

operation * find_by_id(buffer * array, int id);

void append_operation(buffer * target, operation * subject);

#endif // __BUFFER_H