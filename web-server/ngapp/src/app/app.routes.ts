import { Routes } from '@angular/router';
import { SumComponent } from './sum/sum.component';
import { SubtractionComponent } from './subtraction/subtraction.component';
import { HomeComponent } from './home/home.component';

export const routes: Routes = [
  { path: 'sum', component: SumComponent },
  { path: 'subtraction', component: SubtractionComponent },
  { path: 'home', component: HomeComponent },
  { path: '',   redirectTo: '/home', pathMatch: 'full' },
];
