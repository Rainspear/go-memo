import { Injectable } from '@angular/core';
import { Subject, Subscription } from 'rxjs';
import { Memo } from '../models/memo.model';


@Injectable({
  providedIn: 'root'
})
export class MemoDetailService {
  memo?: Memo;
  selectMemo = new Subject<Memo>()
  subscriptionMemo: Subscription

  constructor() {
    this.subscriptionMemo = this.selectMemo.subscribe((memo) => {
        this.memo = memo;
    })
   }
}
