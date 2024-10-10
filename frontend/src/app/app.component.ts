import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ParkingLotComponent } from "./parking-lot/parking-lot.component";

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, ParkingLotComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss'
})
export class AppComponent {
  title = 'frontend';
}
