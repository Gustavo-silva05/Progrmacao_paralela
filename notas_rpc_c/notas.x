
/* notas.x */

#define PROGRAM_NUMBER 0x20000000
#define VERSION_NUMBER 1

program NOTAS_PROG {				   /* Programa */
	version NOTAS_VERSION {		   /* Versao */
		double OBTEM_NOTA(string) = 1; /* Funcao e numero */
	} = VERSION_NUMBER;			   /* Numero de versao */
} = PROGRAM_NUMBER;				   /* Numero de programa */
