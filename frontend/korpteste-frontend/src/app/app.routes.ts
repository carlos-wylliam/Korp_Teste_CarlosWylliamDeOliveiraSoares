import { Routes } from '@angular/router';
import { ProdutosForm } from './features/produtos/produtos-form/produtos-form';
import { ProdutosList } from './features/produtos/produto-list/produto-list';
import { NotasForm } from './features/notas/notas-form/notas-form';
import { NotasList } from './features/notas/notas-list/notas-list';
import { Home } from './features/home/home';

export const routes: Routes = [
    { path: '', component: Home },
    { path: 'produtos/novo', component: ProdutosForm },
    { path: 'produtos', component: ProdutosList },
    { path: 'notas/novo', component: NotasForm },
    { path: 'notas', component: NotasList }
];
