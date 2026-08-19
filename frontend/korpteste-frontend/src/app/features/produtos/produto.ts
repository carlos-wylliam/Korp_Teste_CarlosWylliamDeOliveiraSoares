import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})

export class Produto {
  private apiUrl = 'http://localhost:8000/produtos';

  constructor(private http: HttpClient) {}

  cadastrar(produto: { codigo: string; descricao: string, saldo: number}) {
    return this.http.post(this.apiUrl, produto);
  }

  listar() {
    return this.http.get<any[]>(this.apiUrl);
  }

  excluir(codigo: string) {
  return this.http.delete(`${this.apiUrl}/${codigo}`);
  }
}
