import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterLink } from '@angular/router';
import { Produto } from '../produto';

@Component({
  selector: 'app-produtos-list',
  imports: [CommonModule, RouterLink],
  templateUrl: './produto-list.html',
  styleUrl: './produto-list.css'
})
export class ProdutosList implements OnInit {
  produtos: any[] = [];
  erro: string | null = null;

  constructor(private produtoService: Produto, private cdr: ChangeDetectorRef) {}

  ngOnInit() {
    this.carregar();
  }

  carregar() {
    this.produtoService.listar().subscribe({
      next: (data) => {
        this.produtos = data;
        this.cdr.detectChanges();
      },
      error: (err) => console.error('Erro ao listar produtos', err)
    });
  }

  excluir(codigo: string) {
    if (!confirm(`Deseja realmente excluir o produto ${codigo}?`)) {
      return;
    }

    this.produtoService.excluir(codigo).subscribe({
      next: () => this.carregar(),
      error: (err) => {
        this.erro = err.error?.error || 'Erro ao excluir produto';
        this.cdr.detectChanges();
        setTimeout(() => {
          this.erro = null;
          this.cdr.detectChanges();
        }, 4000);
      }
    });
  }
}