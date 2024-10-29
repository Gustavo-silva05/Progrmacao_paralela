CC = gcc
CFLAGS = -I/usr/include/tirpc/ -ltirpc

all:
	rpcgen notas.x
	$(CC) -o cliente cliente.c $(CFLAGS)
	$(CC) -o servidor notas_svc.c servidor.c $(CFLAGS)
		
clean:
	rm -f notas_svc.c notas_clnt.c notas.h servidor cliente *~
