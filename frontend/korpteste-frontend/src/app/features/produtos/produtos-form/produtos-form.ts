import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { Produto } from '../produto';

@Component({
  selector: 'app-produtos-form',
  imports: [ReactiveFormsModule],
  templateUrl: './produtos-form.html',
  styleUrl: './produtos-form.css',
})
export class ProdutosForm {
  form: FormGroup;

  constructor(private fb: FormBuilder, private produtoService: Produto) {
    this.form = this.fb.group({
      codigo: ['', Validators.required],
      descricao: ['', Validators.required],
      saldo: [0, [Validators.required, Validators.min(0)]]
    });
  }

  salvar() {
    if (this.form.valid) {
      this.produtoService.cadastrar(this.form.value).subscribe({
        next: () => console.log('Produto cadastrado com sucesso'),
        error: (err) => console.error('Erro ao cadastrar produto', err)
      });
    }
  }
}
