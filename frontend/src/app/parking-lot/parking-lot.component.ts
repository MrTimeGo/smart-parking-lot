import { Component } from '@angular/core';
import { ActionType, CarInfo, LogInfo } from '../models/CarInfo';
import { CommonModule } from '@angular/common';
import { SignalrService } from '../services/signalr.service';

export interface LotInfo{
  place: number;
  car?: CarInfo;
}
@Component({
  selector: 'app-parking-lot',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './parking-lot.component.html',
  styleUrl: './parking-lot.component.scss',
})
export class ParkingLotComponent {
  constructor(private signalrService: SignalrService) {
    this.signalrService.startConnection().subscribe(() => {
      this.signalrService.logInfo().subscribe((log) => {
        this.logInfoProcessing(log);
      });
    });
  }

  logInfoProcessing(log: LogInfo) {
    const lotIndex = this.lots.findIndex((l) => l.place == log.place);
    if (log.action == ActionType.Enter) {
      this.lots[lotIndex].car = {
        image: log.image,
        at: log.at,
        plateNumber: log.plateNumber,
      } as CarInfo;
    } else {
      this.lots[lotIndex].car = undefined;
    }
    this.logs.push(log);
    this.currentPicture = log.image;
  }

  public ActionType = ActionType;
  public currentPicture =
    'https://spn-sta.spinny.com/blog/20221004191046/Hyundai-Venue-2022.jpg?compress=true&quality=80&w=1200&dpr=2.6';
  public logs = [
    {
      image:
        'https://images.unsplash.com/photo-1525609004556-c46c7d6cf023?q=80&w=1937&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      plateNumber: 'AA223344FF',
      at: new Date(),
      action: ActionType.Enter,
      place: 2,
    } as LogInfo,
    {
      image:
        'https://images.unsplash.com/photo-1525609004556-c46c7d6cf023?q=80&w=1937&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
      plateNumber: 'AA223344CC',
      at: new Date(),
      action: ActionType.Exit,
      place: 3,
      cost: 12,
    } as LogInfo,
  ];

  public lots = [
    {
      place: 1,
    } as LotInfo,
    {
      place: 2,
    } as LotInfo,
    {
      place: 3,
      car: {
        image:
          'https://images.unsplash.com/photo-1525609004556-c46c7d6cf023?q=80&w=1937&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
        plateNumber: 'AA223344FF',
        at: new Date(),
      } as CarInfo,
    } as LotInfo,
    {
      place: 4,
    } as LotInfo,
    {
      place: 5,
    } as LotInfo,
    {
      place: 6,
    } as LotInfo,
    {
      place: 7,
    } as LotInfo,
    {
      place: 8,
    } as LotInfo,
    {
      place: 9,
    } as LotInfo,
    {
      place: 10,
    } as LotInfo,
    {
      place: 11,
    } as LotInfo,
    {
      place: 12,
    } as LotInfo,
    {
      place: 13,
    } as LotInfo,
    {
      place: 14,
    } as LotInfo,
    {
      place: 15,
    } as LotInfo,
    {
      place: 16,
    } as LotInfo,
    {
      place: 17,
    } as LotInfo,
    {
      place: 18,
    } as LotInfo,
    {
      place: 19,
    } as LotInfo,
    {
      place: 20,
    } as LotInfo,
  ];
}

