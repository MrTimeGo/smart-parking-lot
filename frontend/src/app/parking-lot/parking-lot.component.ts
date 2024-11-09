import { Component, inject } from '@angular/core';
import { ActionType, LogInfo, LotInfo } from '../models/car-info.model';
import { CommonModule } from '@angular/common';
import { ParkingService } from '../services/parking.service';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-parking-lot',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './parking-lot.component.html',
  styleUrl: './parking-lot.component.scss',
})
export class ParkingLotComponent {
  private parkingService = inject(ParkingService);

  public ActionType = ActionType;

  logs$: Observable<LogInfo[]> = this.parkingService.logs$;

  lots$: Observable<LotInfo[]> = this.parkingService.lots$;

  defaultPicture =
    'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ-0WWezXf0CnIZ_wI-_upBPzq9-CVOu4_q_g&s';

  constructor() {}
}
