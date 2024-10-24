import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-sum',
  standalone: true,
  imports: [FormsModule],
  styleUrl: './sum.component.css',
  template: `

<div class="calculator">
<h3>{{title}}</h3>
  <div>
    <input type="number" [(ngModel)]="num1" placeholder="Enter first number" />
  </div>
  <div>
    <input type="number" [(ngModel)]="num2" placeholder="Enter second number" />
  </div>
  <button (click)="calculateSum()">Calculate</button>
  @if (result !== null) {
    <p>The sum is: {{ result }}</p>
  }

</div>


  `
})
export class SumComponent {
  title = "Sum Calculator";
  num1: number = 0;
  num2: number = 0;
  result: number | null = null;

  calculateSum() {
    this.result = this.num1 + this.num2;
  }

}
