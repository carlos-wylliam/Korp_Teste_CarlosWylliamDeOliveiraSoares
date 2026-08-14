import { Component } from '@angular/core';
import { FormBuilder, FormGroup, FormArray, Validators, ReactiveFormsModule } from '@angular/forms';

@Component({
  selector: 'app-notas-form',
  imports: [ReactiveFormsModule],
  templateUrl: './notas-form.html',
  styleUrl: './notas-form.css',
})

export class NotasForm {
  form: FormGroup;

  constructor(private fb: FormBuilder) {
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
}
