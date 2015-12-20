#include "operation.h"

operation * create_insert_operation(int parent_id, int client_id, int index, char * insertion){
  //TODO: implement method
}

operation * create_delete_operation(int parent_id, int client_id, int index, int end_index){
  //TODO: implement method
}

char * operation_to_string(operation * subject){
  //TODO: implement method
}

operation * string_to_operation(char * string){
  //TODO: implement method
}

void transform(operation * subject, operation * applied){
  //TODO: implement method
}