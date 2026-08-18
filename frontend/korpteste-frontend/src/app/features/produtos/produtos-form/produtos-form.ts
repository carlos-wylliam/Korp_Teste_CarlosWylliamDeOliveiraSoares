import { Component, ChangeDetectorRef } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { Produto } from '../produto';

@Component({
  selector: 'app-produtos-form',
  imports: [ReactiveFormsModule, CommonModule],
  templateUrl: './produtos-form.html',
  styleUrl: './produtos-form.css',
})


export class ProdutosForm {
  mensagemSucesso = false;
  mensagemErro: string | null = null;
  form: FormGroup;

  constructor(private fb: FormBuilder, private produtoService: Produto, private cdr: ChangeDetectorRef) {
    this.form = this.fb.group({
      codigo: ['', Validators.required],
      descricao: ['', [Validators.required, Validators.minLength(3)]],
      saldo: [0, [Validators.required, Validators.min(0)]]
    });
  }

  salvar() {
    if (this.form.valid) {
      this.mensagemErro = null;

      this.produtoService.cadastrar(this.form.value).subscribe({
        next: () => {
          this.mensagemSucesso = true;
          this.form.reset({ codigo: '', descricao: '', saldo: 0});
          this.cdr.detectChanges();
          setTimeout(() => {
            this.mensagemSucesso = false;
            this.cdr.detectChanges();
          }, 3000);
        },
        error: (err) => {
          this.mensagemErro = err.error?.error || 'Erro ao cadastrar produto';
          this.cdr.detectChanges();
          setTimeout(() => {
            this.mensagemErro = null;
            this.cdr.detectChanges();
          }, 4000);
        }
      });
    }
  }
}