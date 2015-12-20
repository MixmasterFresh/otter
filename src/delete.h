/*
An Implementation of the delete method and some helper methods
*/

#ifndef __DELETE_H
#define __DELETE_H

struct operation;

void delete(int document_id, int start_of_range, int end_of_range, int parent_id, int client_id);

char * apply_delete(operation * subject, char * string);

#endif // __DELETE_H