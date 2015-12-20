/*
An implementation of the insert method and all relevant helper methods
*/

#ifndef __INSERT_H
#define __INSERT_H

struct operation;

int insert(int document_id, int insertion_point, char * insertion_contents, int parent_id, int client_id);

char * apply_insert(operation * subject, char * string);

#endif // __INSERT_H