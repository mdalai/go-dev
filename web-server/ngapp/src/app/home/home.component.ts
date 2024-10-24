import { Component } from '@angular/core';
import { SumComponent } from '../sum/sum.component';
import { SubtractionComponent } from '../subtraction/subtraction.component';

@Component({
  selector: 'app-home',
  standalone: true,
  imports: [SumComponent, SubtractionComponent],
  template: `
    <app-sum></app-sum>
    <app-subtraction></app-subtraction>
  `,
  styles: ``
})
export class HomeComponent {

}
