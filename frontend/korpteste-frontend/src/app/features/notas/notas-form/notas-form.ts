import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { FormBuilder, FormGroup, FormArray, Validators, ReactiveFormsModule } from '@angular/forms';
import { Nota } from '../nota';

@Component({
  selector: 'app-notas-form',
  imports: [ReactiveFormsModule, CommonModule],
  templateUrl: './notas-form.html',
  styleUrl: './notas-form.css',
})

export class NotasForm {
  form: FormGroup;

  constructor(private fb: FormBuilder, private notaService: Nota) {
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
      this.notaService.criar(this.form.value).subscribe({
        next: () => console.log('Nota cadastrada com sucesso'),
        error: (err) => console.error('Erro ao cadastrar nota', err)
      });
    }
  }
}
