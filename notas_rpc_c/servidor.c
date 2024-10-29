#include <stdio.h>
#include <string.h>

#define TAM_MAX_NOME 9
#define NUM_ALUNOS 6

typedef struct {
	char nome[TAM_MAX_NOME+1];
	double nota;
} NOTAS;

double *obtem_nota_1_svc(char **nome) {
	static NOTAS notas[NUM_ALUNOS] = {
		{"Alexandre", 9.5},
		{"Barbara",   8.5},
		{"Joao",      6.5},
		{"Maria",     9.0},
		{"Paulo",    10.0},
		{"Pedro",     7.0}
	};
 
	static double erro = -1.0;
	int i;

	for	(i = 0; i < NUM_ALUNOS; ++i) {
		if (strcmp(notas[i].nome,*nome) == 0) {
			return &(notas[i].nota);
		}
	}

	return &erro;
}

