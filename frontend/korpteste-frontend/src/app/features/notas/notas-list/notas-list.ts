import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Nota } from '../nota';

@Component({
  selector: 'app-notas-list',
  imports: [CommonModule],
  templateUrl: './notas-list.html',
  styleUrl: './notas-list.css',
})

export class NotasList implements OnInit {
  notas: any[] = [];

  constructor(private notaService: Nota) {}

  ngOnInit() {
    this.notaService.listar().subscribe({
      next: (data) => this.notas = data,
      error: (err) => console.error('Erro ao listar notas', err)
    });
  }
}