import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { Nota } from '../nota';

@Component({
  selector: 'app-notas-list',
  imports: [CommonModule, RouterLink],
  templateUrl: './notas-list.html',
  styleUrl: './notas-list.css'
})
export class NotasList implements OnInit {
  notas: any[] = [];
  imprimindo: number | null = null;
  erro: string | null = null;
  notaSelecionada: any = null;

  constructor(private notaService: Nota, private cdr: ChangeDetectorRef) {}

  ngOnInit() {
    this.carregar();
  }

  carregar() {
    this.notaService.listar().subscribe({
      next: (data) => {
        this.notas = data;
        this.cdr.detectChanges();
      },
      error: (err) => console.error('Erro ao listar notas', err)
    });
  }

  verDetalhes(nota: any) {
    this.notaSelecionada = nota;
  }

  fecharDetalhes() {
    this.notaSelecionada = null;
  }

  imprimir(numero: number) {
    this.imprimindo = numero;
    this.erro = null;

    this.notaService.imprimir(numero).subscribe({
      next: () => {
        this.imprimindo = null;
        this.carregar();
      },
      error: (err) => {
        this.imprimindo = null;
        this.erro = err.error?.error || 'Erro ao imprimir nota';
        this.cdr.detectChanges();
      }
    });
  }
}