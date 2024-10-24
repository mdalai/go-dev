import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { SumComponent } from './sum/sum.component';
import { SubtractionComponent } from './subtraction/subtraction.component';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, SumComponent, SubtractionComponent],
  styles: ``,
  template: `

<router-outlet></router-outlet>

  `
})
export class AppComponent {
  title = 'webapp1';
}
