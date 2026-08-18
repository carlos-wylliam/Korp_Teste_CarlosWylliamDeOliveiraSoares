import { CommonModule } from '@angular/common';
import { Component, ChangeDetectorRef } from '@angular/core';
import { FormBuilder, FormGroup, FormArray, Validators, ReactiveFormsModule } from '@angular/forms';
import { Nota } from '../nota';

@Component({
  selector: 'app-notas-form',
  imports: [ReactiveFormsModule, CommonModule],
  templateUrl: './notas-form.html',
  styleUrl: './notas-form.css',
})

export class NotasForm {
  mensagemSucesso = false;
  mensagemErro: string | null = null;
  form: FormGroup;

  constructor(private fb: FormBuilder, private notaService: Nota, private cdr: ChangeDetectorRef) {
    this.form = this.fb.group({
      status: ['Aberta'],
      itens: this.fb.array([])
    });
  }

  get itens() {
    return this.form.get('itens') as FormArray;
  }

  adicionarItem() {
    this.itens.push(this.fb.group({
      produtoCodigo: ['', Validators.required],
      quantidade: [1, [Validators.required, Validators.min(1)]]
    }));
  }

  removerItem(index:number) {
    this.itens.removeAt(index);
  }
  
  salvar() {
    if (this.form.valid) {
      this.mensagemErro = null;

      this.notaService.criar(this.form.value).subscribe({
        next: () => {
          this.mensagemSucesso = true;
          this.itens.clear();
          this.cdr.detectChanges();
          setTimeout(() => {
            this.mensagemSucesso = false;
            this.cdr.detectChanges();
          }, 3000);
        },
        error: (err) => {
          this.mensagemErro = err.error?.error || 'Erro ao cadastrar nota';
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
