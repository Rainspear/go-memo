import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Memo } from 'src/app/models/memo.model';

@Component({
  selector: 'app-memo-item',
  templateUrl: './memo-item.component.html',
  styleUrls: ['./memo-item.component.scss']
})
export class MemoItemComponent implements OnInit {
  @Output() clickDeleteMemo = new EventEmitter<string>();
  @Input() memo?: Memo;

  onClickDeleteMemo () {
    this.clickDeleteMemo.emit(this.memo?.id);
  }

  constructor() { }

  ngOnInit(): void {
  }

}
