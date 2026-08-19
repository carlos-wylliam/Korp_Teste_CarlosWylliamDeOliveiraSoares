import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class Nota {
  private apiUrl = 'http://localhost:8001/notas';

  constructor(private http: HttpClient) {}

  criar(nota: { status: string; itens: { produtoCodigo: string; quantidade: number }[] }) {
    return this.http.post(this.apiUrl, nota);
  }

  listar() {
    return this.http.get<any[]>(this.apiUrl);
  }

  imprimir(numero:number) {
    return this.http.put(`${this.apiUrl}/${numero}/imprimir`, {});
  }

  excluir(numero: number) {
  return this.http.delete(`${this.apiUrl}/${numero}`);
  }
}