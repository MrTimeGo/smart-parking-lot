export interface CarInfo {
  image: string;
  plateNumber: string;
  at: Date;
  //place: number;
}

export interface LogInfo extends CarInfo {
  cost?: number;
  place: number;
  action: ActionType;
}

export interface LotInfo {
  place: number;
  plateNumber?: string;
}

export enum ActionType {
  Enter,
  Exit
}
