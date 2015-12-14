/*
An implementation of the document data type
*/

#ifndef __DOCUMENT_H
#define __DOCUMENT_H

struct buffer;
struct buffer_container;

void create_document(char * document_contents, char * document_id);

char * create_document(char * document_contents);

void update_document(char * updated_document_contents, char * document_id);

char ** get_document_ids(int start, int quantity);

buffer_container * get_buffer_container(char * document_id);

buffer_container * get_buffer_container_without_id(char * document_id, int client_id);

buffer * get_buffer_with_id(char * document_id, int client_id);

void close_document(char * document_id);

#endif // __DOCUMENT_H
