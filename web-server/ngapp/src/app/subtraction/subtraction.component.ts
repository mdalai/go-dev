import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';

@Component({
  selector: 'app-subtraction',
  standalone: true,
  imports: [FormsModule],
  styleUrl: './subtraction.component.css',
  template: `

<div class="calculator">
  <h3>{{title}}</h3>
  <div>
    <input type="number" [(ngModel)]="num1" placeholder="Enter first number" />
  </div>
  <div>
    <input type="number" [(ngModel)]="num2" placeholder="Enter second number" />
  </div>
  <button (click)="calculateSubtraction()">Calculate</button>
  @if (result !== null) {
    <p>The subtraction is: {{ result }}</p>
  }

</div>

  `
})
export class SubtractionComponent {

  title = "Subtraction Calculator";

  num1: number = 0;
  num2: number = 0;
  result: number | null = null;

  calculateSubtraction() {
    this.result = this.num1 - this.num2;
  }


}
