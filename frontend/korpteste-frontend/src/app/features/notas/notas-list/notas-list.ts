import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { Nota } from '../nota';
import jsPDF from 'jspdf';

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

  constructor(private notaService: Nota, private cdr: ChangeDetectorRef) { }

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

    const nota = this.notas.find(n => n.numero === numero);

    this.notaService.imprimir(numero).subscribe({
      next: () => {
        this.imprimindo = null;
        this.gerarPdf(nota);
        this.carregar();
      },
      error: (err) => {
        this.imprimindo = null;
        this.erro = err.error?.error || 'Erro ao imprimir nota';
        this.cdr.detectChanges();
      }
    });
  }

  gerarPdf(nota: any) {
    const doc = new jsPDF();

    doc.setFontSize(18);
    doc.text('Nota Fiscal', 14, 20);

    doc.setFontSize(11);
    doc.text(`Número: ${nota.numero}`, 14, 32);
    doc.text('Status: Fechada', 14, 39);

    doc.setFontSize(12);
    doc.text('Itens:', 14, 52);

    let y = 60;
    nota.itens.forEach((item: any) => {
      doc.text(`${item.produtoCodigo}  -  Quantidade: ${item.quantidade}`, 14, y);
      y += 8;
    });

    doc.save(`nota-fiscal-${nota.numero}.pdf`);
  }

  excluir(numero: number) {
    if (!confirm(`Deseja realmente excluir a nota ${numero}?`)) {
      return;
    }

    this.notaService.excluir(numero).subscribe({
      next: () => this.carregar(),
      error: (err) => {
        this.erro = err.error?.error || 'Erro ao excluir nota';
        this.cdr.detectChanges();
        setTimeout(() => {
          this.erro = null;
          this.cdr.detectChanges();
        }, 4000);
      }
    });
  }
}