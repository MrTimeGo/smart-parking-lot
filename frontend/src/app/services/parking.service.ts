import { inject, Injectable } from '@angular/core';
import * as signalR from '@microsoft/signalr';
import { BehaviorSubject } from 'rxjs';
import { ActionType, LogInfo, LotInfo } from '../models/car-info.model';
import { environment } from '../../environments/environment';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class ParkingService {
  private baseUrl = environment.baseUrl;
  private hubConnection: signalR.HubConnection;
  private _logs$: BehaviorSubject<LogInfo[]> = new BehaviorSubject<LogInfo[]>(
    []
  );
  private _lots$: BehaviorSubject<LotInfo[]> = new BehaviorSubject<LotInfo[]>(
    Array.from({ length: 20 }).map((_, i) => ({ place: i + 1 } as LotInfo))
  );

  private http = inject(HttpClient);

  get logs$() {
    return this._logs$.asObservable();
  }

  get lots$() {
    return this._lots$.asObservable();
  }

  constructor() {
    this.loadLogs();
    this.loadLots();

    this.hubConnection = new signalR.HubConnectionBuilder()
      .withUrl(this.baseUrl + '/parkingLotHub') // SignalR hub URL
      .build();

    this.connect();

    this.registerActionLogHandler();
  }

  connect() {
    this.hubConnection
      .start()
      .then(() => {
        console.log('Connection established with SignalR hub');
      })
      .catch((error) => {
        console.error('Error connecting to SignalR hub:', error);
      });
  }

  registerActionLogHandler() {
    this.hubConnection.on('action_update', (log: LogInfo) => {
      this._logs$.next([log, ...this._logs$.value]);
      this.updateLot(
        log.place,
        log.action == ActionType.Enter ? log.plateNumber : undefined
      );
    });
  }

  updateLot(place: number, plateNumber?: string) {
    const lots = this._lots$.value;
    const lotIndex = lots.findIndex((l) => l.place == place);
    lots[lotIndex].plateNumber = plateNumber;

    this._lots$.next(lots);
  }

  loadLots() {
    this.http
      .get<LotInfo[]>(`${this.baseUrl}/api/ParkingLot`)
      .subscribe((lots) => {
        lots.forEach((lot) => this.updateLot(lot.place, lot.plateNumber));
      });
  }

  loadLogs() {
    this.http
      .get<LogInfo[]>(`${this.baseUrl}/api/ParkingLot/action-logs`)
      .subscribe((logs) => {
        this._logs$.next(logs);
      });
  }
}
