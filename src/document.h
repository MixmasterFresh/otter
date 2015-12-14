/*
An implementation of the document data type
*/

#ifndef __DOCUMENT_H
#define __DOCUMENT_H

void create_document(char * document_contents, char * document_id);

char * create_document(char * document_contents);

void update_document(char * updated_document_contents, char * document_id);

void close_document(char * document_id);

#endif // __DOCUMENT_H
