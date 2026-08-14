import { Routes } from '@angular/router';
import { ProdutosForm } from './features/produtos/produtos-form/produtos-form';
import { ProdutosList } from './features/produtos/produto-list/produto-list';

export const routes: Routes = [
    { path: 'produtos/novo', component: ProdutosForm },
    { path: 'produtos', component: ProdutosList }
];
