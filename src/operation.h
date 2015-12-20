/*
A definition of the operation data type and some helpful methods.
*/

#ifndef __OPERATION_H
#define __OPERATION_H

#define DELETE 0
#define INSERT 1

typedef struct operation {
  int parent_id;
  int client_id;
  int type;
  char *insertion;
  int index;
  int end_index;
}

operation * create_insert_operation(int parent_id, int client_id, int index, char * insertion);
operation * create_delete_operation(int parent_id, int client_id, int index, int end_index);

char * operation_to_string(operation *subject);
operation * string_to_operation(char *string);

void transform(operation *subject, operation *applied);

#endif // __OPERATION_H