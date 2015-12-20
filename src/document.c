#include "document.h"
#include "operation.h"
#include "buffer.h"

void create_document(char * document_contents, char * document_id){
  //TODO: implement method
}

char * create_document(char * document_contents){
  //TODO: implement method
}

void update_document(char * updated_document_contents, char * document_id){
  //TODO: implement method
}

char ** get_document_ids(int start, int quantity){
  //TODO: implement method
}

buffer_container * get_buffer_container(char * document_id){
  //TODO: implement method
}

buffer_container * get_buffer_container_without_id(char * document_id, int client_id){
  //TODO: implement method
}

buffer * get_buffer_with_id(char * document_id, int client_id){
  //TODO: implement method
}

void close_document(char * document_id){
  //TODO: implement method
}
