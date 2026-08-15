import { Component, OnInit, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Produto } from '../produto';

@Component({
  selector: 'app-produto-list',
  imports: [CommonModule],
  templateUrl: './produto-list.html',
  styleUrl: './produto-list.css',
})
export class ProdutosList implements OnInit{
  produtos: any[] = [];

  constructor(private produtoService: Produto, private cdr: ChangeDetectorRef) {}

  ngOnInit() {
    this.produtoService.listar().subscribe({
      next: (data) => {
        this.produtos = data;
        this.cdr.detectChanges();
      },
      error: (err) => console.error('Erro ao listar produtos', err)
    });
  }
}
